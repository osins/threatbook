package scene

import (
	"encoding/json"
	"fmt"
	"testing"

	"tb.humanrisk.cn/threatbook/dict"
)

func TestIpReputation(t *testing.T) {
	api := New()
	res, err := api.IpReputation(&QueryRequest{
		Resource: "158.247.226.121",
		Lang:     dict.LANG_ZH,
	})

	if err != nil {
		fmt.Println("ip reputation test is error: ", err.Error())
	} else {
		if b, err := json.Marshal(res); err == nil {
			fmt.Printf("\nip reputation test response: %s", b)
		}
	}
}

func TestDNS(t *testing.T) {
	api := New()
	res, err := api.DNS(&QueryRequest{
		Resource: "158.247.226.121",
		Lang:     dict.LANG_ZH,
	})

	if err != nil {
		fmt.Println("ip reputation test is error: ", err.Error())
	} else {
		if b, err := json.Marshal(res); err == nil {
			fmt.Printf("\nip reputation test response: %s", b)
		}
	}
}

func TestDomainContext(t *testing.T) {
	api := New()
	res, err := api.DomainContext(&QueryRequest{
		Resource: "158.247.226.121",
		Lang:     dict.LANG_ZH,
	})

	if err != nil {
		fmt.Println("ip reputation test is error: ", err.Error())
	} else {
		if b, err := json.Marshal(res); err == nil {
			fmt.Printf("\nip reputation test response: %s", b)
		}
	}
}
