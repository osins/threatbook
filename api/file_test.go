package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	api_route "tb.humanrisk.cn/api/dict"
	api_root "tb.humanrisk.cn/threat/tb"
	API_dict "tb.humanrisk.cn/threat/tb/dict"
)

// TestFileUpload method
func TestFileUpload(t *testing.T) {
	conf := &WebAPITestConfig{
		T:           t,
		Description: "file upload",
		Method:      string(API_dict.RestMethodPost),
		Route:       string(api_route.RouteV1FileUpload),
		Error:       false,
		Code:        200,
		Body:        `{"response_code":0,"verbose_msg":"成功"}`,
	}

	wat := NewWebAPITest(conf)

	filename := "D:/temps/test.docx"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(filename)
	defer file.Close()

	part1, errFile1 := writer.CreateFormFile("file", filepath.Base(filename))
	assert.Nilf(t, errFile1, conf.Description)
	_, errFile1 = io.Copy(part1, file)
	assert.Nilf(t, errFile1, conf.Description)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	_ = writer.WriteField("sandbox_type", "")
	_ = writer.WriteField("run_time", "")
	err := writer.Close()
	assert.Nilf(t, err, conf.Description)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest(
		conf.Method,
		conf.Route,
		payload,
	)

	req.Header.Set("Content-Type", writer.FormDataContentType())

	assert.Nilf(t, err, conf.Description)

	var res []byte
	res, err = wat.Run(req)
	assert.Nilf(t, err, conf.Description)
	if err != nil {
		fmt.Println(err.Error())
	}

	testFileReport(res, t)
	testFileReportMultiengines(res, t)
}

func GetSha256(body []byte, t *testing.T) interface{} {
	fmt.Println("file report input body: ", string(body))

	res := &api_root.Response{}
	err := json.Unmarshal(body, res)
	assert.Nilf(t, err, "test input res error.")
	if err != nil {
		fmt.Println("body to map error: ", err)
		fmt.Println("body: ", body)
		return ""
	}

	fmt.Println("response data: ", res)

	sha256, _ := res.Data.(map[string]interface{})["sha256"]
	fmt.Println("body: ", string(body))
	fmt.Println("sha256: ", sha256)

	return sha256
}

// testFileReport method
func testFileReport(body []byte, t *testing.T) {
	sha256 := GetSha256(body, t)
	conf := &WebAPITestConfig{
		T:           t,
		Description: "file report",
		Method:      string(API_dict.RestMethodGet),
		Route:       fmt.Sprintf("%s?sha256=%v&lang=en", api_route.RouteV1FileReportMultiengines, sha256),
		Error:       false,
		Code:        200,
		Body:        `{"response_code":0,"verbose_msg":"成功"}`,
	}

	fmt.Println("conf: ", conf)

	wat := NewWebAPITest(conf)

	req, err := http.NewRequest(
		conf.Method,
		conf.Route,
		nil,
	)

	assert.Nilf(t, err, conf.Description)

	wat.Run(req)
}

// testFileReportMultiengines method
func testFileReportMultiengines(body []byte, t *testing.T) {
	sha256 := GetSha256(body, t)
	conf := &WebAPITestConfig{
		T:           t,
		Description: "file report multiengines",
		Method:      string(API_dict.RestMethodGet),
		Route:       fmt.Sprintf("%s?sha256=%v&lang=en", api_route.RouteV1FileReportMultiengines, sha256),
		Error:       false,
		Code:        200,
		Body:        `{"response_code":0,"verbose_msg":"成功"}`,
	}

	fmt.Println("conf: ", conf)

	wat := NewWebAPITest(conf)

	req, err := http.NewRequest(
		conf.Method,
		conf.Route,
		nil,
	)

	assert.Nilf(t, err, conf.Description)

	wat.Run(req)
}
