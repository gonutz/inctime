This is a Go program to change the system time on your Windows computer.

Example Usage
-------------

	inctime 5h       Add 5 hours to the current system time.
	inctime -5h      Subtract 5 hours from the current system time.
	inctime 24h      Add a day to the current system time.
	inctime 5h4m3s   Add 5 hours, 4 minutes and 3 seconds to the current system time.

Supported suffixes are h (hours), m (minutes), s (seconds) and ms (milliseconds).

Installation
------------

You need the Go programing language installed: https://golang.org/
Use the go tool to get the program:

	go get -u -ldflags="-H=windowsgui" github.com/gonutz/inctime

where -u is for getting the latest version online and -ldflags="-H=windowsgui" has Go build a Windows GUI program which does not open a console window if started from the "Run As" dialog.
