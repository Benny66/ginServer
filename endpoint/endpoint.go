package endpoint

import "github.com/Benny66/ginServer/routers"

type Endpoint interface {
	Router() *routers.Endpoint
}
