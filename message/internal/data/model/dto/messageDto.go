// Package dto
// @Description
// @Author  Ymri  2023/8/13 21:41
// @Update 2023/8/13 21:41
package dto

import "message/internal/data/model"

type MessageDto struct {
	model.Message
	CreateTime string `gorm:"-" json:"create_time"`
}
