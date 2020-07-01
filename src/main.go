/*
	Summary

	Defer
		Used to delay execution of a statement until function exits
		Useful to group 'open' and '' close functions together
			be carefu; in loops
		Run in LIFO(Last-in, first-out) order
		Arguments evaluated at time defer is executed, not at time of called function execution
	Panic
		Occur when program cannot continue at all
			Dont use when file cant be opened, unless its critical
			Use for unrecoverable events-cannot obtaion TCP port for web server
		Function will stop executing
			Deferred functions will still fire
		If nothing handles panic, program will exit
	Recover
		Used to recover from panics
		Only useful in deferred functions
		Current function will not attempt to continue, but higher functions in call stack will
*/

package main

import (
	"fmt"
	"log"
)

func main(){
	helloWorldExample()

	// it works like first in last out!!

	// we can use defer to close http connections
	example() // start
	// in go we dont have exeptions
	// in go there are panic

	// also we can make our own panics!

	/*
	panic("something bad happened!") // panic: something bad happened!
	*/

	// in this example is shown how to avoid panicking, and make an exeption
	panicking()
}

func helloWorldExample(){
	// we can use defer to make go execute command before the return statement

	defer fmt.Println("Bye world")
	fmt.Println("Hello world")
	/*
	so it will execute like this:
	
	Hello world!
	Bye world!
	*/
}

func example(){
	a:= "start"
	defer fmt.Println(a)
	a = "end"
	// the result of this function will be start!!!!!
}

func panicking(){
	fmt.Println("About to panic")
	defer func(){
		if err := recover();err != nil{
			log.Println("Error: ", err)
		}
	}()
	panic("Something bad happened!")
	fmt.Println("Done panicking")
}

