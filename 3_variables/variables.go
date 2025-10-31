package main


func main(){
	//variables
	var a int = 10
	var b string = "hello"
	var c bool = true
	var d float64 = 3.14

	//printing variables
	println(a)
	println(b)
	println(c)
	println(d)

	//short variable declaration
	x := 20
	y := "short variable declaration"
	z := false
	w := 6.28

	//printing short declared variables
	println(y)
	println(x)
	println(z)
	println(w)

	// another way to declare short variables
	var name = "another way to declare short variables"
	var age = 5
	var isGoFun = true

	println(name)
	println(age)
	println(isGoFun)

	// using init function to initialize variables
	var age2 int = 5

	println(age2)


	//if you declare value later then you need to add type
	var score int
	score = 100
	println(score)


}


