package infra

import (
	"github.com/google/wire"
)

// ProviderSet is infra providers.
var ProviderSet = wire.NewSet(InitDB, InitRDB)
