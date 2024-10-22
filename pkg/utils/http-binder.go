package utils

import (
	"encoding/json"
	"io"
)

func HttpBind(body io.ReadCloser, model json.Unmarshaler) error {
	b, err := io.ReadAll(body)

	if err != nil {
		return err
	}

	return model.UnmarshalJSON(b)
}
