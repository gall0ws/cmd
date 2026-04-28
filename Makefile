CC?=		clang
GO?=		go
INSTALL?=	install
CFLAGS?=	-O2 -g
BINDIR?=	$(HOME)/bin

CFLAGS+=\
	-Wall\
	-Wextra\
	-Werror\
	-Wno-format-zero-length\
	-Wno-missing-braces\
	-Wno-parentheses\
	-Wno-sign-compare

BINS=	markov monty setsid tsize usleep

all:	$(BINS)

markov: markov.go
	$(GO) build -o $@ $^

monty:  monty.go
	$(GO) build -o $@ $^

setsid: setsid.o setsid.c
	$(CC) -o $(@) $<

tsize:  tsize.o tsize.c
	$(CC) -lncurses -o $@ $<

usleep: usleep.o usleep.c
	$(CC) -o $(@) $<

%.o:	%.c
	$(CC) $(CFLAGS) -c -o $@ $(@:.o=.c)

clean:
	rm -f *.o $(BINS)

install: $(BINS)
	$(INSTALL) -Csv $^ $(BINDIR)

.PHONY:	all clean install
