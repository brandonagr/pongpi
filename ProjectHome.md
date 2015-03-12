Uses an [RGB LED Strip](http://adafruit.com/products/306) and a Raspberry Pi to make a 1D pong game. Made with the help and laser cutter of the [Dallas Makerspace](http://dallasmakerspace.org)

There are some [existing similar projects](http://hackaday.com/2012/08/22/one-dimensional-pong-is-a-great-use-for-led-strips/), this will be different in that it uses the greater processing power in the Raspberry Pi to do more advanced effects / animations. All of the code is written in Go, using the [wiringPi](https://projects.drogon.net/raspberry-pi/wiringpi/) GPIO utility and hardware SPI to enable the communications with the buttons and LED strip.

<a href='http://www.youtube.com/watch?feature=player_embedded&v=_Y8oRO0JS2Q' target='_blank'><img src='http://img.youtube.com/vi/_Y8oRO0JS2Q/0.jpg' width='425' height=344 /></a>
![http://wiki.pongpi.googlecode.com/git/boxWithLights.jpg](http://wiki.pongpi.googlecode.com/git/boxWithLights.jpg)
![http://wiki.pongpi.googlecode.com/git/batterypi.jpg](http://wiki.pongpi.googlecode.com/git/batterypi.jpg)
![http://wiki.pongpi.googlecode.com/git/buildingBox.jpg](http://wiki.pongpi.googlecode.com/git/buildingBox.jpg)