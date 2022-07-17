package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

// SendGetHttpRequest
// Param reqUrl is destination url,
// Return body response as string and error if exist
func SendGetHttpRequest(reqUrl string) (string, error) {
	response, err := http.Get(reqUrl)
	if err != nil {
		return "", err
	}

	defer PanicIfError(response.Body.Close())

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// SendPostHttpRequest
// Param reqUrl is destination url,
// Param requestBody can be struct or map[string]string,
// Return body response as string and error if exist
func SendPostHttpRequest(reqUrl string, requestBody interface{}) (string, error) {
	reqBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	response, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// SendPostFormHttpRequest
// Param reqUrl is destination url,
// Param formDataRequest can be struct or map[string]string,
// Return body response as string and error if exist
func SendPostFormHttpRequest(reqUrl string, formDataRequest map[string][]string) (map[string]interface{}, error) {
	var formData url.Values
	var result map[string]interface{}

	for key, each := range formDataRequest {
		formData[key] = each
	}

	response, err := http.PostForm(reqUrl, formData)
	if err != nil {
		return result, err
	}

	defer PanicIfError(response.Body.Close())

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// SendCustomHttpRequest
// Param method is http.Method ex http.MethodGet, http.MethodPost etc.
// Param reqUrl is destination url,
// Param formDataRequest is request body as map[string]interface{},
// Return body response as string and error if exist
func SendCustomHttpRequest(method string, reqUrl string, requestBody map[string]interface{}) (string, error) {
	reqBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	timeout := time.Duration(15 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest(method, reqUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer PanicIfError(response.Body.Close())

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// SendMultipartFormHttpRequest
// Param reqUrl is destination url,
// Param fieldFiles is map[string]string that contain fileName as key and file directory as value,
// it would look like this map[fileName]fileDir and filed name of each file will auto generate,
// Param fieldsValue is map[string]string that contain filedName and value
// it would look like this map[filedName]value,
// Return body of request as type map[string]interface{} and error if exist
func SendMultipartFormHttpRequest(reqUrl string, fieldFiles map[string]string, fieldsValue map[string]string) (map[string]interface{}, error) {
	var number int
	var requestBody bytes.Buffer
	var result map[string]interface{}
	multiPartWriter := multipart.NewWriter(&requestBody)

	for fileName, fileDir := range fieldFiles {
		err := func() error {
			file, err := os.Open(filepath.Join(fileName, fileDir))
			if err != nil {
				return err
			}

			defer PanicIfError(file.Close())

			number++
			fieldName := fmt.Sprintf("file_field%d", number)
			fileWriter, err := multiPartWriter.CreateFormFile(fieldName, fileName)
			if err != nil {
				return err
			}

			_, err = io.Copy(fileWriter, file)
			if err != nil {
				return err
			}

			return nil
		}()

		if err != nil {
			return nil, err
		}
	}
	for field, value := range fieldsValue {
		err := func() error {
			fieldWriter, err := multiPartWriter.CreateFormField(field)
			if err != nil {
				return err
			}

			_, err = fieldWriter.Write([]byte(value))
			if err != nil {
				return err
			}
			return nil
		}()

		if err != nil {
			return nil, err
		}
	}

	PanicIfError(multiPartWriter.Close())

	request, err := http.NewRequest(http.MethodPost, reqUrl, &requestBody)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	response, err := client.Do(request)
	if err != nil && errors.Is(err, http.ErrUseLastResponse) {
		fmt.Println(err)
	}
	if err != nil {
		return nil, err
	}
	defer PanicIfError(response.Body.Close())

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
