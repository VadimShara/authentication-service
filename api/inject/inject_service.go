package inject

import (
	"vk-test-task/internal/service/auth"

	"github.com/google/wire"
)

var serviceSet = wire.NewSet( // nolint
	provideAuthService,
)

func provideAuthService(s stores) (auth.Service, error) {
	return auth.New(s.users)
}