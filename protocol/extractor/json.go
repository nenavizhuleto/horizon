package extractor

import (
	"encoding/json"
)

func JSON[T any](from any) (T, error) {
	var ret T
	v, err := json.Marshal(from)
	if err != nil {
		return ret, err
	}

	if err := json.Unmarshal(v, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}
