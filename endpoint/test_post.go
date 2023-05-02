package endpoint

import (
	"github.com/Benny66/ginServer/routers"
	api "github.com/Benny66/ginServer/service"
	"github.com/Benny66/ginServer/utils/format"
	"github.com/gin-gonic/gin"
)

func init() {
	routers.R.AddEndpointSchema(&testPostApiImpl{})
}

type testPostApiImpl struct{}

func (r *testPostApiImpl) Group() string {
	return "api"
}
func (r *testPostApiImpl) Auth() string {
	return ""
}

func (r *testPostApiImpl) Method() string {
	return "POST"
}
func (r *testPostApiImpl) URL() string {
	return "/test/post"
}

func (r *testPostApiImpl) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params api.TestPostReqParams
		//todo: 这里只提供简单的参数处理映射到结构体中
		if err := ctx.BindJSON(&params); err != nil {
			// 处理错误
			format.NewResponseJson(ctx).Error(1, err.Error())
			return
		}
		api.TestPostHandler(ctx, params)
	}
}
