// Package utils
// @Description
// @Author  Ymri  2023/8/12 21:42
// @Update 2023/8/12 21:42
package utils

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

// Init 初始化雪花算法
func Init(machineID int64) (err error) {
	var st time.Time
	// 格式化 1月2号下午3时4分5秒  2006年
	// 获得现在时间
	st = time.Now()
	if err != nil {
		fmt.Println(err)
		return
	}
	snowflake.Epoch = st.UnixNano() / 1e6
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

// GenID 生成 64 位的 雪花 ID
func GenID() int64 {
	return node.Generate().Int64()
}

// func main() {
// 	if err := Init(1); err != nil {
// 		fmt.Println("Init() failed, err = ", err)
// 		return
// 	}

// 	id := GenID()
// 	fmt.Println(id)
// }
