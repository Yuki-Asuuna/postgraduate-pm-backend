package upload_image

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"postgraduate-pm-backend/model"
)

const (
	url       = "http://124.221.197.218:80/api/1/upload/"
	key       = "a6f6c396a4fe645bc1a0ce16f14287df"
	nameField = "source"
)

func GetImageUrl(f *multipart.FileHeader) (string, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	formFile, err := writer.CreateFormFile(nameField, f.Filename)
	if err != nil {
		return "", err
	}

	// f.Open() 得到图片的文件对象
	img, err := f.Open()
	if err != nil {
		return "", err
	}

	_, err = io.Copy(formFile, img)
	if err != nil {
		return "", err
	}

	if err = img.Close(); err != nil {
		return "", err
	}
	if err = writer.Close(); err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	query_params := req.URL.Query()
	query_params.Add("key", key)
	query_params.Add("format", "json")
	req.URL.RawQuery = query_params.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(content))
	obj := &model.GetImageUrlResponse{}
	err = json.Unmarshal(content, obj)
	if err != nil {
		return "", err
	}
	if obj.StatusCode == http.StatusOK {
		return obj.Image.Url, nil
	}
	return "", nil
}
