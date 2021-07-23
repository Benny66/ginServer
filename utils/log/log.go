package log

import (
	"encoding/json"
	"fmt"
	"ginServer/config"
	"ginServer/utils/function"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const LOG_LEVEL_DEBUG = "debug"

var (
	SystemLogger *fileLogger
	SQLLogger    *fileLogger
	APILogger    *fileLogger
	WSLogger     *fileLogger
	UDPLogger    *fileLogger
)

func init() {
	SystemLogger = NewFileLogger(function.GetAbsPath(config.Config.GetSystemLogPath()), "system", config.Config.Mode, config.Config.LogExpire)
	SQLLogger = NewFileLogger(function.GetAbsPath(config.Config.GetSqlLogPath()), "sql", config.Config.Mode, config.Config.LogExpire)
	APILogger = NewFileLogger(function.GetAbsPath(config.Config.GetApiLogPath()), "api", config.Config.Mode, config.Config.LogExpire)
	WSLogger = NewFileLogger(function.GetAbsPath(config.Config.GetWsLogPath()), "ws", config.Config.Mode, config.Config.LogExpire)
	UDPLogger = NewFileLogger(function.GetAbsPath(config.Config.GetUDPLogPath()), "udp", config.Config.Mode, config.Config.LogExpire)
}

/*
* description: 系统日志记录器
* author: shahao
* created on: 19-11-20 上午11:40
* param param_1:
* param param_2:
* return return_1:
 */
func SystemLog(msg interface{}) {
	line, functionName := 0, "???"
	pc, _, line, ok := runtime.Caller(1)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
	}
	SystemLogger.writeLog(fmt.Sprintf("[System] %v | %s:%d | %s\n", time.Now().Format("2006/01/02 - 15:04:05"), functionName, line, fmt.Sprintf("%s", msg)))
}

/*
* description: API接口日志记录器
* author: shahao
* created on: 19-11-20 上午11:40
* param param_1:
* param param_2:
* return return_1:
 */
func APILog(msg string, params ...interface{}) {
	line, functionName := 0, "???"
	pc, _, line, ok := runtime.Caller(1)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
	}
	APILogger.writeLog(fmt.Sprintf("[API] %v | %s:%d | %s\n", time.Now().Format("2006/01/02 - 15:04:05"), functionName, line, fmt.Sprintf(msg, params...)))
}

/*
* description: ws接口日志记录器
* author: jiangjm
* created on: 2021-04-08 上午11:20
* param param_1:
* param param_2:
* return return_1:
 */
func WSLog(msg string, params ...interface{}) {
	line, functionName := 0, "???"
	pc, _, line, ok := runtime.Caller(1)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
	}
	WSLogger.writeLog(fmt.Sprintf("[WS] %v | %s:%d | %s\n", time.Now().Format("2006/01/02 - 15:04:05"), functionName, line, fmt.Sprintf(msg, params...)))
}

/*
* description: 数据库日志记录器
* author: shahao
* created on: 19-11-20 上午11:40
* param param_1:
* param param_2:
* return return_1:
 */
func SQLLog(msg interface{}) {
	SQLLogger.writeLog(fmt.Sprintf("[SQL] %v | %s\n", time.Now().Format("2006/01/02 - 15:04:05"), fmt.Sprintf("%s", msg)))
}

func UdpLog(code, msg string, content []byte) {
	line, functionName := 0, "???"
	pc, _, line, ok := runtime.Caller(1)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
	}

	UDPLogger.writeLog(fmt.Sprintf("[UDP] %v | %s:%d | %s\n", time.Now().Format("2006/01/02 - 15:04:05"), functionName, line, fmt.Sprintf("%s:%s---%s", code, msg, function.BytetoHexSpace(content))))
}

type fileLogger struct {
	FileDir   string
	Prefix    string
	WriteChan chan string
	Mode      string
	LogExpire int
}

func NewFileLogger(dir, prefix, mode string, expire int) *fileLogger {
	function.FileNotExistsAndCreate(dir)
	l := &fileLogger{
		FileDir:   dir,
		Prefix:    prefix,
		WriteChan: make(chan string, 20),
		Mode:      mode,
		LogExpire: expire,
	}
	go l.work()
	go l.timeClean()
	return l
}

func (this *fileLogger) work() {
	for content := range this.WriteChan {
		savePath := filepath.Join(this.FileDir, this.Prefix+"-"+time.Now().Format("20060102")+".log")
		f, err := os.OpenFile(savePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		_, err = f.Write([]byte(content))
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if err := f.Close(); err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
}

func (this *fileLogger) timeClean() {
	now := time.Now()
	next := now.Add(24 * time.Hour)
	next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
	t := time.NewTimer(next.Sub(now))
	defer t.Stop()

	for {
		SystemLog("启动日志文件定期清除监测:" + time.Now().Format("2006/01/02 - 15:04:05"))
		select {
		case <-t.C:
			files, err := ioutil.ReadDir(this.FileDir)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			for _, f := range files {
				if f.ModTime().Before(time.Now().Add(-1 * time.Duration(this.LogExpire) * 24 * time.Hour)) {
					err := os.Remove(filepath.Join(this.FileDir, f.Name()))
					if err != nil {
						fmt.Println(err.Error())
						continue
					}
				}
			}

			time.Sleep(1 * time.Second)

			now := time.Now()
			next := now.Add(24 * time.Hour)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t.Reset(next.Sub(now))
		}
	}
}

func (this *fileLogger) writeLog(msg string) {
	this.WriteChan <- msg
	if this.IsDebugging() {
		os.Stdout.WriteString(msg)
	}
}

func (this *fileLogger) Print(v ...interface{}) {
	data, _ := json.Marshal(v)
	this.writeLog(string(data))
}

func (this *fileLogger) IsDebugging() bool {
	return this.Mode == LOG_LEVEL_DEBUG
}

func (this *fileLogger) Write(p []byte) (n int, err error) {
	this.writeLog(string(p))
	return 0, nil
}
