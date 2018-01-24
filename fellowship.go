package main

import (
	"fmt"
)

type token struct {
	emitter int
	receiver int
	message string
}

func Operator(in, out chan token, id int) {
	operator := <-in
	fmt.Println("Channel - ", id)
	close(in)
	if operator.message == "yep"{
		out <- operator 
		return
	}
	if id == operator.receiver {
			fmt.Println("here is your message", operator.message)
			operator.message = "yep"
			if operator.emitter-operator.receiver > 1{
				out <- operator 
			}
			return
	}else if operator.emitter == id{
		fmt.Println("not found")
		return
	} else{
		out <- operator
	}
	
}

func main() {
	const channelr int = 10
	emitter := 5
	receiver:=5
	message := "data"
	var chans [channelr]chan token 
	for i := range chans {
		chans[i] = make(chan token)
	}

	for i := 0; i < channelr-1; i++ {
		go Operator(chans[i], chans[i+1], i)

	}
	go Operator(chans[channelr-1], chans[0], 9)
	chans[emitter+1] <- token{emitter, receiver, message}
	<-chans[emitter]
}