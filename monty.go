// Monty Hall problem simulator I wrote to win an argument on reddit

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const Winner = true

var (
	rounds = flag.Int("r", 1000, "rounds")
	doors  = flag.Int("d", 3, "doors")
)

func monty(player Player, doors int) bool {
	winning := rand.Intn(doors)
	chosed := player.Pick(doors)
	/*
	 * Monty opens (doors-2) losing doors
	 */
	if player.WantSwitch() {
		return chosed != winning
	}
	return chosed == winning
}

func test(player Player) (wins int) {
	for i := 0; i < *rounds; i++ {
		if monty(player, *doors) == Winner {
			wins++
		}
	}
	return wins
}

func main() {
	flag.Parse()
	if *doors < 2 {
		fmt.Fprintln(os.Stderr, "invalid number of doors")
		return
	}
	rand.Seed(time.Now().Unix())
	fmt.Println("switcher:", test(new(Switcher)))
	fmt.Println("keeper:  ", test(new(Keeper)))
}

type Player interface {
	Pick(n int) int
	WantSwitch() bool
}

type player struct{}

func (p player) Pick(n int) int {
	return rand.Intn(n)
}

type Switcher struct{ player }

func (s Switcher) WantSwitch() bool {
	return true
}

type Keeper struct{ player }

func (k Keeper) WantSwitch() bool {
	return false
}
