package mediareadiness

import "fmt"

func payloadFor(kind string) ([]byte, error) {
	switch kind {
	case "", "empty":
		return []byte("{}"), nil
	case "fixture_query":
		return []byte(`{"query":"sample media readiness"}`), nil
	case "fixture_service_netflix":
		return []byte(`{"service":"netflix"}`), nil
	default:
		return nil, fmt.Errorf("unknown media readiness payload kind %q", kind)
	}
}
