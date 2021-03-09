package main

import (
	"bytes"
	"fmt"
	"time"
)

const finalWord string = "Go!"
const countDownStart int = 3

func main() {
	CountdownInt(4)
}

func CountdownInt(a int) {
	for ; a > 0; a-- {
		fmt.Printf("%d\n", a)
		time.Sleep(1 / 2 * time.Second)
	}
	fmt.Println("Go!")
}

func Countdown(out *bytes.Buffer) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, finalWord)
}
