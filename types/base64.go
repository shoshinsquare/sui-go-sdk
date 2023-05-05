package types

import (
	"encoding/base64"
	"encoding/json"
)

type Base64Data []byte

func NewBase64Data(str string) (*Base64Data, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	b64 := Base64Data(data)
	return &b64, nil
}

func (h Base64Data) Data() []byte {
	return h
}

func (h Base64Data) Length() int {
	return len(h)
}
func (h Base64Data) String() string {
	return base64.StdEncoding.EncodeToString(h)
}

func (h Base64Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.String())
}

func (h *Base64Data) UnmarshalJSON(data []byte) error {
	str := ""
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	tmp, err := NewBase64Data(str)
	if err == nil {
		*h = *tmp
	}
	return err
}

func (h Base64Data) MarshalBCS() ([]byte, error) {
	return h, nil
}
