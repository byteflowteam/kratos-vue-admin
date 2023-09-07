package service

import (
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewSysuserService,
	NewMenusService,
	NewRolesService,
	NewApiService,
	NewDeptService,
	NewPostService,
	NewDictDataService,
	NewDictTypeService,
)
