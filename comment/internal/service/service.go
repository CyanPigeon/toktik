package service

import (
	comment "comment/internal/service/comment"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(comment.NewCommentActionService, comment.NewCommentListService)
