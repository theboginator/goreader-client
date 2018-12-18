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
  #include "C_Libraries/i2cdriver.c"
*/
import "C"
import (
	"log"
	"net"
)

func setup() { //Sets up the network connection and LCD, RFID reader, and LEDs
	_, err := net.Dial("tcp", "192.168.1.1:8000") //setup the network connection (set IP to server IP)
	for err != nil {                              //handle connection error
		log.Fatal(err)
		continue
	}
	C.setupPi() //Setup the LCD
	printLcd("testing... wait")
	flashLED(lpin)
	printLcd("Checking LED...")

}

func flashLED(pin int) {
	C.writePin(pin, 1)

}

func printLcd(input string) {
	ctxt := C.CString(input) //Print the test message to LCD
	C.printLcd(ctxt)
}

func main() {
	setup()

}
