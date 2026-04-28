#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv)
{
	int us;
	char *p;

	if (argc < 2) {
		fputs("usage: usleep number\n", stderr);
		return 1;
	}
	us = strtol(argv[1], &p, 0);
	if (*p != '\0') {
		fprintf(stderr, "usleep: invalid time interval: %s\n", argv[1]);
		return 1;
	}
	usleep(us);
	return 0;
}
