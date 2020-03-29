package converter

import (
	"encoding/json"
	"github.com/siller174/meetingHelper/pkg/logger"
)

func StructToJsonString(v interface{}) (string, error) {
	bytes, err := StructToJsonByte(v)
	return string(bytes), err
}

func StructToJsonByte(v interface{}) ([]byte, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		logger.Error("Can not convert %v to json")
	}
	return bytes, err
}
