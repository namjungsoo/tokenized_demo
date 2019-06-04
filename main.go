package main

/*
comment
string
format
json
request

go get github.com/tokenized/specification/dist/golang/protocol
*/

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/tokenized/specification/dist/golang/protocol"
)

func getTx(txid string) string {
	api_key := "5SjuhZKrguAho2uPNvAmVB6TSeLi6bmD11zuGXbjw5GBcgjQhsCa5V9DHkYwfhHRYx"
	url := "https://api.bitindex.network/api/v3/main/tx/" + txid

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("api_key", api_key)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))
	var jsonmap map[string]interface{}
	if err := json.Unmarshal(bytes, &jsonmap); err != nil {
		panic(err)
	}
	fmt.Println(jsonmap["vout"])

	voutmap := jsonmap["vout"].([]interface{})
	voutmapstr := voutmap[2].(map[string]interface{})
	fmt.Println(voutmapstr)

	scriptPubKey := voutmapstr["scriptPubKey"].(map[string]interface{})
	hex := scriptPubKey["hex"]
	fmt.Println(hex)
	// str := string(bytes)
	// fmt.Println(str)

	return hex.(string)
}

func getRawTx(txid string) string {
	// bitindex.network api key
	api_key := "5SjuhZKrguAho2uPNvAmVB6TSeLi6bmD11zuGXbjw5GBcgjQhsCa5V9DHkYwfhHRYx"
	url := "https://api.bitindex.network/api/v3/main/rawtx/" + txid

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("api_key", api_key)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	bytes, _ := ioutil.ReadAll(res.Body)
	var strmap map[string]string
	if err := json.Unmarshal(bytes, &strmap); err != nil {
		panic(err)
	}

	str := strmap["rawtx"]
	//fmt.Println(str)

	return str
}

func getBitDB() []interface{} {
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
		panic(err)
	}

	bytes, _ := ioutil.ReadAll(res.Body)

	var dat map[string][]interface{}
	fmt.Println(reflect.TypeOf(dat))
	if err := json.Unmarshal(bytes, &dat); err != nil {
		panic(err)
	}
	cslice := dat["c"]
	return cslice
}

func main() {

	//fmt.Println(res.Body)

	/*
		res.Body를 ReadAll 해야함
	*/
	//str := string(bytes) //바이트를 문자열로
	//fmt.Println(str)

	// key: string
	// value: []interface{}

	cslice := getBitDB()

	for i, tx := range cslice {
		/*
			항상 type casting 해야 함
		*/
		tx2 := tx.(map[string]interface{})
		tx3 := tx2["tx"].(map[string]interface{})
		txid := tx3["h"].(string)
		fmt.Println(txid)

		//txid를 구하였으니 이걸 가지고 bitindex.network에 api 호출하여 rawtx를 얻자.
		// rawtx := getRawTx(txid)
		// fmt.Println(rawtx)

		hexstr := getTx(txid)

		//fmt.Println(txstr)

		hexbytes, err := hex.DecodeString(hexstr)
		fmt.Println(hexbytes)

		ret, err := protocol.Deserialize(hexbytes, true)
		if err != nil {
			fmt.Println("protocol error", err)
		}
		fmt.Println(ret)
		break

		// str := "6a0e746573742e746f6b656e697a65644cb70043310c44756d6d7920436f75706f6e0120000000cefc3ac1bc4798477e9a1f6302c615777e6e73b8de52e205ee4f9fe7500e3c130000000000000000000000000000804104eaa315000d44756d6d79204c696d697465644300000000000000000000000000000000000000000011737570706f72744064756d6d792e636f6d000000000008009249249249249248e803000000000000000100000000000000000000c6592dae734834382c3b0cc3e9c3973d08c0888200000000"
		// bytes := []byte(str)
		// fmt.Println(bytes)

		// ret, err := protocol.Deserialize(bytes, true)
		// if err != nil {
		// 	fmt.Println("protocol error", err)
		// }
		// fmt.Println(ret)

		break

		out := tx2["out"].([]interface{})

		for j, out2 := range out {
			out3 := out2.(map[string]interface{})

			if out3["s1"] != nil && strings.Compare(out3["s1"].(string), "test.tokenized") == 0 {
				fmt.Printf("i=%d\n", i)
				fmt.Printf("j=%d\n", j)
				//fmt.Println("test")

				// b2: base64 encoded
				// 결론적으로 b2 s2는 같은 내용이다
				// 어느것으로 parsing 하느냐에 달려 있다

				// fmt.Println(reflect.TypeOf(out3["b2"]))
				// fmt.Println(out3["b2"])
				dec, _ := base64.StdEncoding.DecodeString(out3["b2"].(string))
				fmt.Println(reflect.TypeOf(dec))
				fmt.Println(dec)

				// fmt.Println(reflect.TypeOf(out3["s2"]))
				fmt.Println(out3["s2"])

				// defer func() {
				// 	fmt.Println("defer called")
				// 	if err := recover(); err != nil {
				// 		fmt.Println("recover called")
				// 		fmt.Println(err)
				// 	}
				// }()
				msg, err := protocol.Deserialize(dec, true)
				if err == nil {
					fmt.Println(msg.String(), msg.Type(), msg.Validate())
				} else {
					fmt.Println("error")
					fmt.Println(err)
				}

				//break
				// contractOffer := new(protocol.ContractOffer)
				// contractOffer.write(dec)

				os.Exit(0)
			}
		}
	}
}
