package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)


func main() {
	incorrectName()
	fmt.Println("\033[0;31mred\033[0m")
	fmt.Println("\033[0;32mgreen\033[0m")
	fmt.Println("\033[0;33myellow\033[0m")
	fmt.Println("\033[0;34mblue\033[0m")
	fmt.Println("\033[0;34mblue\033[0m")

}

//func F[T any](p T) {
//	fmt.Println(T)
//}

type User struct {
	LastName  string `json:"lastname"`
	FirstName string `json:"name"`
	firstName string `json:"name"`
}

//incorrectName
func incorrectName() {
	for _, v := range os.Environ() {
		fmt.Println(v)
	}
}

func handle(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())

	if req.Method != http.MethodGet {
		return
	}

	defer cancel()
	doHandle(ctx)
}

func doHandle(ctx context.Context) {
	// TODO
}

func MyFunc() string {
	return ""
}

func _() {
	time.Parse("", "12.01.2020")
}
