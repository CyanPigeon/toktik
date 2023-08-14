package biz

import (
	"github.com/google/wire"
	biz "message/internal/biz/message"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(biz.NewBizMessageService)
