package infra

import (
	"github.com/google/wire"
)

// ProviderSet is icon providers.
var ProviderSet = wire.NewSet(InitDB, InitRDB)
