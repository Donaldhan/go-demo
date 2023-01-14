package mod

/**
第三方工具包测试
*/

import (
	"log"

	// "github.com/labstack/echo"
	jsoniter "github.com/json-iterator/go"
	// "github.com/samber/lo"
	// lop "github.com/samber/lo/parallel"
)

func init() {
	log.Println("==============packagex mod package init")
}

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

/**
 * mod 第三方引用包测试
 */
// https://github.com/json-iterator/go
func JsonTest() {
	//转换对象为json字符串
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := jsoniter.Marshal(group)
	if err == nil {
		log.Println("JsonTest group:", string(b))
	} else {
		log.Println("JsonTest err:", err.Error())
	}
	// 获取JSON对象
	val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	result := jsoniter.Get(val, "Colors", 0).ToString()
	log.Println("JsonTest result:", result)
}
