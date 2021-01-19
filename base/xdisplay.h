#pragma once
#ifndef XDISPLAY_H
#define XDISPLAY_H

#include <X11/Xlib.h>

/* Returns the main display, closed either on exit or when closeMainDisplay()
 * is invoked. This removes a bit of the overhead of calling XOpenDisplay() &
 * XCloseDisplay() everytime the main display needs to be used.
 *
 * Note that this is almost certainly not thread safe. */
Display *XGetMainDisplay(void);

/* Closes the main display if it is open, or does nothing if not. */
void XCloseMainDisplay(void);

#ifdef __cplusplus
extern "C"
{
#endif

char *getXDisplay(void);
void setXDisplay(char *name);
#if defined(USE_X11)
char * strdup(char* srcStr)
#endif
#ifdef __cplusplus
}
#endif

#endif /* XDISPLAY_H */
