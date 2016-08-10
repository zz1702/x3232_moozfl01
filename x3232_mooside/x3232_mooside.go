package main

/*
#include <termios.h>
#include <unistd.h>
*/
import "C"

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetPrefix("x3232_mooside ")
	log.SetFlags(log.Ltime)

	fd0 := C.int(os.Stdin.Fd())
	if C.isatty(fd0) == 1 {
		log.Print("stdin is a tty")
	} else {
		log.Print("stdin is NOT a tty")
		return
	}

	{
		var attr C.struct_termios
		if status, err := C.tcgetattr(fd0, &attr); status != 0 {
			log.Print("tcgetattr on fd0: ", err)
			return
		}
		defer C.tcsetattr(fd0, C.TCSANOW, &attr)
	}

	{
		var attr C.struct_termios
		if status, err := C.tcgetattr(fd0, &attr); status != 0 {
			log.Print("tcgetattr on fd0: ", err)
			return
		}
		attr.c_lflag &= ^C.tcflag_t(C.ECHO)
		if status, err := C.tcsetattr(fd0, C.TCSANOW, &attr); status != 0 {
			log.Print("tcsetattr on fd0: ", err)
			return
		}
	}

	ttySname := "ttyS0"
	ttyS, err := os.OpenFile("/dev/"+ttySname, os.O_RDWR, 0666)
	if err != nil {
		log.Print(err)
		return
	}
	defer ttyS.Close()

	ttySfd := C.int(ttyS.Fd())
	if C.isatty(ttySfd) == 1 {
		log.Print(ttySname, " is a tty")
	} else {
		log.Print(ttySname, " is NOT a tty")
		return
	}

	{
		var attr C.struct_termios

		// See http://elinux.org/RPi_Serial_Connection
		attr.c_cflag = C.CS8 | C.CREAD | C.CLOCAL // 8N1
		if status, err := C.cfsetispeed(&attr, C.B115200); status != 0 {
			log.Print("cfsetispeed: ", err)
			return
		}
		if status, err := C.cfsetospeed(&attr, C.B115200); status != 0 {
			log.Print("cfsetospeed: ", err)
			return
		}

		// blocking read
		attr.c_cc[C.VMIN] = 1
		attr.c_cc[C.VTIME] = 0

		if status, err := C.tcsetattr(ttySfd, C.TCSANOW, &attr); status != 0 {
			log.Print("tcsetattr ", ttySname, ": ", err)
			return
		}
	}

	go func() {
		inbox := make([]byte, 64)
		for {
			n, err := ttyS.Read(inbox)
			fmt.Print(string(inbox[:n]))
			if err != nil {
				log.Print("Read from ", ttySname, ": ", err)
				return
			}
		}
	}()

	{
		inbox := make([]byte, 64)
		for {
			n, err := os.Stdin.Read(inbox)
			if err != nil {
				log.Print("Read from stdin: ", err)
				return
			}
			if _, err = ttyS.Write(inbox[:n]); err != nil {
				log.Print("Write to ", ttySname, ": ", err)
				return
			}
		}
	}
}
