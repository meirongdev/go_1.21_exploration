package main

import (
	"flag"
	"fmt"
)

// go run main.go -log
func main() {
	flag.BoolFunc("log", "logs a dummy message", func(s string) error {
		fmt.Println("dummy message:", s)
		return nil
	})
	flag.Parse() // dummy message: true

	// fs := flag.NewFlagSet("ExampleBoolFunc", flag.ContinueOnError)
	// fs.SetOutput(os.Stdout)

	// fs.BoolFunc("log", "logs a dummy message", func(s string) error {
	// 	fmt.Println("dummy message:", s)
	// 	return nil
	// })
	// fs.Parse([]string{"-log"})
	// fs.Parse([]string{"-log=0"})
}
