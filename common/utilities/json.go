package utilities

import "encoding/json"

func Unmarshal[T any](object interface{}, data *T) {

	jason, _ := json.Marshal(object)

	json.Unmarshal([]byte(jason), &data)
}
