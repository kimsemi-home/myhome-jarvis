package localfinancereadiness

import (
	"encoding/json"
	"errors"
	"io"
)

func decodeOne[T any](reader io.Reader) (T, error) {
	var value T
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&value); err != nil {
		return value, err
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return value, errors.New("JSON input must contain exactly one value")
	}
	return value, nil
}
