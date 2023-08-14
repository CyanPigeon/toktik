package service

import (
	fav "favorite/internal/service/favorite"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(fav.NewLikeListService, fav.NewLikeActionService)
