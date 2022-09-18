package helper

import (
	"encoding/json"
	"fmt"
)

func PrettyJson(v interface{}) string {
	jsonVal, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}

	return string(jsonVal)
}
