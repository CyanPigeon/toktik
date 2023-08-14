// Package dto
// @Description
// @Author  Ymri  2023/8/13 14:42
// @Update 2023/8/13 14:42
package dto

import (
	"comment/internal/data/model"
)

type CommentDto struct {
	model.Comment
	User       *UserDto `gorm:"ForeignKey:UserUID;AssociationForeignKey:UID" json:"user"`
	CreateData string   `gorm:"-" json:"create_date"`
}
