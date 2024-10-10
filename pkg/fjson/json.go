package fjson

import (
	"encoding/json"
	"io"
)

func ParseBody(b io.ReadCloser, t json.Unmarshaler) error {
	var err error
	var body []byte

	body, err = io.ReadAll(b)
	if err != nil {
		return err
	}

	err = t.UnmarshalJSON(body)
	if err != nil {
		return err
	}

	return nil
}
