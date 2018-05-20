package parking_car

import (
	"fmt"
)

var Input int

func Init() {
	fmt.Println("enter 1 for command or 2 for file")
	fmt.Scanf("%d", &Input)
}
