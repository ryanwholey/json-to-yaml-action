package main

import (
	"encoding/json"
	"fmt"
	"strings"

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

	merged = append(first, second...)

	b, err := json.Marshal(merged)
	if err != nil {
		githubactions.Fatalf("Error: %s", err)
	}

	str := string(b)

	fmt.Println("should set sensitive", githubactions.GetInput("sensitive"), inputs.Bool(githubactions.GetInput("sensitive")), strings.EqualFold(githubactions.GetInput("sensitive"), "true"))
	if strings.EqualFold(githubactions.GetInput("sensitive"), "true") {
		fmt.Println("masking....")
		githubactions.AddMask(str)
	}

	githubactions.SetOutput("result", str)
}
