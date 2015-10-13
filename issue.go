package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type issue struct {
	Fields fields `json:"fields"`
}

type fields struct {
	Type        issuetype `json:"issuetype"`
	Project     project   `json:"project"`
	Title       string    `json:"summary"`
	Description string    `json:"description"`
	Labels      []string  `json:"labels"`
}

type issuetype struct {
	Name string `json:"name"`
}

type project struct {
	Key string `json:"key"`
}

func newIssue(f []string) *issue {
	return &issue{
		Fields: fields{
			Title: f[0],
			Type: issuetype{
				Name: toIssueType(f[1]),
			},
			Description: f[2],
			Labels:      strings.Split(f[3], ","),
			Project: project{
				Key: *prj,
			},
		},
	}
}

func toJSON(i *issue) []byte {
	b, _ := json.Marshal(i)
	return b
}

func toIssueType(t string) string {
	fmt.Println(t)
	switch strings.ToLower(t) {
	case "feature":
		return "Story"
	case "release":
		return "Story"
	case "bug":
		return "Bug"
	case "chore":
		return "Task"
	}
	return "unrecognized_type"
}
