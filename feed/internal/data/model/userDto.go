// Package model
// @Description
// @Author  Ymri  2023/8/12 00:34
// @Update 2023/8/12 00:34
package model

type UserDto struct {
	DbUser
	IsFollow bool `gorm:"-" json:"is_follow,omitempty"` // 是否关注
}
