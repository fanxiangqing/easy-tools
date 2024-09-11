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
