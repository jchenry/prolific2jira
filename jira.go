package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
)

// func init() {
// 	client =
// }

func submitIssue(i []string) error {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	issue := newIssue(i)
	issueJSON := toJSON(issue)

	uri := fmt.Sprintf("%s/rest/api/2/issue/", svr)
	req, _ := http.NewRequest("POST", uri, bytes.NewBuffer(issueJSON))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(*user, *pass)
	/*resp*/ _, err := client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
