package main //package is main so it can be run with "go run unittests/GlobalVariables.go" from terminal

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ip := os.Getenv("CAMERA_IP_ADDRESS") //this is how you reference variables in .env
	user1 := os.Getenv("CAMERA_USERNAME")
	pass := os.Getenv("CAMERA_PASSWORD")

	fmt.Println(ip)
	fmt.Println(user1)
	fmt.Println(pass)
}
