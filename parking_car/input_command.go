package parking_car

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cars = []car{}
var current_slot_empty int = 1
var slot_total int

func CmdOpr() {
	fmt.Println("")
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\n", "", -1)

		args := strings.Split(text, " ")

		switch args[0] {
		case "create_parking_lot":

			slot_total, _ = strconv.Atoi(args[1])
      fmt.Println("Created a parking lot with ",slot_total," slots")

		case "park":

			if len(cars) == slot_total {

				fmt.Println("Sorry, parking lot is full")

			} else {
				i := 1

				if is_slot_empty {
					for i <= slot_total {
						if i != cars[i-1].slot_no {
							c := car{
								slot_no: i,
								car_no:  args[1],
								color:   args[2],
							}
							cars = append(cars[:i-1], append([]car{c}, cars[i-1:]...)...)
							fmt.Println("Allocated slot number:", i)
							is_slot_empty = !is_slot_empty
							break
						}
						i++
					}

				} else {
					c := car{
						slot_no: current_slot_empty,
						car_no:  args[1],
						color:   args[2],
					}
					cars = append(cars, c)
					fmt.Println("Allocated slot number:", current_slot_empty)
					current_slot_empty++
				}
			}

		case "leave":

			i, _ := strconv.Atoi(args[1])
			j := 0

			for _, c := range cars {
				if c.slot_no == i {
					cars = append(cars[:j], cars[j+1:]...)
					is_slot_empty = !is_slot_empty
					fmt.Println("Slot No ", i, " is free")
				}
				j++
			}

		case "status":

			fmt.Println("Slot No.\t Registration No \t Colour")

      for _, c := range cars {

      	fmt.Println(c.slot_no, "\t\t", c.car_no, "\t\t", c.color)

      }

    case "registration_numbers_for_cars_with_colour":
      for _, c := range cars {

    		if c.color == args[1] {

    			fmt.Printf("%s, \t", c.car_no)

    		}
    	}



    	fmt.Printf("\n")

    case "slot_numbers_for_cars_with_colour":
      for _, c := range cars {

        if c.color == args[1] {

          fmt.Printf("%s, \t", strconv.Itoa(c.slot_no))

        }
			}

      fmt.Printf("\n")

      //slot_number_for_registration_number
    case "slot_number_for_registration_number":

      for _, c := range cars {

        if c.car_no == args[1] {

          fmt.Println(c.slot_no)

        }
        if c.car_no != args[1]{

        }else{
          fmt.Println("not found")
        }

      }

			// jika salah command
		default:

			fmt.Println("Incorrect Command")

		}
	}
}
