package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(5 * time.Millisecond)
		channel <- "Ryan"
	}()

	data := <-channel
	fmt.Println(data)
}

func gimmeReturn(channel chan string, string2 string) {
	time.Sleep(5 * time.Millisecond)
	channel <- string2
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go gimmeReturn(channel, "Ryan")
	go gimmeReturn(channel, "Dewi")
	go gimmeReturn(channel, "Kanaya")

	fmt.Println(<-channel)
}
