package utilities

import "encoding/json"

func Unmarshal[T any](object interface{}, data *T) error {

	jason, _ := json.Marshal(object)

	errors := json.Unmarshal([]byte(jason), &data)
	return errors
}
