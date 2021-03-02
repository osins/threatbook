package threatbook

import (
	"fmt"
	"net/http"

	"github.com/fatih/structs"
	"github.com/go-resty/resty/v2"
	"tb.humanrisk.cn/threatbook/dict"
)

func NewRest() Rest {
	return &rest{}
}

type Rest interface {
	Get(url string, data interface{}) (*Response, error)
	Post(url string, data interface{}) (*Response, error)
	Error() error
}

type rest struct {
	req *http.Request
	res *http.Response
	err error
}

func (s *rest) Error() error {
	return s.err
}

func (s *rest) Get(url string, data interface{}) (*Response, error) {
	return s.build(dict.REST_METHOD_GET, url, data)
}

func (s *rest) Post(url string, data interface{}) (*Response, error) {
	return s.build(dict.REST_METHOD_POST, url, data)
}

func (s *rest) build(method dict.MethodType, url string, data interface{}) (*Response, error) {
	fmt.Printf("\n\n###########################################################\n")
	fmt.Println("    api:", url)
	fmt.Printf("###########################################################\n")
	if !structs.IsStruct(data) {
		s.err = fmt.Errorf("data not is struct.")
		return nil, s.err
	}

	fmt.Println(fmt.Printf("request data: %v", data))

	d := structs.New(data)
	m := make(map[string]string)
	rs := make(map[string]dict.FileType)

	m["apikey"] = "e9196f0aa4eb4d90aa3658fbc46a82e28a5c58de7c994f43a5cbc0f416328c8e"

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
		dict.REST_METHOD_GET: func() (*resty.Response, error) {
			for i, v := range m {
				fmt.Println("i: ", i, ",\tv: ", v)
				req.SetQueryParam(i, v)
			}

			req.EnableTrace().SetResult(&Response{}).ForceContentType("application/json")
			return req.Get(url)
		},
		dict.REST_METHOD_POST: func() (*resty.Response, error) {
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
