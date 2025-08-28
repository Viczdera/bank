package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Viczdera/bank/token"
	"github.com/Viczdera/bank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestAuthMiddleware(t *testing.T) {

	testCases := []struct {
		name           string
		setupAuth      func(*testing.T, *http.Request, token.Maker)
		requiredChecks func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				accessToken, err := tokenMaker.CreateToken(util.RandomOwner(), time.Minute)
				require.NoError(t, err)

				//setauth header
				authorizationHeader := fmt.Sprintf("%s %s", AUTH_TYPE, accessToken)
				request.Header.Set(AUTH_HEADER_KEY, authorizationHeader)
			},
			requiredChecks: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for i := range testCases {
		testCase := testCases[i]

		t.Run(testCase.name, func(t *testing.T) {
			//create dummy server
			server := newTestServer(t, nil)

			//dummy auth path
			authPath := "/auth"
			server.router.GET(authPath, func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{})
			})

			//recorder to record the request
			recoder := httptest.NewRecorder()

			//make requests
			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			testCase.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recoder, request)

			//verify reesults
			testCase.requiredChecks(t, recoder)
		})
	}

}
