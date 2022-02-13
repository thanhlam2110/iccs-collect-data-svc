package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jmcvetta/napping"
)

func ProducerMessage(message, topic string) string {
	url := "http://13.250.2.116:1323/api/kafka/produce"
	tokenInputString := `
						{"message":"` + message + `",
						"topic":"` + topic + `"
						}`
	s := napping.Session{}
	h := &http.Header{}
	h.Set("X-Custom-Header", "")
	s.Header = h
	var jsonStr = []byte(tokenInputString)

	var data map[string]json.RawMessage
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := s.Post(url, &data, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(resp.RawText()), &result)
	return result["code"].(string)

}
