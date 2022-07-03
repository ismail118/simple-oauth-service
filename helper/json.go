package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&result)
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	err := encoder.Encode(&response)
	PanicIfError(err)
}

func WriteToResponseBodyWithCookie(w http.ResponseWriter, c *http.Cookie, response interface{}) {
	http.SetCookie(w, c)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	err := encoder.Encode(&response)
	PanicIfError(err)
}

// JsonStrConvert
// Param jsonStr is json as string,
// Param dest can be type struct that can contain value from jsonStr,
// Return err if exist
func JsonStrConvert(jsonStr string, dest interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), dest)
	if err != nil {
		return err
	}

	return nil
}

// JsonStrToMap
// Param jsonStr is json as string,
// Return map[string]interface{} and err if exist
func JsonStrToMap(jsonStr string) (map[string]interface{}, error) {
	var result map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// ObjectToJsonStr
// Param obj is can be type struct or map[string]interface{},
// Return json as string and err if exist
func ObjectToJsonStr(obj interface{}) (string, error) {
	jsonByte, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(jsonByte), nil
}
