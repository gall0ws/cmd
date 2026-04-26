#include <curses.h>
#include <stdio.h>

int main()
{
	WINDOW *w;
	int y, x;

	w = initscr();
	if (w == NULL) {
		return 1;
	}
	getmaxyx(w, y, x);
	endwin();

	printf("%dx%d\n", x, y);
	return 0;
}
