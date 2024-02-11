package helper

import (
	"encoding/json"
	"errors"
	"fmt"
)

type JSON json.RawMessage

func ScanJSON(value interface{}) (json.RawMessage, error) {
	bytes, ok := value.([]byte)
	if !ok {
		return nil, errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
