package main

import (
	"flag"
	"fmt"
	"ginServer/config"
	_ "ginServer/migrations"
	"ginServer/routers"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kardianos/service"

	_ "net/http/pprof"
)

var (
	help    bool
	version bool
	errChan = make(chan error)
	logger  service.Logger
)

func init() {
	flag.BoolVar(&help, "help", false, "help")
	flag.BoolVar(&version, "version", false, "version")
	flag.Usage = usage
	flag.Parse()
}

func main() {
	if help {
		flag.Usage()
		return
	}

	if version {
		fmt.Println(config.Config.GetAppVersion())
		return
	}

	svcConfig := &service.Config{
		Name:        "ginServer",
		DisplayName: "ginServer",
		Description: "基于go-gin的web服务框架",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			s.Install()
			logger.Info("服务安装成功")
			return
		}

		if os.Args[1] == "remove" {
			s.Uninstall()
			logger.Info("服务卸载成功")
			return
		}
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
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

func (p *program) Stop(s service.Service) error {
	return nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of Server:\n")
	flag.PrintDefaults()
}
