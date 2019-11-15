package main

import (
	"fmt"
	"os"

	"github.com/jiangew/grpc-todo-rest-microservice/pkg/cmd/server"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
