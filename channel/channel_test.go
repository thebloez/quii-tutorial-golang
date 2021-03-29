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
