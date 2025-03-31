/*
Package pocketlog exposes an API to log your work.

First, instantiate a logger with pocketlog.New, and giving it a threshold level.
Messages of lesser criticality won't be logged.

Sharing the  logger is responsibility of the caller.

The logger can be called to log messages on five levels:

	-Debug: used to debug code
	-Info: valuable information
	-Warn: Possible problems
	-Error: error messages
	-Fatal: info about fatal errors
*/
package pocketlog
