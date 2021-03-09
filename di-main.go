package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	_ "quii-tutorial/dependency-injection"
)

func main() {
	Greet(os.Stdout, "Elodie")
	http.ListenAndServe(":5000", http.HandlerFunc(myGreeterHandler))

}

func myGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
