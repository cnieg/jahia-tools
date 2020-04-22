package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type jahiaConnectInfo struct {
	url      string
	siteKey  string
	user     string
	password string
}

func get(connectInfo jahiaConnectInfo, url string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", connectInfo.url + url, nil)
	req.SetBasicAuth(connectInfo.user, connectInfo.password)

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(parseBody(resp).String())
	}
	defer resp.Body.Close()

	return parseJsonBody(resp)
}

func post(connectInfo jahiaConnectInfo, url string) (map[string]interface{}, error) {
	req, err := http.NewRequest("POST", connectInfo.url + url, nil)
	req.SetBasicAuth(connectInfo.user, connectInfo.password)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(parseBody(resp).String())
	}
	defer resp.Body.Close()

	return parseJsonBody(resp)
}

func downloadFile(connectInfo jahiaConnectInfo, filepath string, url string) error {
	req, err := http.NewRequest("GET", connectInfo.url + url, nil)
	req.SetBasicAuth(connectInfo.user, connectInfo.password)

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(parseBody(resp).String())
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func postFile(connectInfo jahiaConnectInfo, url string, params map[string]string, paramName, filePath string) (map[string]interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(filePath))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", connectInfo.url + url, body)
	req.SetBasicAuth(connectInfo.user, connectInfo.password)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return parseJsonBody(resp)
}

func parseJsonBody(resp *http.Response) (map[string]interface{}, error) {
	var target map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&target)
	return target, err
}

func parseBody(resp *http.Response) *bytes.Buffer {
	body := &bytes.Buffer{}
	_, err := body.ReadFrom(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	return body
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}