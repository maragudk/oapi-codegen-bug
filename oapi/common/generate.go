package common

import (
	_ "github.com/deepmap/oapi-codegen/v2/pkg/util"
)

//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml ../common.yaml
