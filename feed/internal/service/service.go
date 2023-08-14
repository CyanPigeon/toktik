package service

import (
	feed "feed/internal/service/feed"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(feed.NewFeedService)
