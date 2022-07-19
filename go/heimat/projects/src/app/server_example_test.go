package app_test

import (
	"fmt"
	"time"

	"sprinteins.com/go-bank/src/app"
)

func ExampleServer() {
	s := app.NewServer("0.0.0.0", "8080")
	stop := s.Start()

	time.Sleep(time.Second)

	<-stop()
	fmt.Println("Server stopped")
}
