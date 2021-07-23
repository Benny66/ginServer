package language

import (
	"fmt"
	"ginServer/config"
	"log"
)

const (
	LANG_TYPE_CHINESE = "zh-cn"
	LANG_TYPE_EN      = "en"
)

var Lang *lang

func init() {
	Lang = NewLang()
}

func NewLang() *lang {
	l := &lang{
		lang: LANG_TYPE_CHINESE,
	}
	l.setLanguage(config.Config.GetLanguage())
	return l
}

type lang struct {
	lang string
}

func (l *lang) setLanguage(lang string) {
	if lang != LANG_TYPE_CHINESE && lang != LANG_TYPE_EN {
		log.Fatal("language setting failed!")
	}
	l.lang = lang
}

func (l *lang) Msg(errorCode int, params ...interface{}) string {
	switch l.lang {
	case LANG_TYPE_CHINESE:
		if msg, ok := chineseMsgMap[errorCode]; !ok {
			return "invalid error code"
		} else {
			return fmt.Sprintf(msg, params...)
		}
	case LANG_TYPE_EN:
		if msg, ok := englishMsgMap[errorCode]; !ok {
			return "invalid error code"
		} else {
			return fmt.Sprintf(msg, params...)
		}
	}
	return ""
}
