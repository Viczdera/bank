package api

//same package as crud code
//main_test used to set up connection with query object

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// testing
func TestMain(t *testing.M) {

	gin.SetMode(gin.TestMode)

	//run and report back to test runner via the o.exit command
	os.Exit(t.Run())

}
