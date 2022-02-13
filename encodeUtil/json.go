package encodeUtil

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func JSONMarshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func JSONMarshalString(v interface{}) (string, error) {
	res, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func JSONUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func JSONUnmarshalString(data string, v interface{}) error {
	return json.UnmarshalFromString(data, v)
}
