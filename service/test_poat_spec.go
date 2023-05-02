package api

type TestPostReqParams struct {
	Name string `form:"name" json:"name"`
	Age  int    `form:"age" json:"age"`
}
