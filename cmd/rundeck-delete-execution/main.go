package main

import (
	"fmt"
	"os"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v19"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	id = kingpin.Arg("id", "id to delete").Required().String()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()
	err := client.DeleteExecution(*id)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
