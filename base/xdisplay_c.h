#include "xdisplay.h"
#include <stdio.h> /* For fputs() */
#include <stdlib.h> /* For atexit() */
#include <string.h>

static Display *mainDisplay = NULL;
static int registered = 0;
static char *displayName = ":0.0";
static int hasDisplayNameChanged = 0;

Display *XGetMainDisplay(void)
{
	/* Close the display if displayName has changed */
	if (hasDisplayNameChanged) {
		XCloseMainDisplay();
		hasDisplayNameChanged = 0;
	}

	if (mainDisplay == NULL) {
		/* First try the user set displayName */
		mainDisplay = XOpenDisplay(displayName);

		/* Then try using environment variable DISPLAY */
		if (mainDisplay == NULL) {
			mainDisplay = XOpenDisplay(NULL);
		}

		if (mainDisplay == NULL) {
			fputs("Could not open main display\n", stderr);
		} else if (!registered) {
			atexit(&XCloseMainDisplay);
			registered = 1;
		}
	}

	return mainDisplay;
}

void XCloseMainDisplay(void)
{
	if (mainDisplay != NULL) {
		XCloseDisplay(mainDisplay);
		mainDisplay = NULL;
	}
}
char *strdupTemp(const char *s)
{
   size_t  len = strlen(s) +1;
   void *new = malloc(len);
   if (new == NULL)
      return NULL;
   return (char *)memcpy(new,s,len);
}
void setXDisplay(char *name)
{
	displayName = strdupTemp(name);
	hasDisplayNameChanged = 1;
}

char *getXDisplay(void)
{
	return displayName;
}


