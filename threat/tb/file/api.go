package file

import (
	api_root "tb.humanrisk.cn/threat/tb"
	"tb.humanrisk.cn/threat/tb/dict"
)

// New method
func New() API {
	return &api{}
}

// API interface
type API interface {
	Upload(req *UploadRequest) (*api_root.Response, error)
	Report(req *ReportRequest) (*api_root.Response, error)
	ReportMultiengines(req *ReportRequest) (*api_root.Response, error)
}

// API type
type api struct {
}

//  inner method
func (s *api) Upload(req *UploadRequest) (*api_root.Response, error) {
	res, err := api_root.NewRest().Post(string(dict.APIV3FileUpload), req)

	return res, err
}

//  inner method
func (s *api) Report(req *ReportRequest) (*api_root.Response, error) {
	res, err := api_root.NewRest().Post(string(dict.APIV3FileReport), req)

	return res, err
}

//  inner method
func (s *api) ReportMultiengines(req *ReportRequest) (*api_root.Response, error) {
	res, err := api_root.NewRest().Post(string(dict.APIV3FileReportMultiengines), req)

	return res, err
}
