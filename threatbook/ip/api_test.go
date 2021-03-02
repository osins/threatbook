package ip

import (
	"encoding/json"
	"fmt"
	"testing"

	"tb.humanrisk.cn/threatbook/dict"
)

func TestQuery(t *testing.T) {
	api := New()
	res, err := api.Query(&QueryRequest{
		Resource: "158.247.226.121",
		Exclude:  dict.EXCLUDE_INTELLIGENCES,
		Lang:     dict.LANG_EN,
	})

	if err != nil {
		fmt.Println("ip query test is error: ", err.Error())
	} else {
		if b, err := json.Marshal(res); err == nil {
			fmt.Printf("\nip query test response: %s", b)
		}
	}
}

func TestAdvQuery(t *testing.T) {
	api := New()
	res, err := api.AdvQuery(&QueryRequest{
		Resource: "158.247.226.121",
		Exclude:  dict.EXCLUDE_INTELLIGENCES,
		Lang:     dict.LANG_EN,
	})

	if err != nil {
		fmt.Println("ip query test is error: ", err.Error())
	} else {
		if b, err := json.Marshal(res); err == nil {
			fmt.Printf("\nip query test response: %s", b)
		}
	}
}
