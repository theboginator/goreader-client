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
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	replyChan = make(chan string) //Channel for accepted/denied message coming from server
)

func configureConnections(conn net.Conn) { //Setup connection manager to handle incoming/outgoing data
	go replyReader(conn)
}

func sendID(conn net.Conn, id string) { //Send a "accepted/declined" message in the form of a boolean
	fmt.Fprintln(conn, id)
}

func replyReader(conn net.Conn) { //Read userID sent from Pi
	input := bufio.NewScanner(conn)
	for input.Scan() {
		data := input.Text()
		replyChan <- data
	}
}

func setup() { //Sets up the network connection and LCD, RFID reader, and LEDs

	C.setupPi() //Setup the LCD
	printLcd("testing... wait")

}

func printLcd(input string) {
	ctxt := C.CString(input) //Print the test message to LCD
	C.printLcd(ctxt)
}

func main() {
	conn, err := net.Dial("tcp", "192.168.1.1:6000") //setup the network connection (set IP to server IP)
	for err != nil {                                 //handle connection error
		log.Fatal(err)
		continue
	}
	setup()
	for {
		go configureConnections(conn)
		fmt.Println("Enter account #: ")
		reader := bufio.NewReader(os.Stdin)
		account, _ := reader.ReadString('\n')
		fmt.Println("Enter transaction amount: ")
		charge, _ := reader.ReadString('\n')
		msg := "Charging account " + account
		C.printLcd(msg)
		transmit := account + "," + charge
		sendID(conn, transmit)
		reply := <-replyChan
		C.printLcd(reply)
	}

}
