package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Benny66/ginServer/config"
	"github.com/Benny66/ginServer/endpoint"
	"github.com/Benny66/ginServer/middleware"
	"github.com/Benny66/ginServer/routers"

	_ "net/http/pprof"
)

var (
	errChan = make(chan error)
)
var p = &app{}

type app struct {
	endpoint.Endpoint
	middleware.Middleware
}

func main() {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	//启动web服务
	go func() {
		r := p.Router()
		r.Register()
		if err := r.Engine().Run(config.Config.IPAddress + ":" + config.Config.Port); err != nil {
			fmt.Println(err)
			errChan <- err
		}
	}()

	if err := <-errChan; err != nil {
		log.Fatal(err)
	}
}

// func (p *app) Name() string {
// 	return "ginServer"
// }

func (p *app) Router() routers.RouterRegister {
	return *routers.R
}
