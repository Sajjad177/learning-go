package main

import "fmt"

// For loop is only way to loop in Go
func main() {
	//! while loop
	// sum := 1
	// for sum <= 3 {
	// 	println(sum)
	// 	sum += 1
	// }

	fmt.Println("traditional for loop.....")
	for i := 1; i <= 5; i++ {

		if(i == 3) {
			continue // skip the rest of the loop when i is 3
		}

		println(i)
	}

	fmt.Println("while style loop....." )
	sum :=1 
	for sum <= 5 {
		println(sum)
		sum +=1
	}


fmt.Println("Infinite loop example.....")
	count := 0
	for {
		fmt.Println("Count:", count)
		count++
		if count == 3 {
			break 
		}
	}

	fmt.Println("Range loop example.....")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Println("Index:", index, "Value:", value)
	}


	fmt.Println("Range loop with only value.....")
	person := map[string]string{
		"name": "Sajjad",
		"job":  "Developer",
	}
	for key, value := range person {
		fmt.Println(key, ":", value)
	}




}
