package biz

import (
	diz "comment/internal/biz/comment"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(diz.NewDizCommentServiceImpl)
