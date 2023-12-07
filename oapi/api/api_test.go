package api_test

import (
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/maragudk/oapi-codegen-bug/oapi/api"
)

func TestRegisterHandlers(t *testing.T) {
	g := gin.New()
	api.RegisterHandlers(g, &server{})
}

type server struct{}

func (s *server) GetPing(c *gin.Context) {}
