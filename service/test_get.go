package api

import (
	"fmt"

	"github.com/Benny66/ginServer/utils/format"
	"github.com/gin-gonic/gin"
)

func TestGetHandler(context *gin.Context, req TestGetReqParams) {
	fmt.Println(req)
	format.NewResponseJson(context).Success(req)
}
