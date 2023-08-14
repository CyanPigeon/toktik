// Package model
// @Description
// @Author  Ymri  2023/8/13 00:31
// @Update 2023/8/13 00:31
package model

type VideoDto struct {
	Video
	IsFavorite bool `gorm:"-" json:"is_favorite"` // 视频是否点赞
	Author     User `gorm:"ForeignKey:UserUID;AssociationForeignKey:UID"`
}
