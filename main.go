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

	flag.Parse()

	// fmt.Printf("%v", os.Args)
	// fmt.Printf(*svr)

	bufrdr := bufio.NewReader(os.Stdin)
	bufrdr.ReadLine()
	// processIssue(csvrdr, printJSON)

	processIssue(bufrdr, submitIssue, *user, *pass, *svr, *prj)

}

type procFunc func(i []string, user string, pass string, srv string, prj string) error

func processIssue(r io.Reader, process procFunc, user string, pass string, svr string, prj string) error {

	csvrdr := csv.NewReader(r)

	for {
		record, err := csvrdr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		err = process(record, user, pass, svr, prj)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func printCSV(issue []string, user string, pass string, svr string, prj string) error {
	fmt.Println(issue)
	return nil
}
func printJSON(issue []string, user string, pass string, svr string, prj string) error {
	i := newIssue(issue, prj)
	fmt.Println(string(toJSON(i)))
	return nil
}
