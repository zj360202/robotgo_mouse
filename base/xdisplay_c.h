#include "xdisplay.h"
#include <stdio.h> /* For fputs() */
#include <stdlib.h> /* For atexit() */

#if defined(USE_X11)
char * strdup(char* srcStr)
{
     int len=0;
     char *start=srcStr;
      if(srcStr!=NULL)
     {
     while(*srcStr++!='\0')
             len++;

         char *address=(char*)malloc(len+1);
         assert(address!=NULL);

         while((*address++=*start++)!='\0')
             NULL;
         return address-(len+1);
     }
     return NULL;
}
#endif

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

void setXDisplay(char *name)
{
	displayName = strdup(name);
	hasDisplayNameChanged = 1;
}

char *getXDisplay(void)
{
	return displayName;
}


