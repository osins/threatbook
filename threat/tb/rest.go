package threatbook

import (
	"fmt"
	"net/http"

	"github.com/fatih/structs"
	"github.com/go-resty/resty/v2"
	"tb.humanrisk.cn/threat/tb/dict"
)

// NewRest method
func NewRest() Rest {
	return &rest{}
}

// Rest interface
type Rest interface {
	Get(url string, data interface{}) (*Response, error)
	Post(url string, data interface{}) (*Response, error)
	Error() error
}

// rest type
type rest struct {
	req *http.Request
	res *http.Response
	err error
}

// Error inner method
func (s *rest) Error() error {
	return s.err
}

//  inner method
func (s *rest) Get(url string, data interface{}) (*Response, error) {
	return s.build(dict.RestMethodGet, url, data)
}

//  inner method
func (s *rest) Post(url string, data interface{}) (*Response, error) {
	return s.build(dict.RestMethodPost, url, data)
}

//  inner method
func (s *rest) build(method dict.MethodType, url string, data interface{}) (*Response, error) {
	fmt.Printf("\n\n###########################################################\n")
	fmt.Println("    API:", url)
	fmt.Printf("###########################################################\n")
	if !structs.IsStruct(data) {
		s.err = fmt.Errorf("data not is struct: %s", url)
		return nil, s.err
	}

	fmt.Println(fmt.Printf("request data: %v", data))

	d := structs.New(data)
	m := make(map[string]string)
	rs := make(map[string]dict.FileType)

	m["APIkey"] = "e9196f0aa4eb4d90aa3658fbc46a82e28a5c58de7c994f43a5cbc0f416328c8e"

	for _, f := range d.Fields() {
		key := f.Tag("json")

		if v, ok := f.Value().(*dict.FileType); ok {
			fmt.Println("name: ", key, ",\tvalue: ", f.Value(), ", convert to string: %b", ok)
			rs[key] = *v
			break
		}

		fmt.Println("name: ", key, ", kind: ", f.Kind(), ",\tvalue: ", f.Value())
		m[key] = fmt.Sprint(f.Value())
	}

	if m["lang"] == "" {
		m["lang"] = "zh"
	}

	client := resty.New()
	req := client.R()

	sets := map[dict.MethodType]func() (*resty.Response, error){
		dict.RestMethodGet: func() (*resty.Response, error) {
			for i, v := range m {
				fmt.Println("i: ", i, ",\tv: ", v)
				req.SetQueryParam(i, v)
			}

			req.EnableTrace().SetResult(&Response{}).ForceContentType("application/json")
			return req.Get(url)
		},
		dict.RestMethodPost: func() (*resty.Response, error) {
			for _, v := range rs {
				req.SetFileReader(v.Key, v.Filename, v.Reader)
			}

			req.SetFormData(m)
			req.EnableTrace().SetResult(&Response{}).ForceContentType("application/json")
			return req.Post(url)
		},
	}

	resp, err := sets[method]()

	if err != nil {
		s.err = err
		return nil, err
	}

	fmt.Println("rest get response: ", resp)
	return resp.Result().(*Response), nil
}
