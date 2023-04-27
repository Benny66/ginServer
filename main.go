package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Benny66/ginServer/config"
	"github.com/Benny66/ginServer/routers"

	_ "net/http/pprof"
)

var (
	errChan = make(chan error)
)

func main() {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	//启动web服务
	go func() {
		if err := routers.Router.Init().Run(config.Config.IPAddress + ":" + config.Config.Port); err != nil {
			fmt.Println(err)
			errChan <- err
		}
	}()

	if err := <-errChan; err != nil {
		log.Fatal(err)
	}
}
