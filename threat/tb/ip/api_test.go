package ip

import (
	"encoding/json"
	"fmt"
	"testing"

	"tb.humanrisk.cn/threat/tb/dict"
)

// TestQuery method
func TestQuery(t *testing.T) {
	API := New()
	res, err := API.Query(&QueryRequest{
		Resource: "158.247.226.121",
		Exclude:  dict.ExcludeIntelligences,
		Lang:     dict.LangEn,
	})

	if err != nil {
		fmt.Println("ip query test is error: ", err.Error())
	} else {
		if b, err := json.Marshal(res); err == nil {
			fmt.Printf("\nip query test response: %s", b)
		}
	}
}

// TestAdvQuery method
func TestAdvQuery(t *testing.T) {
	API := New()
	res, err := API.AdvQuery(&QueryRequest{
		Resource: "158.247.226.121",
		Exclude:  dict.ExcludeIntelligences,
		Lang:     dict.LangEn,
	})

	if err != nil {
		fmt.Println("ip query test is error: ", err.Error())
	} else {
		if b, err := json.Marshal(res); err == nil {
			fmt.Printf("\nip query test response: %s", b)
		}
	}
}
