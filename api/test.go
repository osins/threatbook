package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// NewWebAPITest method
func NewWebAPITest(conf *WebAPITestConfig) WebAPITest {
	return &webAPITest{
		Conf: conf,
	}
}

// WebAPITestConfig type
type WebAPITestConfig struct {
	T           *testing.T
	Description string `json:"description"`
	Method      string `json:"method"`
	Route       string `json:"route"`
	Error       bool   `json:"error"`
	Code        int    `json:"code"`
	Body        string `json:"body"`
}

// WebAPITest interface
type WebAPITest interface {
	Run(req *http.Request) (resp []byte, err error)
}

// webAPITest type
type webAPITest struct {
	Conf *WebAPITestConfig
}

//  inner method
func (s *webAPITest) Run(req *http.Request) (resp []byte, err error) {
	app := fiber.New()

	RouteInit(app)

	fmt.Println("start test: ", req.Method)

	res, err := app.Test(req, -1)
	if err != nil {
		fmt.Println("fiber test error: ", err.Error())
		return nil, err
	}

	assert.Equalf(s.Conf.T, s.Conf.Code, res.StatusCode, s.Conf.Description)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(s.Conf.T, err, s.Conf.Description)
	if err != nil {
		return nil, err
	}

	fmt.Println("fiber test response: ", string(body))

	if s.Conf.Error {
		return nil, err
	}

	reqMap := hashmap.New()
	err = reqMap.FromJSON([]byte(s.Conf.Body))
	assert.Nilf(s.Conf.T, err, s.Conf.Description)
	if err != nil {
		return nil, err
	}
	fmt.Println("test conf body to map: ", reqMap) // HashMap map[b:2 a:1]

	resMap := hashmap.New()
	err = resMap.FromJSON(body)
	assert.Nilf(s.Conf.T, err, s.Conf.Description)
	if err != nil {
		return nil, err
	}
	fmt.Println("test result to map: ", resMap) // HashMap map[b:2 a:1]

	reqCode, _ := reqMap.Get("reponse_code")
	resCode, _ := resMap.Get("reponse_code")

	reqMsg, _ := reqMap.Get("verbose_msg")
	resMsg, _ := resMap.Get("verbose_msg")

	assert.Equalf(s.Conf.T, reqCode, resCode, s.Conf.Description)
	assert.Equalf(s.Conf.T, reqMsg, resMsg, fmt.Sprintf("%s, error: %s", s.Conf.Description, resMsg))

	fmt.Println(fmt.Sprintf("\n\ntest %s result: ok\n\n", s.Conf.Description))
	fmt.Println(fmt.Sprintf("test complete, response: %s, error: %s\n\n", body, err))

	return body, err
}
