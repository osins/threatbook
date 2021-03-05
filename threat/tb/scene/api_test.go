package scene

import (
	"encoding/json"
	"fmt"
	"testing"

	"tb.humanrisk.cn/threat/tb/dict"
)

// TestIPReputation method
func TestIPReputation(t *testing.T) {
	API := New()
	res, err := API.IPReputation(&QueryRequest{
		Resource: "158.247.226.121",
		Lang:     dict.LangZh,
	})

	if err != nil {
		fmt.Println("ip reputation test is error: ", err.Error())
	} else {
		if b, err := json.Marshal(res); err == nil {
			fmt.Printf("\nip reputation test response: %s", b)
		}
	}
}

// TestDNS method
func TestDNS(t *testing.T) {
	API := New()
	res, err := API.DNS(&QueryRequest{
		Resource: "158.247.226.121",
		Lang:     dict.LangZh,
	})

	if err != nil {
		fmt.Println("ip reputation test is error: ", err.Error())
	} else {
		if b, err := json.Marshal(res); err == nil {
			fmt.Printf("\nip reputation test response: %s", b)
		}
	}
}

// TestDomainContext method
func TestDomainContext(t *testing.T) {
	API := New()
	res, err := API.DomainContext(&QueryRequest{
		Resource: "158.247.226.121",
		Lang:     dict.LangZh,
	})

	if err != nil {
		fmt.Println("ip reputation test is error: ", err.Error())
	} else {
		if b, err := json.Marshal(res); err == nil {
			fmt.Printf("\nip reputation test response: %s", b)
		}
	}
}
