package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Viczdera/bank/token"
	"github.com/gin-gonic/gin"
)

var (
	AUTH_HEADER_KEY  = "authorization"
	AUTH_TYPE        = "bearer"
	AUTH_PAYLOAD_KEY = "authorization_payload"
)

func abort(ctx *gin.Context, errMessage string) {
	err := errors.New(errMessage)
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
	return
}

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {

	//to return a gin context handler function
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AUTH_HEADER_KEY)
		if len(authorizationHeader) == 0 {
			abort(ctx, "authorization not found")
		}

		//split header
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			abort(ctx, "invalid authorization ")
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != AUTH_TYPE {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			abort(ctx, err.Error())
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			abort(ctx, err.Error())
		}

		//token valid.
		//next, store payload in context, then forward request to next handler
		ctx.Set(AUTH_PAYLOAD_KEY, payload)
		ctx.Next()
	}
}
