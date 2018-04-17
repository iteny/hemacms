package common

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

//LogCtrl is local log object.
type LogCtrl struct {
}

//Log return a LogCtrl对象
func Log() *LogCtrl {
	return &LogCtrl{}
}

//Log return a LogCtrl对象
func (c *BaseCtrl) Log() *LogCtrl {
	return &LogCtrl{}
}
func init() {
	zerolog.TimeFieldFormat = "[2006-01-02 15:04:05]"
	logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	// logger.Debug().Msg("filtered out message")
	// logger.Info().Msg("routed message")
	// logger.Info().Str("foo", "bar").Str("ss", "111").Msg("Hello world")
}

//Debug info
func (c *LogCtrl) Debug() *zerolog.Event {
	// _, file, line, ok := runtime.Caller(1)
	return logger.Debug()
}

//CheckErr check error info
func (c *LogCtrl) CheckErr(msg string, err error) {
	var skip int
	if msg == "Template Error" || msg == "Response Json" {
		skip = 2
	} else {
		skip = 1
	}
	if err != nil {
		pc, file, line, ok := runtime.Caller(skip)
		l := strconv.Itoa(line)
		if ok {
			f := runtime.FuncForPC(pc)
			detail := "<span style='color:#3498db'>Package&nbsp;:&nbsp;</span>" + f.Name() + "<br/>" + "<span style='color:#e67e22'>File&nbsp;:&nbsp;</span>" + file + "<br/>" + "<span style='color:#9b59b6'>Line&nbsp;:&nbsp;</span>" + strconv.Itoa(line)
			s := make([]string, 2)
			s[0] = file
			s[1] = l
			terminalLog := Base().Config().Value("common", "terminalLog")
			if terminalLog == "on" {
				logger.Error().Strs("[PATH]", s).Err(err).Msg(msg)
			}
			sqlLog := Base().Config().Value("common", "sqlLog")
			if sqlLog == "on" {
				username := "system"
				logSql := "INSERT INTO hm_oprate_log(username,oprate_time,oprate_ip,useragent,detail,info,url,method,excute_time,status) VALUES(?,?,?,?,?,?,?,?,?,?)"
				tx := Base().Sql().MustBegin()
				tx.MustExec(logSql, username, time.Now().Unix(), "127.0.0.1", "System error", detail, err.Error(), "System error", "System error", "System error", 4)
				err = tx.Commit()
				if err != nil {
					fmt.Println(err.Error())
				} else {
					clearOprateLog := Base().Cache().Items()
					for k, _ := range clearOprateLog {
						if strings.Contains(k, "oprateLog") {
							Base().Cache().Del(k)
						}
					}
				}
			}
		}
	}
}

//CheckErrSkip check error info,Skip before to function
func (c *LogCtrl) CheckErrSkip(msg string, skip int, err error) {
	if err != nil {
		_, file, line, ok := runtime.Caller(skip)
		l := strconv.Itoa(line)
		if ok {
			s := make([]string, 2)
			s[0] = file
			s[1] = l
			logger.Error().Strs("[PATH]", s).Err(err).Msg(msg)
		}

	}
}

//Info echo a level-1 information
// func (c *LogCtrl) Info(s interface{}) {
// 	if i, e := Cfg.Int("Debug", "model"); e != nil {
// 		fmt.Print(e)
// 	} else {
// 		if i > 0 {
// 			_, file, line, ok := runtime.Caller(1)
// 			// defer glog.Flush()
// 			if !ok {
// 				fmt.Print(FmtColor.Green("\n[PATH:] "), "not found file path!", "\n[LINE:] ", "not found file path!")
// 			} else {
// 				fmt.Print(FmtColor.Green("\n[PATH:] "), file, FmtColor.Green("\n[LINE:] "), line, FmtColor.Green("\n[INFO:] "), s)
// 			}
// 		}
// 	}
// }
//
// //Warning echo a level-2 information
// func (c *LogCtrl) Warning(s interface{}) {
// 	if i, e := Cfg.Int("Debug", "model"); e != nil {
// 		fmt.Print(e)
// 	} else {
// 		if i > 1 {
// 			_, file, line, ok := runtime.Caller(1)
// 			// defer glog.Flush()
// 			if !ok {
// 				fmt.Print(FmtColor.Yellow("\n[PATH:] "), "not found file path!", "\n[LINE:] ", "not found file path!")
// 			} else {
// 				fmt.Print(FmtColor.Yellow("\n[PATH:] "), file, FmtColor.Yellow("\n[LINE:] "), line, FmtColor.Yellow("\n[INFO:] "), s)
// 			}
// 		}
// 	}
// }
//
// //Error echo a level-3 information and save log file
// func (c *LogCtrl) Error(s interface{}) {
// 	if i, e := Cfg.Int("Debug", "model"); e != nil {
// 		fmt.Print(e)
// 	} else {
// 		if i > 2 {
// 			_, file, line, ok := runtime.Caller(1)
// 			defer glog.Flush()
// 			if !ok {
// 				fmt.Print("\n")
// 				glog.Info("\n[PATH:] ", "not found file path!", "\n[LINE:] ", "not found file path!")
// 			} else {
// 				fmt.Print("\n")
// 				glog.Info("\n[PATH:] ", file, "\n[LINE:] ", line, "\n[INFO:] ", s, "\n")
// 			}
// 		}
// 	}
// }
//
// //ErrorCheck echo a level-3 information and save log file.
// //The incoming parameter is error type.
// func (c *LogCtrl) ErrorCheck(err error) {
// 	if i, e := Cfg.Int("Debug", "model"); e != nil {
// 		fmt.Print(e)
// 	} else {
// 		if i > 2 {
// 			_, file, line, ok := runtime.Caller(1)
// 			defer glog.Flush()
// 			if !ok {
// 				fmt.Print("\n")
// 				glog.Info("\n[PATH:] ", "not found file path!", "\n[LINE:] ", "not found file path!")
// 			} else {
// 				if err != nil {
// 					fmt.Print("\n")
// 					glog.Info("\n[PATH:] ", file, "\n[LINE:] ", line, "\n[INFO:] ", err, "\n")
// 				}
// 			}
// 		}
// 	}
// }
