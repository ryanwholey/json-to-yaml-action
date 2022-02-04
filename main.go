package main

import (
	"encoding/json"

	inputs "github.com/bendrucker/go-githubactions-inputs"
	"github.com/sethvargo/go-githubactions"
)

func main() {
	var first []interface{}
	var second []interface{}
	var merged []interface{}

	if err := json.Unmarshal([]byte(githubactions.GetInput("first")), &first); err != nil {
		githubactions.Fatalf("Error: %s", err)
	}

	if err := json.Unmarshal([]byte(githubactions.GetInput("second")), &second); err != nil {
		githubactions.Fatalf("Error: %s", err)
	}

	merged = append(first, second)

	b, err := json.Marshal(merged)
	if err != nil {
		githubactions.Fatalf("Error: %s", err)
	}

	str := string(b)

	if inputs.Bool(githubactions.GetInput("sensitive")) {
		githubactions.AddMask(str)
	}

	githubactions.SetOutput("result", str)
}
