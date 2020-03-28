package request

import (
	"fmt"
	"github.com/gorilla/schema"
	"github.com/siller174/meetingHelper/pkg/logger"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func ReadBodyFromFormFileToFile(r *http.Request, formName string, path string) (*multipart.FileHeader, error) {
	file, handler, err := r.FormFile(formName)
	if err != nil {
		logger.Error("Could not get file from form '%s", formName)
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		logger.Error("Could not read file by %s from form '%s'", path, formName)
		return nil, err
	}
	err = ioutil.WriteFile(path, data, 0666)
	if err != nil {
		logger.Error("Could not write request body to file by %s", path)
		return nil, err
	}
	return handler, nil
}

func ReadQueryToStruct(r *http.Request, resultStruct interface{}) error {
	var decoder = schema.NewDecoder()
	err := decoder.Decode(resultStruct, r.URL.Query())
	if err != nil {
		logger.Error("Could not write query %v values in struct", r.URL.Query())
		return err
	}
	return nil
}

func ReadFileHeader(header string, handler *multipart.FileHeader) (string, error) {
	val := handler.Header.Get(header)
	if val == "" {
		return "", fmt.Errorf("%s in FileHeader is empty", header)
	}
	return val, nil
}
