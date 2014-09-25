package main

import "fmt"

func main() {
    //synchronous example
    cs := make(chan string, 1)
    receive(cs)
    sending(cs)
    outputit(cs)
}

func receive(cs chan string){
    message := "test 1"
    cs <- message
}    

func sending(cs chan string){
    message := <-cs
    //alter message and send back to channel
    cs <- message + " and test 2"
}

func outputit(cs chan string){
    output := <-cs
    fmt.Println(output)
}
