package app

import (
	"github.com/google/wire"
	"github.com/isuquo/fortune3-pickaxe/pkg/goals"
)

var WireModule = wire.NewSet(
	goals.WireModule,
)
