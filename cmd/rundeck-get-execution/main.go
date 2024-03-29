package main

import (
	"fmt"
	"os"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v19"
	"github.com/olekukonko/tablewriter"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	id = kingpin.Arg("execution_id", "the id of the execution").Required().String()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()
	data, err := client.GetExecution(*id)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		headers := []string{
			"ID",
			"User",
			"Status",
			"Start Date",
			"End Date",
		}
		table.SetHeader(headers)
		table.Append([]string{data.ID, data.User, data.Status, data.DateStarted, data.DateEnded})
		table.Render()
	}
}
