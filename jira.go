package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
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

	fmt.Printf("* %s :", issue.Fields.Title)

	uri := fmt.Sprintf("%s/rest/api/2/issue/", svr)
	req, _ := http.NewRequest("POST", uri, bytes.NewBuffer(issueJSON))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, pass)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	r := make(map[string]interface{})

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(contents, &r)
	if err != nil {
		return err
	}

	fmt.Printf(" created as %s", r["key"])

	// printResponse(resp)

	return nil
}

func printResponse(resp *http.Response) {
	r, _ := httputil.DumpResponse(resp, true)
	fmt.Printf(string(r))
}
