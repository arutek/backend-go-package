package network

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func ApiGetData(urlString string, query map[string]string) ([]byte, error) {
	queryParams := url.Values{}
	for key, val := range query {
		queryParams.Add(key, val)
	}
	urlString += "?" + queryParams.Encode()
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return resBody, nil
}
func ApiPostData(urlString string, payload interface{}, query map[string]string) ([]byte, error) {
	queryParams := url.Values{}
	for key, val := range query {
		queryParams.Add(key, val)
	}
	urlString += "?" + queryParams.Encode()
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, urlString, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return resBody, nil
}
func ApiPostForm(urlString string, formData url.Values, query map[string]string) ([]byte, error) {
	queryParams := url.Values{}
	for key, val := range query {
		queryParams.Add(key, val)
	}
	urlString += "?" + queryParams.Encode()
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, urlString, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "multipart/form-data;boundary=XxXxX")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return resBody, nil
}
