package main

import "fmt"


func sender(){
	fmt.Println("Sender is sending a message")
}
func main() {

	fmt.Println("LEARNIG GOROUTINE..")
	sender()
}