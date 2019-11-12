package main

import (
	"os"

	flag "github.com/spf13/pflag"
	"github.com/user/HW_Week9/service"
)

const (
	//PORT port number.
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}
	server := service.NewServer()
	server.Run(":" + port)
}
