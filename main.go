package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"luzhibufa/db"
	"luzhibufa/model"
	"os"
)

func main() {
	if err := db.InitDb(); err != nil {
		fmt.Println("数据库链接错误")
		return
	}
	fmt.Println("输入每次结果后，请单独输入[end],然后回车==!...")
	scanner := bufio.NewScanner(os.Stdin)
	var s string
	for scanner.Scan() {
		var input = scanner.Text()
		switch input {
		case "end":
			//
			var w model.W
			if err := json.Unmarshal([]byte(s), &w); err != nil {
				fmt.Println("输入文本内容错误")
			}
			for _, e := range w.AppMsgList {
				db.Db.Create(model.LuzhiPublish{
					Cover:     e.Cover,
					Title:     e.Title,
					Url:       e.Link,
					CreatedAt: e.UpdateTime,
				})
			}
			s = ""
			fmt.Println(w)
		case "exit":
			fmt.Println("结束。。")
			return
		default:
			s = fmt.Sprintf("%s%s", s, input)
		}

	}
}
