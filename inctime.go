package main

import (
	"os"
	"time"
	"unsafe"

	"github.com/gonutz/w32"
)

const adminFlag = "we_are_admin"

func main() {
	fail := func(msg string) {
		w32.MessageBox(0, msg, "Error", w32.MB_OK|w32.MB_TOPMOST|w32.MB_ICONERROR)
		panic(msg)
	}

	// Check the arguments.
	weAreAdmin := false
	if len(os.Args) == 3 && os.Args[2] == adminFlag {
		weAreAdmin = true
		os.Args = os.Args[:2]
	}

	if len(os.Args) != 2 {
		fail(`Please pass exactly one argument, the time to add to your current system time, e.g. "5h" or "-5h5m".`)
	}

	dt, err := time.ParseDuration(os.Args[1])
	if err != nil {
		fail("Failed to parse " + err.Error() + ".")
	}

	if weAreAdmin {
		// We are admin, the duration is valid, get the system time, convert it
		// to a Windows file time which is really a 64 bit integer that we can
		// add our duration to. The file time is in 100 nanoseconds..
		sysTime := w32.GetSystemTime()
		fileTime, ok := w32.SystemTimeToFileTime(sysTime)
		if !ok {
			fail("SystemTimeToFileTime failed.")
		}
		ns := dt.Nanoseconds()
		ticks := ns / 100
		*(*int64)(unsafe.Pointer(&fileTime)) += ticks
		sysTime, ok = w32.FileTimeToSystemTime(fileTime)
		if !ok {
			fail("FileTimeToSystemTime failed.")
		}
		if ok {
			if !w32.SetSystemTime(sysTime) {
				fail("SetSystemTime failed.")
			}
		}
	} else {
		// Only the duration was passed, we need to get admin rights to change
		// the current time so execute ourselves with "runas" which means "run
		// as admin". Also pass ourselves a flag so we know we are admin.
		w32.ShellExecute(
			0,
			"runas",
			os.Args[0],
			os.Args[1]+" "+adminFlag,
			".",
			w32.SW_SHOWNORMAL,
		)
	}
}
