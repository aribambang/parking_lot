package main

import (
	parking_car "parking_lot/parking_car"
)

func main() {
	parking_car.Init()
	//Code for Command Shell Input
	if parking_car.Input == 1 {
		parking_car.CmdOpr()
	} else if parking_car.Input == 2 { //Code for file input
		parking_car.FileOpr()
	} else {

	}

}
