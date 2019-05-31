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
	"strings"
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

	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(res.Body)

	/*
		res.Body를 ReadAll 해야함
	*/
	bytes, _ := ioutil.ReadAll(res.Body)
	//str := string(bytes) //바이트를 문자열로
	//fmt.Println(str)

	// key: string
	// value: []interface{}
	var dat map[string][]interface{}
	fmt.Println(reflect.TypeOf(dat))
	if err := json.Unmarshal(bytes, &dat); err != nil {
		panic(err)
	}
	cslice := dat["c"]

	for i, tx := range cslice {
		/*
			항상 type casting 해야 함
		*/
		tx2 := tx.(map[string]interface{})
		out := tx2["out"].([]interface{})

		for j, out2 := range out {
			out3 := out2.(map[string]interface{})

			if out3["s1"] != nil && strings.Compare(out3["s1"].(string), "test.tokenized") == 0 {
				fmt.Println("i=", i)
				fmt.Println("j=", j)
				fmt.Println("test")
				fmt.Println(out3)
			}
		}
	}
}
