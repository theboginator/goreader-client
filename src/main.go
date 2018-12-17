package main

/*
  #cgo CFLAGS: -I.
  #cgo LDFLAGS: -L. -lwiringPi -lwiringPiDev
  #include "C_Libraries/hello.c"
  #include "C_Libraries/i2cdriver.c"
  #include "
*/
import "C"
import (
	"errors"
	"log"
)

func main() {
	//Call to void function without params
	err := hello()
	if err != nil {
		log.Fatal(err)
	}
}

//Hello is a C binding to the Hello World "C" program. As a Go user you could
//use now the Hello function transparently without knowing that it is calling
//a C function
func hello() error {
	_, err := C.hello() //We ignore first result as it is a void function
	if err != nil {
		return errors.New("error calling Hello function: " + err.Error())
	}
	_, err := C.showDisplay()
	if err != nil {
		return errors.New("error calling Driver function: " + err.Error())
	}
	return nil
}
