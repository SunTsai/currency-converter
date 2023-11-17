package api

import (
	"os"
	"suntsai/currency-converter/util"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func newTestServer(t *testing.T) *Server {
	exchangeRates, err := util.LoadExchangeRates("..")
	require.NoError(t, err)

	server, err := NewServer(exchangeRates)
	require.NoError(t, err)
	return server
}
