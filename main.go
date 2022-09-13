package main

import "github.com/aniruddha2000/kitly/pkg/server"

func main()  {
	server := server.NewServer()

	server.Initialize()
	server.Run("8989")
}
