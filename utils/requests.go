package utils

import (
	"io"
	"io/ioutil"
	"net/http"
)

func CreateRequest(url string, method string, body io.Reader) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func CreateGetRequest(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func CreatePostRequest(url string, contentType string, content io.Reader) (string, error) {
	res, err := http.Post(url, contentType, content)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
