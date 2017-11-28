package output

import (
	"encoding/json"
	"fmt"
)

func toJson(lang interface{}) []byte {
	enc, err := json.Marshal(lang)
	if err != nil {
		fmt.Println(err.Error())
		return []byte(nil)
	}
	return enc
}