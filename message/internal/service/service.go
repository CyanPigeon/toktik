package service

import (
	"github.com/google/wire"
	service "message/internal/service/message"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(service.NewMessageHistoryService, service.NewMessageActionService)
