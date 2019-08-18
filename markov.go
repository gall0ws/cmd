// Markov text generator.
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

type State struct {
	pref []string
	suff []string
}

type Chain struct {
	states  map[string]*State
	current *State
}

func NewChain(rd io.Reader, prefLen int) (*Chain, error) {
	if prefLen <= 0 {
		return nil, errors.New("prefix length must be greater than 0")
	}
	mc := Chain{
		states: make(map[string]*State),
	}
	pref := make([]string, prefLen)
	brd := bufio.NewReader(rd)
	s := bufio.NewScanner(brd)
	s.Split(bufio.ScanWords)

	// Initialise prefix.
	for i := 0; i < prefLen; i++ {
		if !s.Scan() {
			return nil, errors.New("input too small")
		}
		pref[i] = s.Text()
	}

	// Consume the reader until EOF.
	for s.Scan() {
		w := s.Text()
		mc.add(pref, w)
		pref = append(pref[1:], w)
	}
	mc.Reset()
	return &mc, nil
}

func (c *Chain) add(pref []string, w string) {
	key := strings.Join(pref, " ")
	state := c.states[key]
	if state != nil {
		state.suff = append(state.suff, w)
		return
	}
	buf := make([]string, len(pref))
	copy(buf, pref)
	c.states[key] = &State{
		pref: buf,
		suff: []string{w},
	}
}

func (c *Chain) Next() string {
	if c.current == nil {
		return ""
	}
	i := rand.Intn(len(c.current.suff))
	w := c.current.suff[i]

	// Update current state.
	nextPref := append(c.current.pref[1:], w)
	c.current = c.states[strings.Join(nextPref, " ")]
	return w
}

func (c *Chain) Reset() {
	n := rand.Intn(len(c.states))
	i := 0
	for k := range c.states {
		if i == n {
			c.current = c.states[k]
			return
		}
		i++
	}
}

func main() {
	var cfg struct {
		prefixLength int
		wordCount    int
		seed         int64
	}
	flag.IntVar(&cfg.prefixLength, "p", 2, "prefix length")
	flag.IntVar(&cfg.wordCount, "w", 100, "word count output")
	flag.Int64Var(&cfg.seed, "s", time.Now().UnixNano(), "seed to initialise the PRNG")
	flag.Parse()

	rand.Seed(cfg.seed)
	mc, err := NewChain(os.Stdin, cfg.prefixLength)
	if err != nil {
		fmt.Fprintln(os.Stderr, "fatal error:", err)
		os.Exit(1)
	}
	for i := 0; i < cfg.wordCount; i++ {
		w := mc.Next()
		if w == "" {
			mc.Reset()
			continue
		}
		fmt.Printf("%s ", w)
	}
	fmt.Println()
}
