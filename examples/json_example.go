package main

import (
	"fmt"
	"github.com/fanxiangqing/easy-tools/easytools"
)

func main() {
	jsonObject := easytools.JSONObject{}
	jsonObject.Put("a", int64(1))
	jsonObject.Put("b", "2")

	fmt.Println(jsonObject.GetInt64("a"))
	fmt.Println(jsonObject.GetString("b"))

	jsonString, err := easytools.ToJsonString(jsonObject)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(jsonString)

	parseJsonObject, err := easytools.ParseJsonObject("{\"a\":\"b\"}")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(parseJsonObject)
}
