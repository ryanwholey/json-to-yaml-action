package main

import (
	"github.com/ghodss/yaml"
	"github.com/sethvargo/go-githubactions"
)

func main() {
	b, err := yaml.JSONToYAML([]byte(githubactions.GetInput("json")))
	if err != nil {
		githubactions.Fatalf("Error: %s", err)
	}

	githubactions.SetOutput("yaml", string(b))
}
