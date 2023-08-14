package biz

import (
	feed "feed/internal/biz/feed"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(feed.NewFeedServiceImpl)
