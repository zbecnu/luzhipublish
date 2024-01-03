package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"luzhibufa/db"
	"luzhibufa/model"
	"os"
	"path/filepath"
)

func main() {
	if err := db.InitDb(); err != nil {
		fmt.Println("数据库链接错误")
		return
	}

	filepath.Walk("./temp", func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			var scanner = bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			var s string
			for {
				read := scanner.Scan()
				if !read {
					break
				}
				s += scanner.Text()

			}
			var weixin model.Weixin
			if err := json.Unmarshal([]byte(s), &weixin); err != nil {
				fmt.Println("输入文本内容错误")
			}
			fmt.Println(weixin)
			fmt.Println(weixin.PublishPage)
			//解析为first
			var first model.First
			if err := json.Unmarshal([]byte(weixin.PublishPage), &first); err != nil {
				fmt.Println("输入文本内容错误")
			}

			if len(first.PublishList) > 0 {
				for _, publishInfo := range first.PublishList {
					var x model.PublisInfoEle
					if err := json.Unmarshal([]byte(publishInfo.PublishInfo), &x); err != nil {
						fmt.Println("输入文本内容错误")
					}
					for _, e := range x.Appmsgex {
						db.Db.Create(model.LuzhiPublish{
							Cover:     e.Cover,
							Title:     e.Title,
							Url:       e.Link,
							CreatedAt: e.UpdateTime,
						})
					}
				}
			}
		}
		return nil
	})
}
