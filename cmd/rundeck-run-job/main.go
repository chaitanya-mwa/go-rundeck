package main

import (
	"fmt"
	"os"
	"strings"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v19"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	runAs      = kingpin.Flag("runas", "user to run as").String()
	nodeFilter = kingpin.Flag("nodeFilter", "node filter to use").String()
	logLevel   = kingpin.Flag("logLevel", "log level to run the job").Enum("DEBUG", "VERBOSE", "INFO", "WARN", "ERROR")
	jobID      = kingpin.Arg("jobId", "Job ID to run").Required().String()
	argString  = kingpin.Arg("argString", "arguments to pass to job").Strings()
)

func main() {
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()
	jobopts := rundeck.RunOptions{
		AsUser:    *runAs,
		LogLevel:  *logLevel,
		Filter:    *nodeFilter,
		Arguments: strings.Join(*argString, " "),
	}
	res, err := client.RunJob(*jobID, jobopts)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Job %s is %s\n", res.Executions[0].ID, res.Executions[0].Status)
		os.Exit(0)
	}
}
