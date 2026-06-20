package cicache

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func readInputs(root string) (graphFile, string, error) {
	graph, err := readGraph(root)
	if err != nil {
		return graphFile{}, "", err
	}
	workflow, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(workflowPath)))
	if err != nil {
		return graphFile{}, "", err
	}
	return graph, string(workflow), nil
}

func readGraph(root string) (graphFile, error) {
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(graphPath)))
	if err != nil {
		return graphFile{}, err
	}
	var graph graphFile
	if err := json.Unmarshal(body, &graph); err != nil {
		return graphFile{}, fmt.Errorf("%s: %w", graphPath, err)
	}
	return graph, nil
}
