package externalevidence

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func appendJSONLine(root string, rel string, record any) error {
	body, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("write_private_layer")
	}
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return fmt.Errorf("write_private_layer")
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		return fmt.Errorf("write_private_layer")
	}
	defer file.Close()
	if _, err := file.Write(append(body, '\n')); err != nil {
		return fmt.Errorf("write_private_layer")
	}
	return nil
}

func writePrivateFile(root string, rel string, body []byte) error {
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return fmt.Errorf("write_private_layer")
	}
	if err := os.WriteFile(path, body, 0o600); err != nil {
		return fmt.Errorf("write_private_layer")
	}
	return nil
}
