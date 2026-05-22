CC?=		clang
GO?=		go
STRIP?=		strip
INSTALL?=	install
CFLAGS?=	-O2 -g
BINDIR?=	$(HOME)/bin

CFLAGS+=\
	-Wall\
	-Wextra\
	-Werror\
	-Wno-format-zero-length

BINS=	avatar markov monty setsid statfs tsize usleep

all:	$(BINS)

avatar: avatar.go
	$(GO) build -o $@ $^

markov: markov.go
	$(GO) build -o $@ $^

monty:  monty.go
	$(GO) build -o $@ $^

statfs: statfs.o statfs.c
	$(CC) -o $(@) $<

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
	$(STRIP) $^
	$(INSTALL) -Cv $^ $(BINDIR)

.PHONY:	all clean install
