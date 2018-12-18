/*
GoReader Client v0.5
(c) 2018 Jacob Bogner
This application runs on the card reader client (Raspberry Pi v3)
Objectives:
Listen to RFID reader for a new card scan
Upon reading a new scan, prompt user to enter a transaction amount
Connect to server and transmit the user ID and requested amount
Display an "approved" or "denied" message based off the server's reply
*/
package main

/*
  #cgo CFLAGS: -I.
  #cgo LDFLAGS: -L. -lwiringPi -lwiringPiDev
  #include "C_Libraries/hello.c"
  #include "C_Libraries/i2cdriver.c"
*/
import "C"
import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"log"
	"net"
	"os"
	"time"
)

func setup() {
	listener, err := net.Listen("tcp", "192.168.1.1:8000") //setup the network connection (set IP to server IP)
	if err != nil {                                        //handle connection error
		log.Fatal(err)
	}
	C.configureDisplay() //Setup the LCD
	mystring := "testing... wait"
	ctxt := C.CString(mystring) //Print the test message to LCD
	C.printLcd(ctxt)

}

var (
	// Use mcu pin 10, corresponds to physical pin 19 on the pi
	pin = rpio.Pin(10)
)

func main() {
	setup()
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	pin.Output()

	// Toggle pin 20 times
	for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}
}
