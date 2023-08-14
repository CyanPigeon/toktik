package biz

import (
	fav "favorite/internal/biz/favorite"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(fav.NewFavoriteServiceBiz)
