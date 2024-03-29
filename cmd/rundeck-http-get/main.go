package main

import (
	"fmt"
	"os"
	"strings"

	rundeck "github.com/lusis/go-rundeck/pkg/rundeck.v19"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	path            = kingpin.Arg("path", "path to dump (e.g. executions/12234)").Required().String()
	queryParameters = kingpin.Flag("query_params", "key=value query parameter. specify multiple times if neccessary").Strings()
	contentType     = kingpin.Flag("content_type", "an alternate content type if neccessary").Default("application/xml").String()
)

func buildParams(p *map[string]string, value string) error {
	parts := strings.SplitN(value, "=", 2)
	if len(parts) != 2 {
		return fmt.Errorf("expected key=value got '%s'", value)
	}
	(*p)[parts[0]] = parts[1]
	return nil
}

func main() {
	myParams := make(map[string]string)
	kingpin.Parse()
	client := rundeck.NewClientFromEnv()
	var data []byte
	if contentType != nil {
		myParams["content_type"] = *contentType
	}

	for _, param := range *queryParameters {
		e := buildParams(&myParams, param)
		if e != nil {
			fmt.Printf(e.Error())
			os.Exit(1)
		}
	}
	err := client.Get(&data, *path, myParams)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", string(data))
	}
}
