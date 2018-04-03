package admin

import (
	"fmt"
	"hemacms/common"
	"hemacms/common/sql"
	"net/http"
	"strconv"
	"strings"
)

type IndexCtrl struct {
	common.BaseCtrl
}

func IndexCtrlObject() *IndexCtrl {
	return &IndexCtrl{}
}
func (c *IndexCtrl) Skip(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/intendant/login", http.StatusFound)
}
func (c *IndexCtrl) Index(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	session := c.Sess().Load(r)
	username, err := session.GetString("username")
	c.Log().CheckErr("Session Get Error", err)
	data["username"] = username
	c.Template(w, r, data, "./view/admin/index/index.html")
}
func (c *IndexCtrl) Home(w http.ResponseWriter, r *http.Request) {
	c.Template(w, r, nil, "./view/admin/index/home.html")
}

// func (c *IndexCtrl) GetMenu(w http.ResponseWriter, r *http.Request) {
// 	// pid := r.PostFormValue("pid")
// 	// intpid, _ := strconv.Atoi(pid)
// 	sqls := "SELECT * FROM hm_auth_rule"
// 	rule := []sql.AuthRule{}
// 	err := c.Sql().Select(&rule, sqls)
// 	if err != nil {
// 		c.Log().Debug().Err(err).Msg("错误")
// 		c.ResponseJson(4, err.Error(), w, r)
// 	}
// 	ar := sql.RecursiveMenu(rule, 0, 0)
// 	// fmt.Println(rule)
// 	// for k, v := range rule {
// 	// 	srule := []sqlm.AuthRule{}
// 	// 	err = c.Sql().Select(&srule, sqls, v.Id)
// 	// 	if err != nil {
// 	// 		c.Log().Debug().Err(err).Msg("错误")
// 	// 		c.ResponseJson(4, err.Error(), w, r)
// 	// 	}
// 	// 	for tk, _ := range srule {
// 	// 		rule[k].Children = append(rule[k].Children, srule[tk])
// 	// 	}
// 	// }
// 	// fmt.Println(ar)
// 	fmt.Fprint(w, c.RowsJson(ar))
// }

//tab权限
func (c *IndexCtrl) TabNoAuth(w http.ResponseWriter, r *http.Request) {
	uri := r.PostFormValue("uri")
	session := c.Sess().Load(r)
	userId, err := session.GetString("uid")
	c.Log().CheckErr("Session Get Error", err)
	if userId == "" {
		// fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":67}")
		return
	} else if userId == "1" {
		fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":88}")
		return
	} else {
		//获取所有菜单
		rule := []sql.AuthRule{}
		if rows, found := c.Cache().CacheGet("allmenu"); found {
			rule = rows.([]sql.AuthRule)
		} else {
			sqls := "SELECT * FROM hm_auth_rule ORDER BY sort ASC"
			err := c.Sql().Select(&rule, sqls)
			c.Log().CheckErr("Sql Error", err)
			c.Cache().CacheSetAlwaysTime("allmenu", rule)
		}
		//获取用户的用户组id
		roleId := ""
		if roleIdRow, roleIdFound := c.Cache().CacheGet("roleId" + userId); roleIdFound {
			roleId = roleIdRow.(string)
		} else {
			access := sql.AuthRoleAccess{}
			accessSql := "SELECT * FROM hm_auth_role_access WHERE uid = ?"
			err := c.Sql().Get(&access, accessSql, userId)
			c.Log().CheckErr("Sql Error", err)
			roleId = strconv.Itoa(access.RoleId)
			c.Cache().CacheSetAlwaysTime("roleId"+userId, roleId)
		}
		rulerz := "false"
		formatUri := c.FormatUrl(uri)
		for _, v := range rule {
			if formatUri == "/intendant/"+v.Url {
				rulerz = "true"
				//获取用户权限id组
				auths := ""
				if row, fd := c.Cache().CacheGet("auth" + roleId); fd {
					auths = row.(string)
				} else {
					role := sql.AuthRole{}
					roleSql := "SELECT * FROM hm_auth_role WHERE id = ?"
					err = c.Sql().Get(&role, roleSql, roleId)
					c.Log().CheckErr("Sql Error", err)
					auths = role.Rules
					c.Cache().CacheSetAlwaysTime("auth"+roleId, auths)
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
					fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":88}")
					return
				} else {
					fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":66}")
					return
				}
			}
		}
		if rulerz == "true" {
			fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":88}")
			return
		} else {
			fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":66}")
			return
		}
	}
}
