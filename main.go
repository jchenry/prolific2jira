package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var user = flag.String("user", "user", "jira username")
var pass = flag.String("pass", "NotAPassword1", "jira password")
var prj = flag.String("prj", "PROJ", "jira project to add stories to")
var svr = flag.String("svr", "https://jira.example.com", "jira server url")

func main() {
	if len(os.Args) != 9 {
		flag.PrintDefaults()
		return
	}

	bufrdr := bufio.NewReader(os.Stdin)
	bufrdr.ReadLine()
	bufrdr.ReadLine()
	csvrdr := csv.NewReader(bufrdr)
	processIssue(csvrdr, printJSON)
}

func processIssue(r *csv.Reader, process func(issue []string) error) error {
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		err = process(record)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func printCSV(issue []string) error {
	fmt.Println(issue)
	return nil
}
func printJSON(issue []string) error {
	i := newIssue(issue)
	fmt.Println(string(toJSON(i)))
	return nil
}
