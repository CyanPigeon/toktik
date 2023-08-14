// Package model
// @Description
// @Author  Ymri  2023/8/11 15:45
// @Update 2023/8/11 15:45
package model

type TokTikVideo struct {
	Video
	Author     UserDto `gorm:"ForeignKey:UserUID;AssociationForeignKey:UID"`
	IsFavorite bool    `gorm:"-" json:"is_favorite,omitempty"` // 视频是否点赞
}
