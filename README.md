# goreader-client
Card Reader client application for Raspberry Pi, written in Go

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

Use connection diagram to connect RFID reader to GPIO on Raspberry Pi
Use connection diagram to connect LCD to I2C bus on Raspberry Pi
Use connection diagram to connect LEDs, resistors, and Piezo to appropriate GPIO connections on Raspberry Pi

Install GO on Raspberry Pi
Enable I2C, GPIO access on Raspberry Pi

Use the following commands to install wiringPi:
cd
sudo apt-get install libi2c-dev
git clone git://git.drogon.net/wiringPi
cd wiringPi
./build

