package common

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs"
)

var sessionManage = scs.NewCookieManager("asldfjsladasdfasd")

type BaseCtrl struct {
	// LogCtrl
}

func Sess() *scs.Manager {
	return sessionManage
}
func (c *BaseCtrl) Sess() *scs.Manager {
	return sessionManage
}

//regexp return a regexpCtrl对象
func Base() *BaseCtrl {
	return &BaseCtrl{}
}
func init() {
	sessionManage.Lifetime(10 * time.Hour) // Set the maximum session lifetime to 1 hour.
	sessionManage.Persist(false)           // Persist the session after a user has closed their browser.
	// sessionManage.Secure(true)            // Set the Secure flag on the session cookie.
}
func (c *BaseCtrl) Template(w io.Writer, r *http.Request, data interface{}, filenames ...string) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	excuteTime := c.TimeString(r, "excuteTime")
	terminalLog := c.Config().Value("common", "terminalLog")
	if terminalLog == "on" {
		defer Log().Debug().Str("[Method]", r.Method).Str("[Addr]", r.RequestURI).Str("[Ip]", ip).Str("[Status]", "200").Str("[ExcuteTime]", excuteTime).Msg("Get Template")
	}
	t, err := template.ParseFiles(filenames...)
	c.Log().CheckErr("Template Error", err)
	err = t.Execute(w, data)
	c.Log().CheckErr("Template Error", err)
}
func (c *BaseCtrl) TimeString(r *http.Request, contextName string) string {
	contextValue := r.Context().Value("excuteTime")
	if contextValue != nil {
		timetype := time.Unix(r.Context().Value("excuteTime").(int64), 0)
		now := time.Now()
		duration := now.Sub(timetype)
		excuteTime := strconv.Itoa(int(duration.Nanoseconds() / 1000 / 1000))
		return excuteTime
	} else {
		return "0"
	}
}

//响应json
func (c *BaseCtrl) ResponseJson(status interface{}, info interface{}, w io.Writer, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	var infos string = info.(string)
	if infos == "" {
		infos = "no problem"
	}
	excuteTime := c.TimeString(r, "excuteTime")
	terminalLog := c.Config().Value("common", "terminalLog")
	if terminalLog == "on" {
		path := ""
		_, file, line, ok := runtime.Caller(1)
		if ok {
			path = file + " Line:" + strconv.Itoa(line)
		}
		defer Log().Debug().Str("[Method]", r.Method).Str("[Addr]", r.RequestURI).Str("[Ip]", ip).Str("[Status]", "200").Str("[ExcuteTime]", excuteTime).Str("[Path]", path).Str("[Info]", info.(string)).Msg("Response Json")
	}
	sqlLog := c.Config().Value("common", "sqlLog")
	if sqlLog == "on" {
		if r.RequestURI == "/intendant/login" {
		} else {
			detail := ""
			pc, file, line, ok := runtime.Caller(1)
			if ok {
				f := runtime.FuncForPC(pc)
				detail = "<span style='color:#3498db'>Package&nbsp;:&nbsp;</span>" + f.Name() + "<br/>" + "<span style='color:#e67e22'>File&nbsp;:&nbsp;</span>" + file + "<br/>" + "<span style='color:#9b59b6'>Line&nbsp;:&nbsp;</span>" + strconv.Itoa(line)
			}
			session := c.Sess().Load(r)
			username, err := session.GetString("username")
			c.Log().CheckErr("Session Get Error", err)
			logSql := "INSERT INTO hm_oprate_log(username,oprate_time,oprate_ip,useragent,detail,info,url,method,excute_time,status) VALUES(?,?,?,?,?,?,?,?,?,?)"
			tx := c.Sql().MustBegin()
			tx.MustExec(logSql, username, time.Now().Unix(), ip, r.UserAgent(), detail, infos, r.RequestURI, r.Method, excuteTime, status)
			err = tx.Commit()
			if err != nil {
				c.Log().CheckErr("Sql Error", err)
			} else {
				clearOprateLog := c.Cache().Items()
				for k, _ := range clearOprateLog {
					if strings.Contains(k, "oprateLog") {
						c.Cache().Del(k)
					}
				}
			}
		}
	}
	m := make(map[string]interface{})
	m["status"] = status
	m["info"] = info
	mData, err := json.Marshal(m)
	c.Log().CheckErr("Json Error", err)
	fmt.Fprint(w, string(mData))
}

