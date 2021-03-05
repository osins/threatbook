package ip

import (
	"encoding/json"
	"fmt"
	"testing"
)

// TestScan method
func TestScan(t *testing.T) {
	API := New()
	res, err := API.Scan(&QueryRequest{
		URL: "https://baidu.com",
	})

	if err != nil {
		fmt.Println("ip query test is error: ", err.Error())
	} else {
		if b, err := json.Marshal(res); err == nil {
			fmt.Printf("\nip query test response: %s", b)
		}
	}
}

// TestReport method
func TestReport(t *testing.T) {
	API := New()
	res, err := API.Report(&QueryRequest{
		URL: "https://baidu.com",
	})

	if err != nil {
		fmt.Println("ip query test is error: ", err.Error())
	} else {
		if b, err := json.Marshal(res); err == nil {
			fmt.Printf("\nip query test response: %s", b)
		}
	}
}
