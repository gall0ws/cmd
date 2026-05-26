#include <err.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct line {
	char        *str;
	struct line *next;
};

struct line *
push(struct line *l, const char *str)
{
	struct line *p;

	p = malloc(sizeof(*p));
	if (!p) {
		err(1, "malloc");
	}
	p->next = l;
	p->str = strdup(str);
	if (!p->str) {
		err(1, "strdup");
	}

	return p;
}

void
consume(struct line *l)
{
	struct line *p;

	while (l) {
		fputs(l->str, stdout);
		p = l;
		l = l->next;

		free(p->str);
		free(p);
	}
}

void
tac(FILE *fp, const char *s)
{
	char buf[4096];
	struct line *p = NULL;

	while (fgets(buf, sizeof(buf), fp)) {
		p = push(p, buf);
	}
	if (!feof(fp)) {
		warn("error reading %s", s);
	}
	consume(p);
}

int
main(int argc, char **argv)
{
	FILE *fp;
	int i, retv = 0;

	if (argc == 1) {
		tac(stdin, "<stdin>");
	} else {
		for (i=1; i<argc; i++) {
			fp = fopen(argv[i], "r");
			if (!fp) {
				warn("could not open %s", argv[i]);
				retv = 1;
				continue;
			}
			tac(fp, argv[i]);
			fclose(fp);
		}
	}

	return retv;
}
