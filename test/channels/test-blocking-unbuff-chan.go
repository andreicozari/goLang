package main

func main() {
	unBuffChan := make(chan int)

	// blocked operation / can't read any value because the channel has no values
	// You can't read from an unbuffered channel if until now nobody have pushed any value
	// A deadlock here:
	<-unBuffChan
	//fmt.Println("After trying to get a value from an unbuffered channel without being pushed any value in that channel!")
}
