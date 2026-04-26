#include <sys/types.h>
#include <unistd.h>
#include <stdio.h>

int main(int argc, char **argv)
{
	pid_t p;

	if (argc < 2) {
		fputs("usage: setsid command [arguments]\n", stderr);
		return 1;
	}

	p = fork();
	if (p < 0) {
		perror("fork");
		return 1;
	} else if (p > 0) {
		return 0;
	}

	if (setsid() < 0) {
		perror("setsid");
		return 1;
	}
	if (execvp(argv[1], argv+1) < 0) {
		perror("execvp");
		return 1;
	}
	return 0;
}
