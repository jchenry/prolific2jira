package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
)

var client *http.Client

func init() {
	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

}

func submitIssue(i []string, user string, pass string, svr string, prj string) error {

	issue := newIssue(i, prj)
	issueJSON := toJSON(issue)

	uri := fmt.Sprintf("%s/rest/api/2/issue/", svr)
	req, _ := http.NewRequest("POST", uri, bytes.NewBuffer(issueJSON))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, pass)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	fmt.Printf(resp)

	return nil
}
