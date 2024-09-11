# Easy Tools

Easy Tools 提供一些基础快捷的工具包, 可以方便快速开发

## 快速使用

```sh
import github.com/fanxiangqing/easy-tools
```

```sh
go get -u github.com/fanxiangqing/easy-tools
```

## Example

beanutils 使用

```go
package main

import (
	"fmt"
	"github.com/fanxiangqing/easy-tools/easytools"
	"time"
)

type A struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type B struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func main() {
	a := A{
		ID:        1,
		Name:      "张三",
		CreatedAt: time.Now(),
	}

	fmt.Println("A: ", a)

	var b B
	if err := easytools.CopyProperties(a, &b); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("B: ", b)
}
```

json 使用

```go
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
```