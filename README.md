# goreader-client
Card Reader client application for Raspberry Pi, written in Go

To download application onto the Pi,
at command prompt, type "git clone https://github.com/theboginator/goreader-client"

Needed hardware:
Raspberry Pi v3
16x2 LCD
RC522 RFID reader
2x Blue LED
1x Red LED
1x Green LED
2x 200-Ohm Resistors
1x Piezo
Hookup wire

Use connection diagram to connect RFID reader to GPIO on Raspberry Pi - located in setupImages directory
Connect LCD to I2C bus on Raspberry Pi. Ensure power supply is 5v
Connect the Blue LEDs to GPIO pins 5 and 25
Connect the Red LED to GPIO pin 26
Connect the Green LED to GPIO pin 20
The negative leads of the LEDs may be connected directly to ground

Install GO on Raspberry Pi

Enable I2C, GPIO access on Raspberry Pi by using "sudo raspi-config" (reboot if necessary)

Use the following commands to install wiringPi, which provides needed hardware interfaces:
cd
sudo apt-get install libi2c-dev
git clone git://git.drogon.net/wiringPi
cd wiringPi
./build

ensure Python 2.7 is installed
Then use the following commands to install the python SPI Drivers for the RFID reader:
git clone https://github.com/lthiery/SPI-Py.git
cd ~/SPI-Py
sudo python setup.py install

cd ~
git clone https://github.com/pimylifeup/MFRC522-python.git

All prerequisites are theoretically complete

cd goreader-client/src
go build main.go
go run main.go



