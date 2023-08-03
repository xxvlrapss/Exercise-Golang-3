package main 

import "fmt"

func print_hello() {
	fmt.Printf("Hello!\n")
}

func main() { 

		// defer	
		defer print_hello()

		// if, else, if, else

		age := 19
		
		if age < 18 {
			fmt.Printf("Not old enough!\n")
		} else if age > 60 {
			fmt.Printf("Too old!\n")
		} else { 
			fmt.Printf("Welcome!\n")
		}

		// switch 

		// grade := "B"

		// switch grade {
		// case "A":
		// 	fmt.Printf("Good Job!\n")
		// case "B", "C":
		// 	fmt.Printf("Nice Work!\n")
		// 	fmt.Printf("You can do better!\n")
		// case "F":
		// 	fmt.Printf("Please see the teacher\n")
		// 	fallthrough
		// default:
		// 	fmt.Printf("You have to try again!\n")
		// }


		// Looping
		// for i := 0; i < 5; i++ {
		// 	fmt.Printf("Looping %d\n", i)
		// }

		// j := 0
		// for j < 5 {
		// 	fmt.Printf("Looping %d\n", j)
		// 	j += 1
		// }

		// i := 0
		// for {
		// 	i += 1
		// 	if i == 2 {
		// 		continue
		// 	}
		// 			fmt.Printf("Looping %d\n", i)				
		// 			if i == 5 {
		// 			break
		// 	}
		// }

		// numbers := []int64{1, 2, 3, 4, 5}
		// for i, v := range numbers {
		// 	fmt.Printf("I got %d on index %d\n", i, v)
		// }
		
}