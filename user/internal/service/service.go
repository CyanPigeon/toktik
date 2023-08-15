package service

import (
	"github.com/google/wire"
	u "user/internal/service/user"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(u.NewUserLoginService, u.NewUserRegisterService, u.NewUserInfoService)
