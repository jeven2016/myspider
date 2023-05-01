package utils

import (
	"encoding/json"
	"errors"
	"github.com/duke-git/lancet/v2/convertor"
	"net/url"
	"reflect"
	"strings"
)

func ConvertToMap(obj any) (map[string]any, error) {
	return convertor.StructToMap(obj)
}

func ConvertToMaps(objs []any) ([]map[string]any, error) {
	if objs == nil {
		return nil, nil
	}

	//convert into stream messages
	var catalogMap []map[string]any
	for _, msg := range objs {
		scm, convertErr := ConvertToMap(msg)
		if convertErr != nil {
			return nil, convertErr
		}
		catalogMap = append(catalogMap, scm)
	}
	return catalogMap, nil
}

func ConvertMapToStruct(dst any, source map[string]any) error {
	dstValue := reflect.ValueOf(dst)
	if dstValue.Kind() != reflect.Ptr {
		return errors.New("the dst should be a pointer")
	}
	jsonData, err := json.Marshal(source)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, dst)
	return err
}

func CheckHttpUrl(url string) bool {
	lowerUrl := strings.ToLower(url)

	if strings.HasPrefix(lowerUrl, "http://") || strings.HasPrefix(lowerUrl, "https://") {
		return true
	}
	return false
}

func GetBaseUrl(link string) (string, error) {
	var baseUrl string
	//if url doesn't start with http:// or https://, prefix the homeUrl
	if CheckHttpUrl(link) {
		validUrl, err := url.Parse(link)
		if err != nil {
			return baseUrl, err
		}
		baseUrl = validUrl.Scheme + "://" + validUrl.Host
		return baseUrl, err
	}
	return baseUrl, nil
}
