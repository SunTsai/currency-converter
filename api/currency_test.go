package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvertCurrencyAPI(t *testing.T) {
	server := newTestServer(t)

	testCases := []struct {
		source        string
		target        string
		amount        int
		checkResponse func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			source: "TWD",
			target: "JPY",
			amount: 100,
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				resBody, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				var res currencyConversionResponse
				err = json.Unmarshal(resBody, &res)
				require.NoError(t, err)

				require.Equal(t, res.Msg, "success")
				require.Equal(t, res.Amount, "$366.90")
			},
		},
		{
			source: "USD",
			target: "JPY",
			amount: 1234567,
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				resBody, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				var res currencyConversionResponse
				err = json.Unmarshal(resBody, &res)
				require.NoError(t, err)

				require.Equal(t, res.Msg, "success")
				require.Equal(t, res.Amount, "$138,025,825.17")
			},
		},
		{
			source: "JPY",
			target: "JPY",
			amount: 1234567,
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				resBody, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				var res currencyConversionResponse
				err = json.Unmarshal(resBody, &res)
				require.NoError(t, err)

				require.Equal(t, res.Msg, "success")
				require.Equal(t, res.Amount, "$1,234,567.00")
			},
		},
		{
			source: "CAD",
			target: "JPY",
			amount: 100,
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			source: "TWD",
			target: "CAD",
			amount: 100,
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			source: "TWD",
			target: "USD",
			amount: -10,
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/currencyConversion?source=%s&target=%s&amount=%d",
			tc.source, tc.target, tc.amount)

		request, err := http.NewRequest(http.MethodGet, url, bytes.NewReader(nil))
		require.NoError(t, err)

		server.router.ServeHTTP(recorder, request)
		tc.checkResponse(t, recorder)
	}
}
