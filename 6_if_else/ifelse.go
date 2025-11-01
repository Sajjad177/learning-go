package main

func main() {
	age := 20

	if age < 18 {
		println("Minor")
	} else {
		println("Adult")
	}


	mark := 85
	if mark >= 90 {
		println("Grade: A")
	} else if mark >= 80 {
		println("Grade: B")
	} else if mark >= 70 {
		println("Grade: C")
	} else {
		println("Grade: F")
	}


	num := -5
	if num > 0 {
		println("Positive")
	} else if num < 0 {
		println("Negative")
	} else {
		println("Zero")
	}


	if n := 10; n%2 == 0 {
		println("Even")
	} else {
		println("Odd")
	}

	voteAge := 18
	country := "bangladesh"

	if country == "bangladesh" {
		if voteAge <= 18 {
			println("Eligible to vote in Bangladesh")
		} else {
			println("Not eligible to vote in Bangladesh")
		}
	}


	temperature := 30

	if temperature > 30 {
		println("It's a hot day")
	}else if temperature >= 20 && temperature <= 30 {
		println("It's a warm day")
	}else {
		println("It's a cold day")
	}


	var role = "admin"
	var isAccess = true

	if role == "admin" && isAccess {
		println("Access granted to admin")
	} else {
		println("Access denied")
	}


}