package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Ob struct {
	X string `json:"key"`
	Y int    `json:"value"`
}

func encode(i interface{}) []byte {
	data, err := json.Marshal(i)
	if err != nil {
		log.Fatalf("failed to encode: %s", data)
	}
	return data
}

func decode(data string) interface{} {
	var i interface{}
	err := json.Unmarshal([]byte(data), &i)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	return i
}

func decodeM(data string) map[string][]string {
	var m map[string][]string
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}

func main() {
	fmt.Println(string(encode("yeah test")))
	fmt.Println(encode("yeah"))

	o := &Ob{"test", 1}
	fmt.Println(string(encode(o)))

	fmt.Println(decode("100"))

	req := `{"urls": ["a", "b"]}`
	// fmt.Println(decode(req))
	// fmt.Printf("%T\n", decode(req))
	// m, _ := decode(req).(map[string]interface{})
	// i, _ := m["urls"].([]interface{})
	// urls := make([]string, len(i))
	// for k, v := range i {
	// 	urls[k], _ = v.(string)
	// }
	// for _, v := range urls {
	// 	fmt.Println(v)
	// }

	fmt.Println(decodeM(req))

	// s := []int{1, 2}
	// s = append(s, s...)
	// fmt.Println(s)
}
