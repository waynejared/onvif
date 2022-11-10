package main //package is main so it can be run with "go run unittests/GlobalVariables.go" from terminal

import (
	"fmt"
	"os"
)
import _ "github.com/joho/godotenv/autoload"

func main() {
	ip := os.Getenv("IP_ADRESS") //this is how you reference variables in .env
	user := os.Getenv("USERNAME")
	pass := os.Getenv("PASSWORD")

	fmt.Println(ip)
	fmt.Println(user)
	fmt.Println(pass)
}
