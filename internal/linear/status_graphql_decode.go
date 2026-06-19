package linear

import (
	"encoding/json"
	"errors"
	"fmt"
)

func decodeGraphQLPayload(httpStatus int, remaining int, payload []byte, target any) (int, int, error) {
	var envelope graphQLEnvelope
	if err := json.Unmarshal(payload, &envelope); err != nil {
		return httpStatus, remaining, err
	}
	if len(envelope.Errors) > 0 {
		return httpStatus, remaining, fmt.Errorf("Linear GraphQL error: %s", envelope.Errors[0].Message)
	}
	if len(envelope.Data) == 0 || string(envelope.Data) == "null" {
		return httpStatus, remaining, errors.New("Linear response did not include data")
	}
	if target != nil {
		if err := json.Unmarshal(envelope.Data, target); err != nil {
			return httpStatus, remaining, err
		}
	}
	return httpStatus, remaining, nil
}
