package middle

import (
	"context"
	"fmt"
	"hemacms/common"
	"hemacms/common/sql"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func LoginVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uri := r.RequestURI
		session := common.Sess().Load(r)
		userId, err := session.GetString("uid")
		common.Log().CheckErr("Session Get Error", err)
		_, err = r.Cookie("back-language")
		if err != nil {
			common.Log().Debug().Str("[back-language-reload]", "cn").Str("[Error]", err.Error()).Msg("back-stage language set fail")
			homeLanguage := &http.Cookie{
				Name:     "back-language",
				Value:    "cn",
				Path:     "/",
				HttpOnly: false,
				MaxAge:   80000,
			}
			http.SetCookie(w, homeLanguage)
		} else {
			// common.Log().Debug().Str("[back-language]", cookie.Value).Msg("back-stage language set success")
		}
		switch uri {
		case "/intendant/login":
			if userId != "" {
				fmt.Fprint(w, "<script>parent.window.location = '/intendant/index';</script>")
			}
		case "/intendant":
		case "/intendant/tabNoAuth":
			if userId == "" {
				fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":67}")
			}
		default:
			if userId == "" {
				if r.Method == "GET" {
					fmt.Fprint(w, "<script>parent.window.location = '/intendant/login';</script>")
					http.Redirect(w, r, "/intendant/login", http.StatusFound)
					return
				} else if r.Method == "POST" {
					fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":66}")
					return
				} else {
					fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":66}")
					return
				}
			} else if userId == "1" {
				fmt.Println("hi,你进来了")
			} else {
				//获取所有菜单
				fmt.Println("你的id", userId)
				rule := []sql.AuthRule{}
				if rows, found := common.Cache().CacheGet("allmenu"); found {
					rule = rows.([]sql.AuthRule)
				} else {
					sqls := "SELECT * FROM hm_auth_rule ORDER BY sort ASC"
					err := common.Sql().Select(&rule, sqls)
					common.Log().CheckErr("Sql Error", err)
					common.Cache().CacheSetAlwaysTime("allmenu", rule)
				}
				//获取用户的用户组id
				roleId := ""
				if roleIdRow, roleIdFound := common.Cache().CacheGet("roleId" + userId); roleIdFound {
					roleId = roleIdRow.(string)
				} else {
					access := sql.AuthRoleAccess{}
					accessSql := "SELECT * FROM hm_auth_role_access WHERE uid = ?"
					err := common.Sql().Get(&access, accessSql, userId)
					common.Log().CheckErr("Sql Error", err)
					roleId = strconv.Itoa(access.RoleId)
					common.Cache().CacheSetAlwaysTime("roleId"+userId, roleId)
				}
				rulerz := "false"
				formatUri := common.Base().FormatUrl(uri)
				fmt.Println(formatUri)
				for _, v := range rule {
					if formatUri == "/intendant/"+v.Url {
						fmt.Println(v.Id, "草拟吗")
						rulerz = "true"
						//获取用户权限id组
						auths := ""
						if row, fd := common.Cache().CacheGet("auth" + roleId); fd {
							auths = row.(string)
							fmt.Println(auths)

						} else {
							fmt.Println("找不到rules权限")
							role := sql.AuthRole{}
							roleSql := "SELECT * FROM hm_auth_role WHERE id = ?"
							err = common.Sql().Get(&role, roleSql, roleId)
							common.Log().CheckErr("Sql Error", err)
							auths = role.Rules
							common.Cache().CacheSetAlwaysTime("auth"+roleId, auths)
						}
						rules := strings.Split(auths, ",")
						renzheng := "false"
						for _, rv := range rules {
							intv := strconv.Itoa(v.Id)
							if rv == intv {
								renzheng = "true"
							}
						}
						if renzheng == "true" {
							fmt.Println("你拥有权限")
						} else {
							if r.Method == "GET" {
								common.Base().Template(w, r, nil, "./view/admin/index/auth.html")
							} else if r.Method == "POST" {
								fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":66}")
							} else {
								fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":66}")
							}
							return
						}
					}
				}
				if rulerz == "true" {
					fmt.Println("您拥有权限")
				} else {
					if r.Method == "GET" {
						common.Base().Template(w, r, nil, "./view/admin/index/auth.html")
					} else if r.Method == "POST" {
						fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":66}")
					} else {
						fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":66}")
					}
					return
				}
			}
		}

		// if err != nil {
		// 	http.Error(w, http.StatusText(404), 404)
		// 	return
		// }
		ctx := context.WithValue(r.Context(), "ExcuteTime", time.Now().Unix())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//没有路由的情况下
func NoFound(w http.ResponseWriter, r *http.Request) {
	common.Base().Template(w, r, nil, "./view/admin/index/nofound.html")
}

// func formatUrl(s string) string {
// 	uri := s
// 	straddress := ""
// 	for _, v := range uri { // i 是字符的字节位置，v 是字符的拷贝
// 		ss := fmt.Sprintf("%c", v) // 输出单个字符
// 		yz := common.Regexp().Id(ss)
// 		if yz == false {
// 			straddress = straddress + ss
// 		}
// 	}
// 	result := strings.TrimRight(straddress, "/")
// 	return result
// }
