package main

import "api-temp_conv/server"

func main()  {
	server := server.NewServer()

	server.Run()
}