package main

import (
	"context"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
)

var link = "mysql:root:JCPHqknyy8ATR5ME@tcp(192.168.10.47:3306)/oldme?loc=Local"

func main() {
	var ctx = context.Background()
	var input string
	fmt.Scan(&input)
	get(ctx)
	fmt.Scan(&input)
	get(ctx)
}

func get(ctx context.Context) {
	// 这个db包含了driver和core
	db, err := gdb.New(gdb.ConfigNode{
		Link:  link,
		Debug: true,
	})
	if err != nil {
		panic(err)
	}
	db2 := db.Ctx(ctx).Model("article")
	db2 = db2.Where("article.id=?", "2")
	//db2 = db2.Where("author=?", "half")
	//db2 = db2.LeftJoin("article_grp", "article.grp_id=article_grp.id")
	data, _ := db2.All()
	fmt.Println(data)
}
