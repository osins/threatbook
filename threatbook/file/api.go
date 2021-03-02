package file

import (
	"tb.humanrisk.cn/threatbook"
	"tb.humanrisk.cn/threatbook/dict"
)

func New() API {
	return &api{}
}

type API interface {
	Upload(req *UploadRequest) (*threatbook.Response, error)
	Report(req *ReportRequest) (*threatbook.Response, error)
	ReportMultiengines(req *ReportRequest) (*threatbook.Response, error)
}

type api struct {
}

func (s *api) Upload(req *UploadRequest) (*threatbook.Response, error) {
	res, err := threatbook.NewRest().Post(string(dict.API_V3_FILE_UPLOAD), req)

	return res, err
}

func (s *api) Report(req *ReportRequest) (*threatbook.Response, error) {
	res, err := threatbook.NewRest().Post(string(dict.API_V3_FILE_REPORT), req)

	return res, err
}

func (s *api) ReportMultiengines(req *ReportRequest) (*threatbook.Response, error) {
	res, err := threatbook.NewRest().Post(string(dict.API_V3_FILE_REPORT_MULTIENGINES), req)

	return res, err
}
