package shortsfactory

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

func LoadContract(path string) (Contract, error) {
	var value Contract
	err := decode(path, &value)
	return value, err
}

func LoadRequest(path string) (GateRequest, error) {
	var value GateRequest
	err := decode(path, &value)
	return value, err
}

func decode(path string, target any) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		return err
	}
	var extra any
	if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
		return errors.New("JSON file must contain exactly one value")
	}
	return nil
}
