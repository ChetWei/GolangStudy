package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"xingye", "zhangbozhi"}}

	//结构体 -> json  编码过程
	//将结构体里面的字段名称与对应的json标签进行转换
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr= %s\n", jsonStr)

	//jsonstr -> 结构体 解码过程
	my_movie := Movie{}
	err = json.Unmarshal(jsonStr, &my_movie)
	if err != nil {
		fmt.Println("json ummashal", err)
		return
	}

	fmt.Printf("%v\n", my_movie)
}
