package main

/*
comment
string
format
json
request
*/

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

func main() {
	query := []byte(`
{
	"v": 3,
	"q": {
		"db": ["c"],
		"find": {
			"out.b0.op": 106,
			"out.s1": "test.tokenized"
		},
		"limit": 10000000
	}
}`)

	encoded := base64.StdEncoding.EncodeToString(query)
	fmt.Println(encoded)

	url := "https://genesis.bitdb.network/q/1FnauZ9aUH2Bex6JzdcV4eNX7oLSSEbxtN/" + encoded
	fmt.Println(url)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("key", "1Azd2MmsCNpQn6jX1uwxDdUjm28VdNPD3F")
	res, err := client.Do(req)

	fmt.Println(err)
	fmt.Println(res.Body)

	bytes, _ := ioutil.ReadAll(res.Body)
	//str := string(bytes) //바이트를 문자열로
	//fmt.Println(str)

	// key: string
	// value: []interface{}
	var dat map[string][]interface{}
	if err := json.Unmarshal(bytes, &dat); err != nil {
		panic(err)
	}
	//fmt.Println(dat)

	fmt.Println(" \n \n")
	cslice := dat["c"]

	fmt.Println(reflect.TypeOf(cslice))

	for i, tx := range cslice {
		fmt.Println(i)
		fmt.Println(tx)
	}

}
