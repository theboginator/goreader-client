/*
This code is an i2c/GPIO combination driver built with code provided by WiringPi (GPIO) and Freenove (LCD)
*/
#include <stdlib.h>
#include <stdio.h>
#include <wiringPi.h>
#include <pcf8574.h>
#include <lcd.h>
#include <time.h>

#define pcf8574_address 0x27        // default I2C address of Pcf8574
#define BASE 64         // BASE is not less than 64
//////// Define the output pins of the PCF8574, which are directly connected to the LCD1602 pin.
#define RS      BASE+0
#define RW      BASE+1
#define EN      BASE+2
#define LED     BASE+3
#define D4      BASE+4
#define D5      BASE+5
#define D6      BASE+6
#define D7      BASE+7

int lcdhd;// used to handle LCD

int setupPi(void){
    int i;

    if(wiringPiSetup() == -1){ //handle an error should Pi setup fail
        printf("setup wiringPi failed !");
        return 1;
    }
    pcf8574Setup(BASE,pcf8574_address);// initialize PCF8574 LCD controller
    for(i=0;i<8;i++){
        pinMode(BASE+i,OUTPUT);     // set PCF8574 port to output mode
    }
    digitalWrite(LED,HIGH);     // turn on LCD backlight
    digitalWrite(RW,LOW);       // allow writing to LCD
    lcdhd = lcdInit(2,16,4,RS,EN,D4,D5,D6,D7,0,0,0,0);// initialize LCD and return “handle” used to handle LCD
    if(lcdhd == -1){
        printf("lcdInit failed !");
        return 1;
    }

}

void printLcd(char *text){
    lcdPosition(lcdhd,0,0);     // set the LCD cursor position to (0,0)
    lcdPrintf(lcdhd, text);     // Print data to the screen
}


