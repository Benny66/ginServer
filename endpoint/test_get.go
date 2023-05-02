package endpoint

import (
	"github.com/Benny66/ginServer/routers"
	api "github.com/Benny66/ginServer/service"
	"github.com/Benny66/ginServer/utils/format"
	"github.com/gin-gonic/gin"
)

func init() {
	routers.R.AddEndpointSchema(&testGetApiImpl{})
}

type testGetApiImpl struct{}

func (r *testGetApiImpl) Group() string {
	return "api"
}
func (r *testGetApiImpl) Auth() string {
	return ""
}

func (r *testGetApiImpl) Method() string {
	return "GET"
}
func (r *testGetApiImpl) URL() string {
	return "/test/get"
}

func (r *testGetApiImpl) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params api.TestGetReqParams
		//todo: 这里只提供简单的参数处理映射到结构体中,可根据需要处理
		if err := ctx.Bind(&params); err != nil {
			// 处理错误
			format.NewResponseJson(ctx).Error(1, err.Error())
			return
		}
		api.TestGetHandler(ctx, params)
	}
}
