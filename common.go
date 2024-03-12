package bpx_api_sdk_go

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"github.com/ThreeAndTwo/bpx-api-sdk-go/constants"
	"github.com/ThreeAndTwo/bpx-api-sdk-go/types"
	"reflect"
	"sort"
	"strings"
	"time"
)

func Sign(
	privateKey *ed25519.PrivateKey,
	instruction constants.InstructionTy,
	queryParams, bodyParams map[string]interface{}) types.Signature {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	allParams := make(map[string]interface{})
	for k, v := range queryParams {
		allParams[k] = v
	}
	for k, v := range bodyParams {
		allParams[k] = v
	}

	allParams["timestamp"] = fmt.Sprintf("%d", timestamp)
	allParams["window"] = fmt.Sprintf("%d", constants.DefaultWindow)

	keys := make([]string, 0, len(allParams))
	for k := range allParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var message string
	message += fmt.Sprintf("instruction=%s", instruction)
	for _, k := range keys {
		message += fmt.Sprintf("&%s=%v", k, allParams[k])
	}

	signature := ed25519.Sign(*privateKey, []byte(message))
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)
	return types.Signature{
		Base64: signatureBase64,
		Bytes:  signature,
		Ts:     fmt.Sprintf("%d", timestamp),
	}
}

func SetHeaders(apiKey string, signature types.Signature) map[string]string {
	headers := make(map[string]string)
	headers["X-API-Key"] = apiKey
	headers["X-Timestamp"] = signature.Ts
	headers["X-Signature"] = signature.Base64
	headers["X-Window"] = fmt.Sprintf("%d", constants.DefaultWindow)
	return headers
}

func formatLimit(limit int64) int64 {
	switch {
	case limit < 0:
		return constants.DefaultLimit
	case limit > constants.MaxLimit:
		return constants.MaxLimit
	default:
		return limit
	}
}

func formatOffset(offset int64) int64 {
	switch {
	case offset < 0:
		return constants.DefaultOffset
	default:
		return offset
	}
}

func StructToMap(data interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, received %v", v.Kind())
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag.Get("json")

		if tag == "" || tag == "-" {
			continue
		}

		if commaIdx := strings.Index(tag, ","); commaIdx > 0 {
			tag = tag[:commaIdx]
		}

		if !isEmptyValue(field) {
			result[tag] = field.Interface()
		}
	}

	return result, nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
