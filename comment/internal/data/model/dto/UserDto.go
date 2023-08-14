// Package dto
// @Description
// @Author  Ymri  2023/8/13 18:52
// @Update 2023/8/13 18:52
package dto

import (
	model "comment/internal/data/model"
)

type UserDto struct {
	model.User
	IsFollow bool `gorm:"-" json:"is_follow"`
}
