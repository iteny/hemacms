package admin

import (
	"bytes"
	"fmt"
	"hemacms/common"
	"hemacms/common/sql"
	"net"
	"net/http"
	"strconv"
	"time"
)

type LoginCtrl struct {
	common.BaseCtrl
}

func LoginCtrlObject() *LoginCtrl {
	return &LoginCtrl{}
}
func (c *LoginCtrl) Login(w http.ResponseWriter, r *http.Request) {
	c.Template(w, r, nil, "./view/admin/login/index.html")
}
func (c *LoginCtrl) LoginSubmit(w http.ResponseWriter, r *http.Request) {
	username, password := r.PostFormValue("username"), r.PostFormValue("password")
	b := bytes.Buffer{}
	b.WriteString("errored")
	b.WriteString(username)
	s := b.String()
	fod, foundd := c.Cache().CacheGet(s)
	if foundd {
		if fod.(int) > 2 {
			// fmt.Fprint(w, c.ResponseJson(4, "密码错误3次，需要等待1分钟后再登录，谢谢！"))
			c.ResponseJson(4, "密码错误3次，需要等待1分钟后再登录，谢谢！", w, r)
			return
		}
	}
	isUser := c.Regexp().Username(username)
	isPwd := c.Regexp().Password(password)
	switch false {
	case isUser:
		// fmt.Fprint(w, c.ResponseJson(4, "用户名的长度为5位到15位！"))
		c.ResponseJson(11, "", w, r)
		return
	case isPwd:
		// fmt.Fprint(w, c.ResponseJson(4, "密码的长度为5位到15位！"))
		c.ResponseJson(12, "", w, r)
		return
	default:
		// common.Log.Critical("1111")
		userone := sql.User{}
		sqls := "SELECT id,username,status FROM hm_user WHERE username = ? AND password = ?"
		err := c.Sql().Get(&userone, sqls, username, c.Sha1PlusMd5(password))
		if err != nil {
			// common.Log.Error(err)
			c.Log().CheckErr("Password Error", err)
			errored := 1
			b := bytes.Buffer{}
			b.WriteString("errored")
			b.WriteString(username)
			s := b.String()
			fo, found := c.Cache().CacheGet(s)
			if found {
				errored = fo.(int) + errored
			}
			c.Cache().CacheSetConfineTime(s, errored)
			// fmt.Fprint(w, c.ResponseJson(4, "用户名或密码错误！"))
			c.ResponseJson(4, "", w, r)
		} else {
			if userone.Id != 0 {
				if userone.Status == 1 {
					ip, _, _ := net.SplitHostPort(r.RemoteAddr)
					if ip == "::1" {
						ip = "127.0.0.1"
					}
					ipinfo := c.TaobaoIP(ip)
					sqls := "INSERT INTO hm_login_log(username,login_time,login_ip,status,area,country,useragent,uid) VALUES(?,?,?,?,?,?,?,?)"
					tx := c.Sql().MustBegin()
					tx.MustExec(sqls, userone.Username, time.Now().Unix(), ip, 1, ipinfo.Data.Area, ipinfo.Data.Country, r.UserAgent(), userone.Id)
					tx.Commit()
					session := c.Sess().Load(r)
					// session, _ := common.Sess.Get(r, "hm-back-stage")
					// session.Values["uid"] = userone.Id
					// session.Values["username"] = userone.Username
					// session.Values["status"] = userone.Status
					// session.Save(r, w)
					access := sql.AuthRoleAccess{}
					fmt.Println("用户id=", userone.Id)
					accessSql := "SELECT * FROM hm_auth_role_access WHERE uid = ?"
					err := c.Sql().Get(&access, accessSql, userone.Id)
					c.Log().CheckErr("Sql Error", err)
					fmt.Println("access=", access.RoleId)
					role := sql.AuthRole{}
					roleSql := "SELECT * FROM hm_auth_role WHERE id = ?"
					err = c.Sql().Get(&role, roleSql, access.RoleId)
					c.Log().CheckErr("Sql Error", err)

					// err = session.PutString(w, "roleId", string(role.Id))
					// c.Log().CheckErr("Session Error", err)
					// err = session.PutString(w, "uidint", strconv.Itoa(userone.Id))
					// c.Log().CheckErr("Session Error", err)
					err = session.PutString(w, "uid", strconv.Itoa(userone.Id))
					c.Log().CheckErr("Session Error", err)
					err = session.PutString(w, "username", string(userone.Username))
					c.Log().CheckErr("Session Error", err)
					err = session.PutString(w, "status", string(userone.Status))
					c.Log().CheckErr("Session Error", err)
					c.Cache().CacheSetAlwaysTime("roleId"+string(userone.Id), strconv.Itoa(role.Id))
					c.Cache().CacheSetAlwaysTime("auth"+string(role.Id), role.Rules)
					// fmt.Fprint(w, c.ResponseJson(1, "登录成功，3秒后为你跳转！"))
					c.ResponseJson(1, "", w, r)
					return
				} else {
					// fmt.Fprint(w, c.ResponseJson(4, "该账号已被封停！"))
					c.ResponseJson(7, "", w, r)
				}
			} else {
				// fmt.Fprint(w, c.ResponseJson(4, "用户名或密码错误！"))
				c.ResponseJson(5, "", w, r)
			}
		}
		// fmt.Printf("%v", userone.Id)

		// row, err := sql.SelectOne(sqls, username, common.Sha1PlusMd5(password))
		// row := *rows
		// fmt.Println(row)
		// if err != nil {
		// 	common.LogerInsertText("./controllers/intendant/login.go:43line", err.Error())
		// 	return
		// } else {
		// 	if row["id"] != "" {
		// 		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		// 		if ip == "::1" {
		// 			ip = "127.0.0.1"
		// 		}
		// 		ipinfo := common.TaobaoIP(ip)
		// 		sqls := "INSERT INTO hm_login_log(username,login_time,login_ip,status,info,area,country,useragent,uid) VALUES(?,?,?,?,?,?,?,?,?)"
		// 		sql.Insert(sqls, row["username"], time.Now().Unix(), ip, 1, "登录成功", ipinfo.Data.Area, ipinfo.Data.Country, r.UserAgent(), row["id"])
		// 		session, _ := common.Sess.Get(r, "hmcms")
		// 		session.Values["uid"] = row["id"]
		// 		session.Values["username"] = row["username"]
		// 		session.Values["status"] = row["status"]
		// 		session.Save(r, w)
		//
		// 		fmt.Fprint(w, common.ResponseJson(1, "登录成功，3秒后为你跳转！"))
		// 		return
		// 	} else {
		// 		errored := 1
		// 		b := bytes.Buffer{}
		// 		b.WriteString("errored")
		// 		b.WriteString(username)
		// 		s := b.String()
		// 		fo, found := common.CacheGet(s)
		// 		if found {
		// 			errored = fo.(int) + errored
		// 		}
		// 		common.CacheSetConfineTime(s, errored)
		// 		fmt.Fprint(w, common.ResponseJson(4, "用户名或密码错误！"))
		// 	}
		// }

	}
}

//退出
func (c *LoginCtrl) LoginOut(w http.ResponseWriter, r *http.Request) {
	session := c.Sess().Load(r)
	err := session.PutString(w, "uid", "")
	if err != nil {
		c.Log().Debug().Err(err).Msg("错误")
		c.ResponseJson(4, err.Error(), w, r)
	} else {
		c.ResponseJson(1, "", w, r)
	}
}
