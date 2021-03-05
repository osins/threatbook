package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"tb.humanrisk.cn/api/dict"
	api_root "tb.humanrisk.cn/threat/tb"
	API_dict "tb.humanrisk.cn/threat/tb/dict"
	API_file "tb.humanrisk.cn/threat/tb/file"
)

// NewFileRoute method
func NewFileRoute() Routes {
	return &file{}
}

// file type
type file struct {
}

// RouteInit inner method
func (s *file) RouteInit(app *fiber.App) error {
	app.Post(string(dict.RouteV1FileUpload), s.Upload)
	app.Get(string(dict.RouteV1FileReport), s.Report)
	app.Get(string(dict.RouteV1FileReportMultiengines), s.ReportMultiengines)
	return nil
}

// Upload inner method
func (s *file) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	uuid := uuid.New()
	tempFilename := fmt.Sprintf("D:/temps/%s.%s", uuid.String(), filepath.Ext(file.Filename))

	err = c.SaveFile(file, tempFilename)
	if err != nil {
		return err
	}

	wordBytes, err := ioutil.ReadFile(tempFilename)
	if err != nil {
		return err
	}

	req := &API_file.UploadRequest{
		File: &API_dict.FileType{
			Key:      "file",
			Filename: tempFilename,
			Reader:   bytes.NewReader(wordBytes),
		},
	}

	err = c.QueryParser(req)
	if err != nil {
		return err
	}

	var resp *api_root.Response
	resp, err = API_file.New().Upload(req)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

// Report inner method
func (s *file) Report(c *fiber.Ctx) error {
	fmt.Println("request body: ", c.Context().QueryArgs())

	req := &API_file.ReportRequest{}
	err := c.QueryParser(req)
	if err != nil {
		return err
	}

	var resp *api_root.Response
	resp, err = API_file.New().Report(req)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

// ReportMultiengines inner method
func (s *file) ReportMultiengines(c *fiber.Ctx) error {
	fmt.Println("request body: ", c.Context().QueryArgs())

	req := &API_file.ReportRequest{}
	err := c.QueryParser(req)
	if err != nil {
		return err
	}

	var resp *api_root.Response
	resp, err = API_file.New().ReportMultiengines(req)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}
