package file

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/mitchellh/mapstructure"
	"tb.humanrisk.cn/threatbook/dict"
)

func TestUpload(t *testing.T) {
	f := &dict.FileType{
		Key:      "file",
		Filename: "D:\\temps\\test.docx",
	}

	wordBytes, err := ioutil.ReadFile(f.Filename)
	if err != nil {
		t.Error(err)

		return
	}

	f.Reader = bytes.NewReader(wordBytes)

	api := New()
	req := &UploadRequest{
		File:        f,
		SandboxType: dict.SANDBOX_WIN7_SP1_ENX64_OFFICE2013,
		RunTime:     100,
	}
	res, err := api.Upload(req)

	if err != nil {
		t.Error("file upload test is error: ", err.Error())
		return
	}

	b, err := json.Marshal(res)
	if err == nil {
		fmt.Printf("\nfile upload test response: %s\n", b)
	}

	fmt.Println("response data: ", res.Data)
	fmt.Println("response data kind: ", reflect.TypeOf(t).Kind())

	var d UploadResponse
	err = mapstructure.Decode(res.Data, &d)
	if err != nil {
		t.Error("upload file response data is error!")
		t.Failed()
		return
	}

	fmt.Println("sha256: ", d.Sha256)

	res, err = api.Report(&ReportRequest{
		SandboxType: req.SandboxType,
		Sha256:      d.Sha256,
	})

	if err != nil {
		t.Error("file report test is error: ", err.Error())
		return
	}

	b, err = json.Marshal(res)
	if err == nil {
		fmt.Printf("\nfile report test response: %s\n", b)
	}

	res, err = api.ReportMultiengines(&ReportRequest{
		SandboxType: req.SandboxType,
		Sha256:      d.Sha256,
	})

	if err != nil {
		t.Error("file report multiengines test is error: ", err.Error())
		return
	}

	b, err = json.Marshal(res)
	if err == nil {
		fmt.Printf("\nfile report multiengines test response: %s\n", b)
	}
}
