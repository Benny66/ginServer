package api

import (
	"fmt"

	"github.com/Benny66/ginServer/utils/format"

	"github.com/gin-gonic/gin"
)

func TestPostHandler(context *gin.Context, req TestPostReqParams) {
	fmt.Println(req)
	format.NewResponseJson(context).Success(req)
}
