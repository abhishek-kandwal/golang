package convert

import (
	"encoding/json"
)

func ByteToInterface (body []byte) (map[string]interface{} , error) {
	var err error
	data := map[string]interface{}{}
	err = json.Unmarshal(body , &data)
	if err != nil {
		return map[string]interface{}{} , err
	}
	return data,nil
}