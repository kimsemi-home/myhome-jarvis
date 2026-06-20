package financeconsent

import "encoding/json"

type testingT interface {
	Helper()
	Fatal(args ...any)
}

func jsonMarshal(value any) ([]byte, error) {
	return json.Marshal(value)
}