//响应数据
func (c *BaseCtrl) ResponseData(data interface{}, w io.Writer, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	dt := data.(map[string]interface{})
	var infos string
	if _, ok := dt["info"]; ok {
		infos = dt["info"].(string)
		if infos == "" {
			infos = "no problem"
		}
	} else {
		infos = "no problem"
	}
	var status interface{}
	if _, ok := dt["status"]; ok {
		status = dt["status"]
	} else {
		status = "1"
	}
	excuteTime := c.TimeString(r, "excuteTime")
	terminalLog := c.Config().Value("common", "terminalLog")
	if terminalLog == "on" {
		path := ""
		_, file, line, ok := runtime.Caller(1)
		if ok {
			path = file + " Line:" + strconv.Itoa(line)
		}
		defer Log().Debug().Str("[Method]", r.Method).Str("[Addr]", r.RequestURI).Str("[Ip]", ip).Str("[Status]", "200").Str("[ExcuteTime]", excuteTime).Str("[Path]", path).Str("[Info]", infos).Msg("Response Data")
	}
	sqlLog := c.Config().Value("common", "sqlLog")
	if sqlLog == "on" {
		if r.RequestURI == "/intendant/site/getOprateLog" {
		} else if r.RequestURI == "/intendant/ajaxPolling" {
		} else {
			detail := ""
			pc, file, line, ok := runtime.Caller(1)
			if ok {
				f := runtime.FuncForPC(pc)
				detail = "<span style='color:#3498db'>Package&nbsp;:&nbsp;</span>" + f.Name() + "<br/>" + "<span style='color:#e67e22'>File&nbsp;:&nbsp;</span>" + file + "<br/>" + "<span style='color:#9b59b6'>Line&nbsp;:&nbsp;</span>" + strconv.Itoa(line)
			}
			session := c.Sess().Load(r)
			username, err := session.GetString("username")
			c.Log().CheckErr("Session Get Error", err)
			logSql := "INSERT INTO hm_oprate_log(username,oprate_time,oprate_ip,useragent,detail,info,url,method,excute_time,status) VALUES(?,?,?,?,?,?,?,?,?,?)"
			tx := c.Sql().MustBegin()
			tx.MustExec(logSql, username, time.Now().Unix(), ip, r.UserAgent(), detail, infos, r.RequestURI, r.Method, excuteTime, status)
			err = tx.Commit()
			if err != nil {
				c.Log().CheckErr("Sql Error", err)
			} else {
				clearOprateLog := c.Cache().Items()
				for k, _ := range clearOprateLog {
					if strings.Contains(k, "oprateLog") {
						c.Cache().Del(k)
					}
				}
			}
		}
	}
	// m := make(map[string]interface{})
	// m["status"] = status
	// m["info"] = info
	mData, err := json.Marshal(dt)
	c.Log().CheckErr("Json Data Error", err)
	fmt.Fprint(w, string(mData))
}

//return rows json
func (c *BaseCtrl) RowsJson(rows interface{}) string {

	mData, err := json.Marshal(rows)
	c.Log().CheckErr("Json Error", err)
	return string(mData)
}
func (c *BaseCtrl) Md5(s string) string {
	hash := md5.New()
	buf := []byte(s)
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
func (c *BaseCtrl) Sha1(s string) string {
	hash := sha1.New()
	buf := []byte(s)
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//Sha1 Plus Md5
func (c *BaseCtrl) Sha1PlusMd5(s string) string {
	return c.Sha1(c.Md5(s))
}

//格式化路由
func (c *BaseCtrl) FormatUrl(s string) string {
	uri := s
	straddress := ""
	for _, v := range uri { // i 是字符的字节位置，v 是字符的拷贝
		ss := fmt.Sprintf("%c", v) // 输出单个字符
		yz := Base().Regexp().Id(ss)
		if yz == false {
			straddress = straddress + ss
		}
	}
	result := strings.TrimRight(straddress, "/")
	return result
}

//分页
func (c *BaseCtrl) Pagination(page, row string) int {
	ipage, _ := strconv.Atoi(page)
	irow, _ := strconv.Atoi(row)
	first := irow * (ipage - 1)
	return first
}

//date转化时间戳
func (c *BaseCtrl) DateToTimestamp(date, format string) int64 {
	loc, _ := time.LoadLocation("Local") //重要：获取时���
	from, _ := time.ParseInLocation(format, date, loc)
	return from.Unix()
}
