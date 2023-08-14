// Package utils
// @Description
// @Author  Ymri  2023/8/12 21:42
// @Update 2023/8/12 21:42
package utils

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/log"
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

// DtoUtils dto转换工具
func DtoUtils(inputClass any, outputClass any) {
	// 把输入转成json
	retJson, err := json.Marshal(inputClass)
	if err != nil {
		log.Error("json序列化转换失败")
		panic(err)
	}
	// json转成输出
	err = json.Unmarshal(retJson, &outputClass)

	if err != nil {
		log.Error("json反序列化转换失败")
		panic(err)
	}
}
