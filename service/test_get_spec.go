package api

type TestGetReqParams struct {
	Name string `form:"name" json:"name"`
	Age  int    `form:"age" json:"age"`
}
