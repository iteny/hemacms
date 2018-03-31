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
	ExcuteTime := c.TimeString(r, "ExcuteTime")
	defer Log().Debug().Str("[Method]", r.Method).Str("[Addr]", r.RequestURI).Str("[Ip]", ip).Str("[Status]", "200").Str("[ExcuteTime]", ExcuteTime).Msg("Get Template")
	t, err := template.ParseFiles(filenames...)
	c.Log().CheckErr("Template Error", err)
	err = t.Execute(w, data)
	c.Log().CheckErr("Template Error", err)
}
func (c *BaseCtrl) TimeString(r *http.Request, contextName string) string {
	contextValue := r.Context().Value("ExcuteTime")
	if contextValue != nil {
		timetype := time.Unix(r.Context().Value("ExcuteTime").(int64), 0)
		now := time.Now()
		duration := now.Sub(timetype)
		return duration.String()
	} else {
		return "0ms"
	}
}

//响应json
func (c *BaseCtrl) ResponseJson(status interface{}, info interface{}, w io.Writer, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	ExcuteTime := c.TimeString(r, "ExcuteTime")
	defer Log().Debug().Str("[Method]", r.Method).Str("[Addr]", r.RequestURI).Str("[Ip]", ip).Str("[Status]", "200").Str("[ExcuteTime]", ExcuteTime).Msg("Response Json")
	m := make(map[string]interface{})
	m["status"] = status
	m["info"] = info
	mData, err := json.Marshal(m)
	c.Log().CheckErr("Json Error", err)
	fmt.Fprint(w, string(mData))
	// return string(mData)
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
