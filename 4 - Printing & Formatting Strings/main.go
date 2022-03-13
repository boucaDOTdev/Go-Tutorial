package main

import "fmt"

func main(){
	age := 22
	name := "Alex"
	
	//print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("hello ninjas!")
	fmt.Println("goodbye ninjas!")
	fmt.Println("My age is", age, "and my name is", name)

	// Printf (formatted string) %_ = format specifier
	fmt.Printf("My age is %v and my name is %v \n",age, name)
	fmt.Printf("My age is %q and my name is %q \n",age, name)
	fmt.Printf("age is of type %T \n",age)
	fmt.Printf("you scored %0.2f points! \n",255.55)

	// Sprintf (save formated strings)
	var str = fmt.Sprintf("My age is %v and my name is %v \n",age, name)
	fmt.Println("the saved string is:", str)
}