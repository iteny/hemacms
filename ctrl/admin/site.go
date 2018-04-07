package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hemacms/common"
	"hemacms/common/sql"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
)

type SiteCtrl struct {
	common.BaseCtrl
}

func SiteCtrlObject() *SiteCtrl {
	return &SiteCtrl{}
}

/**
 * @description 菜单管理页面
 * @English	menu management page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) Menu(w http.ResponseWriter, r *http.Request) {
	c.Template(w, r, nil, "./view/admin/site/menu.html")
}

/**
 * @description 获取菜单
 * @English	get menus
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) GetMenu(w http.ResponseWriter, r *http.Request) {
	rule := []sql.AuthRule{}
	if rows, found := c.Cache().CacheGet("allmenu"); found {
		rule = rows.([]sql.AuthRule)
	} else {
		sqls := "SELECT * FROM hm_auth_rule ORDER BY sort ASC"
		err := c.Sql().Select(&rule, sqls)
		if err != nil {
			c.Log().Debug().Err(err).Msg("Error")
			c.ResponseJson(4, err.Error(), w, r)
		}
		c.Cache().CacheSetAlwaysTime("allmenu", rule)
	}
	ar := sql.RecursiveMenu(rule, 0, 0)
	fmt.Fprint(w, c.RowsJson(ar))
}

/**
 * @description 菜单批量排序
 * @English	batch sorting menus
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) SortMenu(w http.ResponseWriter, r *http.Request) {
	sortMenu := make(map[string]string, 0)
	defer r.Body.Close()
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.ResponseJson(4, ""+err.Error(), w, r)
		c.Log().Debug().Err(err).Msg("Error")
	}
	json.Unmarshal(result, &sortMenu)
	var menuk []string
	for k, _ := range sortMenu {
		menuk = append(menuk, k)
	}
	var bf bytes.Buffer
	bf.WriteString("UPDATE hm_auth_rule SET sort = CASE id ")
	// menu := make(map[string]string, 0)
	for _, v := range menuk {
		mid := c.Regexp().Id(v)
		msort := c.Regexp().Sort(sortMenu[v])
		switch false {
		case mid:
			c.ResponseJson(11, "", w, r)
			return
		case msort:
			c.ResponseJson(12, "", w, r)
			return
		default:
			bf.WriteString(fmt.Sprintf("WHEN %v THEN %v ", v, sortMenu[v]))
		}
	}
	// fmt.Println(menu)
	//实现了implode的功能
	var buffer bytes.Buffer
	for _, v := range menuk {
		buffer.WriteString(v)
		buffer.WriteString(",")
	}
	ids := strings.Trim(buffer.String(), ",")
	bf.WriteString(fmt.Sprintf("END WHERE id IN (%v)", ids))
	// fmt.Println(ids)
	// fmt.Println(bf.String())
	tx := c.Sql().MustBegin()
	tx.MustExec(bf.String())
	err = tx.Commit()
	if err != nil {
		c.ResponseJson(4, ""+err.Error(), w, r)
	} else {
		c.Cache().CacheDel("allmenu")
		c.ResponseJson(1, "", w, r)
	}
}

/**
 * @description 获取图标
 * @English	get icons
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) IconsCls(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile("./static/common/dist/icons.css")
	if err != nil {
		c.Log().Debug().Err(err).Msg("Error")
	}
	var ss []string
	s := string(dat)
	// fmt.Printf("%#v", s)
	for _, v := range strings.Split(s, "\n") {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		ss = append(ss, v)
	}
	fmt.Println(ss)
	// s := strings.Split(dat, "}")
	fmt.Fprint(w, c.RowsJson(ss))
}

/**
 * @description 添加菜单页面
 * @English	add menu page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) AddMenu(w http.ResponseWriter, r *http.Request) {
	rule := []sql.AuthRule{}
	if rows, found := c.Cache().CacheGet("allmenu"); found {
		rule = rows.([]sql.AuthRule)
	} else {
		sqls := "SELECT * FROM hm_auth_rule"
		err := c.Sql().Select(&rule, sqls)
		if err != nil {
			c.Log().Debug().Err(err).Msg("Error")
			c.ResponseJson(4, err.Error(), w, r)
		}
		c.Cache().CacheSetAlwaysTime("allmenu", rule)
	}
	data := make(map[string]interface{})
	ar := sql.RecursiveMenuLevel(rule, 0, 0)
	data["json"] = c.RowsJson(ar)
	pid := chi.URLParam(r, "menuPid")
	data["currentUrl"] = r.RequestURI
	data["pid"] = pid
	c.Template(w, r, data, "./view/admin/site/addMenu.html")
}

/**
 * @description 添加菜单提交
 * @English	add menu submit
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) AddMenuSubmit(w http.ResponseWriter, r *http.Request) {
	pid, name, url, en := r.PostFormValue("pid"), r.PostFormValue("name"), r.PostFormValue("url"), r.PostFormValue("en")
	sort, icon, isshow := r.PostFormValue("sort"), r.PostFormValue("icon"), r.PostFormValue("isshow")
	isPid, isName := c.Regexp().Id(pid), c.Regexp().MenuName(name)
	isSort, isIcon, isIsshow := c.Regexp().Sort(sort), c.Regexp().Icon(icon), c.Regexp().Status(isshow)
	isUrl, isEn := c.Regexp().UrlAdmin(url), c.Regexp().English(en)
	switch false {
	case isPid:
		c.ResponseJson(11, "", w, r)
		return
	case isName:
		c.ResponseJson(12, "", w, r)
		return
	case isSort:
		c.ResponseJson(13, "", w, r)
		return
	case isIcon:
		c.ResponseJson(14, "", w, r)
		return
	case isIsshow:
		c.ResponseJson(15, "", w, r)
		return
	case isUrl:
		c.ResponseJson(16, "", w, r)
		return
	case isEn:
		c.ResponseJson(17, "", w, r)
		return
	default:
		rule := sql.AuthRule{}
		err := c.Sql().Get(&rule, "SELECT * FROM hm_auth_rule WHERE name = ? OR url = ? OR en = ?", name, url, en)
		if err != nil {
			sqls := "INSERT INTO hm_auth_rule(url,name,pid,isshow,sort,icon,level,en) VALUES(?,?,?,?,?,?,?,?)"
			tx := c.Sql().MustBegin()
			tx.MustExec(sqls, url, name, pid, isshow, sort, icon, 1, en)
			err = tx.Commit()
			if err != nil {
				c.Log().Debug().Err(err).Msg("Error")
				c.ResponseJson(4, err.Error(), w, r)
			} else {
				c.Cache().CacheDel("allmenu")
				c.ResponseJson(1, "", w, r)
			}
		} else {
			if rule.Name == name {
				c.ResponseJson(21, "", w, r)
				return
			} else if rule.Url == url {
				c.ResponseJson(22, "", w, r)
				return
			} else if rule.En == en {
				c.ResponseJson(23, "", w, r)
				return
			} else {
				c.ResponseJson(44, "", w, r)
				return
			}
		}
	}
}

/**
 * @description 修改菜单页面
 * @English	edit menu page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) EditMenu(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	id := chi.URLParam(r, "menuId")
	ruleSingle := sql.AuthRule{}
	err := c.Sql().Get(&ruleSingle, "SELECT * FROM hm_auth_rule WHERE id = ?", id)
	if err != nil {
		c.Log().Debug().Err(err).Msg("Error")
		c.ResponseJson(4, err.Error(), w, r)
	}
	data["menu"] = ruleSingle
	allrule := []sql.AuthRule{}
	if rows, found := c.Cache().CacheGet("allmenu"); found {
		allrule = rows.([]sql.AuthRule)
	} else {
		sqls := "SELECT * FROM hm_auth_rule"
		err := c.Sql().Select(&allrule, sqls)
		if err != nil {
			c.Log().Debug().Err(err).Msg("Error")
			c.ResponseJson(4, err.Error(), w, r)
		}
		c.Cache().CacheSetAlwaysTime("allmenu", allrule)
	}
	ar := sql.RecursiveMenuLevel(allrule, 0, 0)
	data["json"] = c.RowsJson(ar)
	data["currentUrl"] = r.RequestURI
	c.Template(w, r, data, "./view/admin/site/editMenu.html")
}

/**
 * @description 修改菜单提交
 * @English	edit menu submit
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) EditMenuSubmit(w http.ResponseWriter, r *http.Request) {
	pid, name, url, en := r.PostFormValue("pid"), r.PostFormValue("name"), r.PostFormValue("url"), r.PostFormValue("en")
	sort, icon, isshow, id := r.PostFormValue("sort"), r.PostFormValue("icon"), r.PostFormValue("isshow"), r.PostFormValue("id")
	isPid, isName, isId := c.Regexp().Id(pid), c.Regexp().MenuName(name), c.Regexp().Id(id)
	isSort, isIcon, isIsshow := c.Regexp().Sort(sort), c.Regexp().Icon(icon), c.Regexp().Status(isshow)
	isUrl, isEn := c.Regexp().UrlAdmin(url), c.Regexp().English(en)
	switch false {
	case isPid:
		c.ResponseJson(11, "", w, r)
		return
	case isName:
		c.ResponseJson(12, "", w, r)
		return
	case isSort:
		c.ResponseJson(13, "", w, r)
		return
	case isIcon:
		c.ResponseJson(14, "", w, r)
		return
	case isIsshow:
		c.ResponseJson(15, "", w, r)
		return
	case isUrl:
		c.ResponseJson(16, "", w, r)
		return
	case isEn:
		c.ResponseJson(17, "", w, r)
		return
	case isId:
		c.ResponseJson(18, "", w, r)
		return
	default:
		rule := sql.AuthRule{}
		err := c.Sql().Get(&rule, "SELECT * FROM hm_auth_rule WHERE id <> ? AND (name = ? OR url = ? OR en = ?)", id, name, url, en)
		if err != nil {
			sqls := "UPDATE hm_auth_rule SET url=?,name=?,pid=?,isshow=?,sort=?,icon=?,level=?,en=? WHERE id = ?"
			tx := c.Sql().MustBegin()
			tx.MustExec(sqls, url, name, pid, isshow, sort, icon, 1, en, id)
			err := tx.Commit()
			if err != nil {
				c.Log().Debug().Err(err).Msg("Error")
				c.ResponseJson(4, err.Error(), w, r)
			} else {
				c.Cache().CacheDel("allmenu")
				c.ResponseJson(1, "", w, r)
			}
		} else {
			if rule.Name == name {
				c.ResponseJson(21, "", w, r)
				return
			} else if rule.Url == url {
				c.ResponseJson(22, "", w, r)
				return
			} else if rule.En == en {
				c.ResponseJson(23, "", w, r)
				return
			} else {
				c.ResponseJson(44, "", w, r)
				return
			}
		}
	}
}

/**
 * @description 删除菜单提交
 * @English	delete menu submit
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) DelMenuSubmit(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	isId := c.Regexp().Id(id)
	if isId == false {
		c.ResponseJson(11, "", w, r)
		return
	} else {
		rule := []sql.AuthRule{}
		if rows, found := c.Cache().CacheGet("allmenu"); found {
			rule = rows.([]sql.AuthRule)
		} else {
			sqls := "SELECT * FROM hm_auth_rule"
			err := c.Sql().Select(&rule, sqls)
			if err != nil {
				c.Log().Debug().Err(err).Msg("错���")
				c.ResponseJson(4, err.Error(), w, r)
			}
			c.Cache().CacheSetAlwaysTime("allmenu", rule)
		}
		ids := sql.RecursiveMenuId(rule, id)
		delSql := fmt.Sprintf("DELETE FROM hm_auth_rule WHERE id IN(%v)", ids)
		tx := c.Sql().MustBegin()
		tx.MustExec(delSql)
		err := tx.Commit()
		if err != nil {
			c.Log().Debug().Err(err).Msg("Error")
			c.ResponseJson(4, err.Error(), w, r)
		} else {
			c.Cache().CacheDel("allmenu")
			c.ResponseJson(1, "", w, r)
		}
	}
}

/**
 * @description 系统设置页面
 * @English	system site page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) System(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["sqlType"] = c.Config().Value("sql", "sqlType")
	data["port"] = c.Config().Int("servSet", "port")
	data["readTimeout"] = c.Config().Int("servSet", "readTimeout")
	data["writeTimeout"] = c.Config().Int("servSet", "writeTimeout")
	data["language"] = c.Config().Value("cookie", "language")
	c.Template(w, r, data, "./view/admin/site/system.html")
}

/**
 * @description 系统设置修改
 * @English	edit system site
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) EditSystem(w http.ResponseWriter, r *http.Request) {
	sqlType, port, language := r.PostFormValue("sqlType"), r.PostFormValue("port"), r.PostFormValue("language")
	readTimeout, writeTimeout := r.PostFormValue("readTimeout"), r.PostFormValue("writeTimeout")
	isSqlType, isPort, isLanguage := c.Regexp().SqlType(sqlType), c.Regexp().Port(port), c.Regexp().English(language)
	isReadTimeout, isWriteTimeout := c.Regexp().ReadTimeout(readTimeout), c.Regexp().WriteTimeout(writeTimeout)
	switch false {
	case isSqlType:
		c.ResponseJson(11, "", w, r)
		return
	case isPort:
		c.ResponseJson(12, "", w, r)
		return
	case isReadTimeout:
		c.ResponseJson(13, "", w, r)
		return
	case isWriteTimeout:
		c.ResponseJson(14, "", w, r)
		return
	case isLanguage:
		c.ResponseJson(15, "", w, r)
		return
	default:
		c.Config().SetValue("sql", "sqlType", sqlType)
		c.Config().SetValue("servSet", "port", port)
		c.Config().SetValue("servSet", "readTimeout", readTimeout)
		c.Config().SetValue("servSet", "writeTimeout", writeTimeout)
		c.Config().SetValue("cookie", "language", language)
		err := c.Config().Save("./ini/hemacms.ini")
		if err != nil {
			c.Log().Debug().Err(err).Msg("Error")
			c.ResponseJson(4, err.Error(), w, r)
		} else {
			homeLanguage := &http.Cookie{
				Name:     "back-language",
				Value:    language,
				Path:     "/",
				HttpOnly: false,
				MaxAge:   80000,
			}
			http.SetCookie(w, homeLanguage)
			err = c.Config().Reload()
			if err != nil {
				c.Log().Debug().Err(err).Msg("Error")
				c.ResponseJson(4, err.Error(), w, r)
			} else {
				c.ResponseJson(1, "", w, r)
			}
		}
	}
}

/**
 * @description 用户管理页面
 * @English	user management page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) User(w http.ResponseWriter, r *http.Request) {
	c.Template(w, r, nil, "./view/admin/site/user.html")
}

/**
 * @description 获取用户数据
 * @English	get users data
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) GetUser(w http.ResponseWriter, r *http.Request) {
	array := make([]sql.User, 0)
	jsonUser := ""
	page, row := r.PostFormValue("page"), r.PostFormValue("rows")
	username, dateFrom, dateTo := r.PostFormValue("username"), r.PostFormValue("dateFrom"), r.PostFormValue("dateTo")
	isUsername, isDateFrom, isDateTo := c.Regexp().Username(username), c.Regexp().Date(dateFrom), c.Regexp().Date(dateTo)
	switch {
	case isUsername == false && username != "":
		jsonUser = "{\"total\":0,\"rows\":[]}"
		// fmt.Println("username")
	case isDateFrom == false && dateFrom != "":
		jsonUser = "{\"total\":0,\"rows\":[]}"
		fmt.Println("datefrom")
	case isDateTo == false && dateTo != "":
		jsonUser = "{\"total\":0,\"rows\":[]}"
		fmt.Println("dateto")
	default:
		ipage, _ := strconv.Atoi(page)
		irow, _ := strconv.Atoi(row)
		first := irow * (ipage - 1)
		timeLayout := "2006-01-02 15:04:05"  //转化所需模板
		loc, _ := time.LoadLocation("Local") //重要：获取时区
		from, _ := time.ParseInLocation(timeLayout, dateFrom, loc)
		to, _ := time.ParseInLocation(timeLayout, dateTo, loc)
		fromTime, toTime := from.Unix(), to.Unix()
		addsql := ""
		if username != "" {
			addsql = "username LIKE '%" + username + "%' AND "
		}
		if dateFrom != "" {
			addsql = addsql + fmt.Sprintf("create_time>='%v' AND ", fromTime)
		}
		if dateTo != "" {
			addsql = addsql + fmt.Sprintf("create_time<='%v' AND ", toTime)
		}
		if addsql != "" {
			addsql = "WHERE " + strings.Trim(addsql, "AND ")
		}
		fmt.Println(addsql)
		if rows, found := c.Cache().CacheGet("allUser"); found {
			// array = rows.([]sql.User)
			jsonUser = rows.(string)
		} else {
			userSql := "SELECT * FROM hm_user " + addsql + " Limit ?,?"
			roleSql := "SELECT * FROM hm_auth_role"
			accessSql := "SELECT uid,role_id FROM hm_auth_role_access"
			user := []sql.User{}
			role := []sql.AuthRole{}
			access := []sql.AuthRoleAccess{}
			err := c.Sql().Select(&user, userSql, first, row)
			if err != nil {
				c.Log().Debug().Err(err).Msg("Error1")
				c.ResponseJson(4, err.Error(), w, r)
			}
			err = c.Sql().Select(&role, roleSql)
			if err != nil {
				c.Log().Debug().Err(err).Msg("Error")
				c.ResponseJson(4, err.Error(), w, r)
			}
			err = c.Sql().Select(&access, accessSql)
			if err != nil {
				c.Log().Debug().Err(err).Msg("Error")
				c.ResponseJson(4, err.Error(), w, r)
			}
			for k, v := range user {
				array = append(array, user[k])
				for _, av := range access {
					if v.Id == av.Uid {
						for gk, gv := range role {
							if av.RoleId == gv.Id {
								array[k].Role = role[gk].Name
								array[k].RoleEn = role[gk].En
							}
						}

					}
				}
			}
			count := len(array)
			if count != 0 {
				jsonUser = "{\"total\":" + strconv.Itoa(count) + ",\"rows\":" + c.RowsJson(array) + "}"
			} else {
				jsonUser = "{\"total\":0,\"rows\":[]}"
			}
			// c.Cache().CacheSetAlwaysTime("allUser", jsonUser)
		}
	}
	fmt.Fprint(w, jsonUser)
}

/**
 * @description 删除单个用户
 * @English	delete single user
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) DelUser(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	isId := c.Regexp().Id(id)
	if isId == false {
		c.ResponseJson(11, "", w, r)
	} else {
		if id == "1" {
			c.ResponseJson(12, "", w, r)
			return
		} else {
			userSql := "DELETE FROM hm_user WHERE id = ?"
			accessSql := "DELETE FROM hm_auth_role_access WHERE uid = ?"
			tx := c.Sql().MustBegin()
			tx.MustExec(userSql, id)
			tx.MustExec(accessSql, id)
			err := tx.Commit()
			if err != nil {
				c.Log().Debug().Err(err).Msg("Error")
				c.ResponseJson(4, err.Error(), w, r)
			} else {
				// c.Cache().CacheDel("allUser")
				c.ResponseJson(1, "", w, r)
			}
		}
	}
}

/**
 * @description 批量删除用户
 * @English	batch delete users
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) BatchDelUser(w http.ResponseWriter, r *http.Request) {
	ids := r.PostFormValue("ids")
	isIds := c.Regexp().IdGroup(ids)
	if !isIds {
		c.ResponseJson(11, "", w, r)
	} else {
		id := strings.Split(ids, ",")
		for _, v := range id {
			if v == "1" {
				c.ResponseJson(12, "", w, r)
				return
			}
		}
		userSql := fmt.Sprintf("DELETE FROM hm_user WHERE id IN (%v)", ids)
		accessSql := fmt.Sprintf("DELETE FROM hm_auth_role_access WHERE uid IN (%v)", ids)
		tx := c.Sql().MustBegin()
		tx.MustExec(userSql)
		tx.MustExec(accessSql)
		err := tx.Commit()
		if err != nil {
			c.Log().Debug().Err(err).Msg("Error")
			c.ResponseJson(4, err.Error(), w, r)
		} else {
			c.Cache().CacheDel("allUser")
			c.ResponseJson(1, "", w, r)
		}
	}
}

/**
 * @description 添加用户页面
 * @English	add user page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) AddUser(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	roleSql := "select * from hm_auth_role"
	role := []sql.AuthRole{}
	err := c.Sql().Select(&role, roleSql)
	if err != nil {
		c.Log().Debug().Err(err).Msg("Error")
		c.ResponseJson(4, err.Error(), w, r)
	}
	data["currentUrl"] = r.RequestURI
	data["role"] = c.RowsJson(role)
	c.Template(w, r, data, "./view/admin/site/addUser.html")
}

/**
 * @description 添加用户提交
 * @English	add user submit
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) AddUserSubmit(w http.ResponseWriter, r *http.Request) {
	username, password, passworded := r.PostFormValue("username"), r.PostFormValue("password"), r.PostFormValue("passworded")
	nickname, email := r.PostFormValue("nickname"), r.PostFormValue("email")
	remark, role, status := r.PostFormValue("remark"), r.PostFormValue("role"), r.PostFormValue("status")
	isUsername, isPassword := c.Regexp().Username(username), c.Regexp().Password(password)
	isNickname, isEmail, isRemark := c.Regexp().Nickname(nickname), c.Regexp().Email(email), c.Regexp().Remark(remark)
	isRole, isStatus := c.Regexp().Id(role), c.Regexp().Status(status)
	switch false {
	case isUsername:
		c.ResponseJson(11, "", w, r)
		return
	case isPassword:
		c.ResponseJson(12, "", w, r)
		return
	case isNickname:
		c.ResponseJson(13, "", w, r)
		return
	case isEmail:
		c.ResponseJson(14, "", w, r)
		return
	case isRemark:
		c.ResponseJson(15, "", w, r)
		return
	case isRole:
		c.ResponseJson(16, "", w, r)
		return
	case isStatus:
		c.ResponseJson(17, "", w, r)
		return
	default:
		if password != passworded {
			c.ResponseJson(18, "", w, r)
			return
		} else {
			user := sql.User{}
			err := c.Sql().Get(&user, "SELECT * FROM hm_user WHERE username = ? OR nickname = ? OR email = ?", username, nickname, email)
			if err != nil {
				ip, _, _ := net.SplitHostPort(r.RemoteAddr)
				if ip == "::1" {
					ip = "127.0.0.1"
				}
				userSql := "INSERT INTO hm_user(username,password,nickname,email,remark,status,create_time,create_ip) VALUES(?,?,?,?,?,?,?,?)"
				accessSql := "INSERT INTO hm_auth_role_access(uid,role_id) VALUES(?,?)"
				tx := c.Sql().MustBegin()
				id, err := tx.MustExec(userSql, username, c.Sha1PlusMd5(password), nickname, email, remark, status, time.Now().Unix(), ip).LastInsertId()
				if err != nil {
					c.Log().Debug().Err(err).Msg("Error")
					c.ResponseJson(4, err.Error(), w, r)
				}
				tx.MustExec(accessSql, id, role)
				err = tx.Commit()
				if err != nil {
					c.Log().Debug().Err(err).Msg("Error")
					c.ResponseJson(4, err.Error(), w, r)
				} else {
					c.Cache().CacheDel("allUser")
					c.ResponseJson(1, "", w, r)
				}
			} else {
				if user.Username == username {
					c.ResponseJson(21, "", w, r)
					return
				} else if user.Nickname == nickname {
					c.ResponseJson(22, "", w, r)
					return
				} else if user.Email == email {
					c.ResponseJson(23, "", w, r)
					return
				} else {
					c.ResponseJson(44, "", w, r)
					return
				}
			}
		}
	}
}

/**
 * @description 修改用户页面
 * @English	change user page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) EditUser(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	id := chi.URLParam(r, "userId")
	userSingle := sql.User{}
	accessSingle := sql.AuthRoleAccess{}
	err := c.Sql().Get(&userSingle, "SELECT * FROM hm_user WHERE id = ?", id)
	if err != nil {
		c.Log().Debug().Err(err).Msg("Error")
		c.ResponseJson(4, err.Error(), w, r)
	}
	err = c.Sql().Get(&accessSingle, "SELECT * FROM hm_auth_role_access WHERE uid = ?", id)
	if err != nil {
		c.Log().Debug().Err(err).Msg("Error")
		c.ResponseJson(4, err.Error(), w, r)
	}
	data["user"] = userSingle
	data["roleme"] = accessSingle
	roleSql := "select * from hm_auth_role"
	role := []sql.AuthRole{}
	err = c.Sql().Select(&role, roleSql)
	if err != nil {
		c.Log().Debug().Err(err).Msg("Error")
		c.ResponseJson(4, err.Error(), w, r)
	}
	data["currentUrl"] = r.RequestURI
	data["role"] = c.RowsJson(role)
	c.Template(w, r, data, "./view/admin/site/editUser.html")
}

/**
 * @description 修改用户提交
 * @English	change user submit
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) EditUserSubmit(w http.ResponseWriter, r *http.Request) {
	username, password, passworded := r.PostFormValue("username"), r.PostFormValue("password"), r.PostFormValue("passworded")
	nickname, email, id := r.PostFormValue("nickname"), r.PostFormValue("email"), r.PostFormValue("id")
	remark, role, status := r.PostFormValue("remark"), r.PostFormValue("role"), r.PostFormValue("status")
	isUsername, isPassword, isId := c.Regexp().Username(username), c.Regexp().Password(password), c.Regexp().Id(id)
	isNickname, isEmail, isRemark := c.Regexp().Nickname(nickname), c.Regexp().Email(email), c.Regexp().Remark(remark)
	isRole, isStatus := c.Regexp().Id(role), c.Regexp().Status(status)
	switch false {
	case isUsername:
		c.ResponseJson(11, "", w, r)
		return
	case isPassword:
		c.ResponseJson(12, "", w, r)
		return
	case isNickname:
		c.ResponseJson(13, "", w, r)
		return
	case isEmail:
		c.ResponseJson(14, "", w, r)
		return
	case isRemark:
		c.ResponseJson(15, "", w, r)
		return
	case isRole:
		c.ResponseJson(16, "", w, r)
		return
	case isStatus:
		c.ResponseJson(17, "", w, r)
		return
	case isId:
		c.ResponseJson(19, "", w, r)
		return
	default:
		if password != passworded {
			c.ResponseJson(18, "", w, r)
			return
		} else {
			user := sql.User{}
			err := c.Sql().Get(&user, "SELECT * FROM hm_user WHERE id <> ? AND (username = ? OR nickname = ? OR email = ?)", id, username, nickname, email)
			if err != nil {
				ip, _, _ := net.SplitHostPort(r.RemoteAddr)
				if ip == "::1" {
					ip = "127.0.0.1"
				}
				userSql := "UPDATE hm_user SET username=?,password=?,nickname=?,email=?,remark=?,status=? WHERE id = ?"
				accessSql := "UPDATE hm_auth_role_access SET role_id=? WHERE uid = ?"
				tx := c.Sql().MustBegin()
				tx.MustExec(userSql, username, c.Sha1PlusMd5(password), nickname, email, remark, status, id)
				tx.MustExec(accessSql, role, id)
				err = tx.Commit()
				if err != nil {
					c.Log().Debug().Err(err).Msg("Error")
					c.ResponseJson(4, err.Error(), w, r)
				} else {
					c.Cache().CacheDel("allUser")
					c.ResponseJson(1, "", w, r)
				}
			} else {
				if user.Username == username {
					c.ResponseJson(21, "", w, r)
					return
				} else if user.Nickname == nickname {
					c.ResponseJson(22, "", w, r)
					return
				} else if user.Email == email {
					c.ResponseJson(23, "", w, r)
					return
				} else {
					c.ResponseJson(44, "", w, r)
					return
				}
			}
		}
	}
}

/**
 * @description 角色管理页面
 * @English	role management page
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) Role(w http.ResponseWriter, r *http.Request) {
	c.Template(w, r, nil, "./view/admin/site/role.html")
}

/**
 * @description 获取角色
 * @English	get roles
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) GetRole(w http.ResponseWriter, r *http.Request) {
	jsonRole := ""
	page, row := r.PostFormValue("page"), r.PostFormValue("rows")
	name := r.PostFormValue("name")
	isName := c.Regexp().Name(name)
	switch {
	case isName == false && name != "":
		jsonRole = "{\"total\":0,\"rows\":[]}"
	default:
		ipage, _ := strconv.Atoi(page)
		irow, _ := strconv.Atoi(row)
		first := irow * (ipage - 1)
		addsql := ""
		if name != "" {
			cookie, err := r.Cookie("back-language")
			if err != nil {
				common.Log().Debug().Str("[back-language-reload]", "cn").Str("[Error]", err.Error()).Msg("back-stage language set fail")
				language := "cn"
				language = c.Config().Value("cookie", "language")
				homeLanguage := &http.Cookie{
					Name:     "back-language",
					Value:    language,
					Path:     "/",
					HttpOnly: false,
					MaxAge:   80000,
				}
				http.SetCookie(w, homeLanguage)
			}
			switch cookie.Value {
			case "cn":
				addsql = "WHERE name LIKE '%" + name + "%'"
			case "en":
				addsql = "WHERE en LIKE '%" + name + "%'"
			default:
				addsql = "WHERE name LIKE '%" + name + "%'"
			}

		}
		if rows, found := c.Cache().CacheGet("allRole"); found {
			// array = rows.([]sql.User)
			jsonRole = rows.(string)
		} else {
			roleSql := "SELECT * FROM hm_auth_role " + addsql + " Limit ?,?"
			role := []sql.AuthRole{}
			err := c.Sql().Select(&role, roleSql, first, row)
			if err != nil {
				c.Log().Debug().Err(err).Msg("Error")
				c.ResponseJson(4, err.Error(), w, r)
			}
			count := len(role)
			if count != 0 {
				jsonRole = "{\"total\":" + strconv.Itoa(count) + ",\"rows\":" + c.RowsJson(role) + "}"
			} else {
				jsonRole = "{\"total\":0,\"rows\":[]}"
			}
			// c.Cache().CacheSetAlwaysTime("allUser", jsonUser)
		}
	}
	fmt.Fprint(w, jsonRole)
}

/**
 * @description 设置角色权限页面
 * @English	set role authority page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) SetRole(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "roleId")
	role := sql.AuthRole{}
	if rows, found := c.Cache().CacheGet("role" + id); found {
		role = rows.(sql.AuthRole)
	} else {
		sqls := "SELECT * FROM hm_auth_role WHERE id = ?"
		err := c.Sql().Get(&role, sqls, id)
		if err != nil {
			c.Log().Debug().Err(err).Msg("Error")
			c.ResponseJson(4, err.Error(), w, r)
		}
		c.Cache().CacheSetAlwaysTime("role"+id, role)
	}
	rule := []sql.AuthRule{}
	if rows, found := c.Cache().CacheGet("allmenu"); found {
		rule = rows.([]sql.AuthRule)
	} else {
		sqls := "SELECT * FROM hm_auth_rule ORDER BY sort ASC"
		err := c.Sql().Select(&rule, sqls)
		if err != nil {
			c.Log().Debug().Err(err).Msg("Error")
			c.ResponseJson(4, err.Error(), w, r)
		}
		c.Cache().CacheSetAlwaysTime("allmenu", rule)
	}
	ids := strings.Split(role.Rules, ",")
	for k, v := range rule {
		iss := 0
		for _, iv := range ids {
			vint, _ := strconv.Atoi(iv)
			if vint == v.Id {
				iss = 1
			}
		}
		rule[k].Checked = iss
	}
	array := sql.RecursiveMenu(rule, 0, 0)
	// for rk, rv := range array {
	// 	if len(rv.Children) != 0 {
	// 		for ck, cv := range rv.Children {
	// 			if len(cv.Children) != 0 {
	// 				for tk, tv := range cv.Children {
	// 					if len(tv.Children) == 0 {
	// 						for _, v := range ids {
	// 							vint, _ := strconv.Atoi(v)
	// 							if vint == tv.Id {
	// 								array[rk].Children[ck].Children[tk].Checked = 1
	// 							}
	// 						}
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	data := make(map[string]interface{})
	data["json"] = c.RowsJson(array)
	data["id"] = id
	data["roleName"] = role.Name
	data["roleEn"] = role.En
	c.Template(w, r, data, "./view/admin/site/setRole.html")
}

/**
* @description 设置角色权限提交
* @English	set role authority submit
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
*/
func (c *SiteCtrl) SetRoleSubmit(w http.ResponseWriter, r *http.Request) {
	id, ids := r.PostFormValue("id"), r.PostFormValue("ids")
	isIds := c.Regexp().IdGroup(ids)
	switch false {
	case isIds:
		c.ResponseJson(11, "", w, r)
		return
	default:
		sqls := "UPDATE hm_auth_role SET rules=? WHERE id = ?"
		tx := c.Sql().MustBegin()
		tx.MustExec(sqls, ids, id)
		err := tx.Commit()
		if err != nil {
			c.Log().Debug().Err(err).Msg("Error")
			c.ResponseJson(4, err.Error(), w, r)
		} else {
			c.Cache().CacheDel("role" + id)
			c.Cache().CacheDel("auth" + id)
			c.ResponseJson(1, "", w, r)
		}
	}
}

/**
 * @description 删除单个角色
 * @English	delete single role
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) DelRole(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	isId := c.Regexp().Id(id)
	if isId == false {
		c.ResponseJson(11, "", w, r)
	} else {
		if id == "1" {
			c.ResponseJson(12, "", w, r)
			return
		} else {
			roleSql := "DELETE FROM hm_auth_role WHERE id = ?"
			tx := c.Sql().MustBegin()
			tx.MustExec(roleSql, id)
			err := tx.Commit()
			if err != nil {
				c.Log().Debug().Err(err).Msg("Error")
				c.ResponseJson(4, err.Error(), w, r)
			} else {
				c.Cache().CacheDel("allrole")
				c.ResponseJson(1, "", w, r)
			}
		}
	}
}

/**
 * @description 批量删除角色
 * @English	batch delete roles
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) BatchDelRole(w http.ResponseWriter, r *http.Request) {
	ids := r.PostFormValue("ids")
	isIds := c.Regexp().IdGroup(ids)
	if !isIds {
		c.ResponseJson(11, "", w, r)
	} else {
		id := strings.Split(ids, ",")
		for _, v := range id {
			if v == "1" {
				c.ResponseJson(12, "", w, r)
				return
			}
		}
		roleSql := fmt.Sprintf("DELETE FROM hm_auth_role WHERE id IN (%v)", ids)
		tx := c.Sql().MustBegin()
		tx.MustExec(roleSql)
		err := tx.Commit()
		if err != nil {
			c.Log().Debug().Err(err).Msg("Error")
			c.ResponseJson(4, err.Error(), w, r)
		} else {
			c.Cache().CacheDel("allrole")
			c.ResponseJson(1, "", w, r)
		}
	}
}

/**
 * @description 添加角色页面
 * @English	add role page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) AddRole(w http.ResponseWriter, r *http.Request) {
	c.Template(w, r, nil, "./view/admin/site/addRole.html")
}

/**
 * @description 添加角色提交
 * @English	add role submit
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) AddRoleSubmit(w http.ResponseWriter, r *http.Request) {
	name, sort, remark, en := r.PostFormValue("name"), r.PostFormValue("sort"), r.PostFormValue("remark"), r.PostFormValue("en")
	isName, isSort, isRemark, isEn := c.Regexp().Name(name), c.Regexp().Sort(sort), c.Regexp().Remark(remark), c.Regexp().English(en)
	switch false {
	case isName:
		c.ResponseJson(11, "", w, r)
		return
	case isSort:
		c.ResponseJson(12, "", w, r)
		return
	case isRemark:
		c.ResponseJson(13, "", w, r)
		return
	case isEn:
		c.ResponseJson(14, "", w, r)
		return
	default:
		role := sql.AuthRole{}
		err := c.Sql().Get(&role, "SELECT * FROM hm_auth_role WHERE name = ? OR en = ?", name, en)
		if err != nil {
			sqls := "INSERT INTO hm_auth_role(name,sort,remark,en) VALUES(?,?,?,?)"
			tx := c.Sql().MustBegin()
			tx.MustExec(sqls, name, sort, remark, en)
			err = tx.Commit()
			if err != nil {
				c.Log().Debug().Err(err).Msg("Error")
				c.ResponseJson(4, err.Error(), w, r)
			} else {
				c.Cache().CacheDel("allrole")
				c.ResponseJson(1, "", w, r)
			}
		} else {
			if role.Name == name {
				c.ResponseJson(21, "", w, r)
				return
			} else if role.En == en {
				c.ResponseJson(22, "", w, r)
				return
			} else {
				c.ResponseJson(44, "", w, r)
				return
			}
		}
	}
}

/**
 * @description 修改角色页面
 * @English	edit role page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) EditRole(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "roleId")
	fmt.Println(id)
	role := sql.AuthRole{}
	sqls := "SELECT * FROM hm_auth_role WHERE id = ?"
	err := c.Sql().Get(&role, sqls, id)
	if err != nil {
		c.Log().Debug().Err(err).Msg("Error")
		c.ResponseJson(4, err.Error(), w, r)
	}
	data := make(map[string]interface{})
	data["role"] = role
	c.Template(w, r, data, "./view/admin/site/editRole.html")
}

/**
 * @description 修改角色提交
 * @English	edit role submit
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) EditRoleSubmit(w http.ResponseWriter, r *http.Request) {
	name, sort, remark, en, id := r.PostFormValue("name"), r.PostFormValue("sort"), r.PostFormValue("remark"), r.PostFormValue("en"), r.PostFormValue("id")
	isName, isSort, isRemark, isEn, isId := c.Regexp().Name(name), c.Regexp().Sort(sort), c.Regexp().Remark(remark), c.Regexp().English(en), c.Regexp().Id(id)
	switch false {
	case isName:
		c.ResponseJson(11, "", w, r)
		return
	case isSort:
		c.ResponseJson(12, "", w, r)
		return
	case isRemark:
		c.ResponseJson(13, "", w, r)
		return
	case isEn:
		c.ResponseJson(14, "", w, r)
		return
	case isId:
		c.ResponseJson(15, "", w, r)
		return
	default:
		role := sql.AuthRole{}
		err := c.Sql().Get(&role, "SELECT * FROM hm_auth_role WHERE id <> ? AND (name = ? OR en = ?)", id, name, en)
		if err != nil {
			sqls := "UPDATE hm_auth_role SET name=?,sort=?,remark=?,en=? WHERE id = ?"
			tx := c.Sql().MustBegin()
			tx.MustExec(sqls, name, sort, remark, en, id)
			err = tx.Commit()
			if err != nil {
				c.Log().Debug().Err(err).Msg("Error")
				c.ResponseJson(4, err.Error(), w, r)
			} else {
				c.Cache().CacheDel("allrole")
				c.ResponseJson(1, "", w, r)
			}
		} else {
			if role.Name == name {
				c.ResponseJson(21, "", w, r)
				return
			} else if role.En == en {
				c.ResponseJson(22, "", w, r)
				return
			} else {
				c.ResponseJson(44, "", w, r)
				return
			}
		}
	}
}

/**
 * @description 修改角色提交
 * @English	edit role submit
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) ColorSchemes(w http.ResponseWriter, r *http.Request) {
	c.Template(w, r, nil, "./view/admin/site/colorSchemes.html")
}

/**
 * @description 登录日志页面
 * @English	Login log page
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) LoginLog(w http.ResponseWriter, r *http.Request) {
	c.Template(w, r, nil, "./view/admin/site/loginLog.html")
}

/**
 * @description 获取登录日志
 * @English	Get login log
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) GetLoginLog(w http.ResponseWriter, r *http.Request) {
	jsonLog := ""
	page, row := r.PostFormValue("page"), r.PostFormValue("rows")
	username, dateFrom, dateTo := r.PostFormValue("username"), r.PostFormValue("dateFrom"), r.PostFormValue("dateTo")
	isUsername, isDateFrom, isDateTo := c.Regexp().Username(username), c.Regexp().Date(dateFrom), c.Regexp().Date(dateTo)
	switch {
	case isUsername == false && username != "":
		jsonLog = "{\"total\":0,\"rows\":[]}"
	case isDateFrom == false && dateFrom != "":
		jsonLog = "{\"total\":0,\"rows\":[]}"
		fmt.Println("datefrom")
	case isDateTo == false && dateTo != "":
		jsonLog = "{\"total\":0,\"rows\":[]}"
		fmt.Println("dateto")
	default:
		ipage, _ := strconv.Atoi(page)
		irow, _ := strconv.Atoi(row)
		first := irow * (ipage - 1)
		timeLayout := "2006-01-02 15:04:05"  //转化所需模板
		loc, _ := time.LoadLocation("Local") //重要：获取时区
		from, _ := time.ParseInLocation(timeLayout, dateFrom, loc)
		to, _ := time.ParseInLocation(timeLayout, dateTo, loc)
		fromTime, toTime := from.Unix(), to.Unix()
		addsql := ""
		if username != "" {
			addsql = "username LIKE '%" + username + "%' AND "
		}
		if dateFrom != "" {
			addsql = addsql + fmt.Sprintf("login_time>='%v' AND ", fromTime)
		}
		if dateTo != "" {
			addsql = addsql + fmt.Sprintf("login_time<='%v' AND ", toTime)
		}
		if addsql != "" {
			addsql = "WHERE " + strings.Trim(addsql, "AND ")
		}
		if rows, found := c.Cache().CacheGet("allLoginLog"); found {
			// array = rows.([]sql.User)
			jsonLog = rows.(string)
		} else {
			logSql := "SELECT * FROM hm_login_log " + addsql + " ORDER BY id DESC Limit ?,?"
			totalSql := "SELECT id FROM hm_login_log " + addsql + " ORDER BY id DESC"
			log := []sql.LoginLog{}
			total := []sql.LoginLog{}
			err := c.Sql().Select(&log, logSql, first, row)
			if err != nil {
				c.Log().Debug().Err(err).Msg("Error")
				c.ResponseJson(4, err.Error(), w, r)
			}
			err = c.Sql().Select(&total, totalSql)
			if err != nil {
				c.Log().Debug().Err(err).Msg("Error")
				c.ResponseJson(4, err.Error(), w, r)
			}
			count := len(total)
			if count != 0 {
				jsonLog = "{\"total\":" + strconv.Itoa(count) + ",\"rows\":" + c.RowsJson(log) + "}"
			} else {
				jsonLog = "{\"total\":0,\"rows\":[]}"
			}
			// c.Cache().CacheSetAlwaysTime("allLoginLog", jsonLog)
		}
	}
	fmt.Fprint(w, jsonLog)
}

/**
 * @description 删除一个月前的登录日志
 * @English	Delete the logon log a month ago
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) DelLoginLog(w http.ResponseWriter, r *http.Request) {
	addsql := fmt.Sprintf("login_time<='%v'", time.Now().Unix()-86400*30)
	delSql := "DELETE FROM hm_login_log WHERE " + addsql
	tx := c.Sql().MustBegin()
	tx.MustExec(delSql)
	err := tx.Commit()
	if err != nil {
		c.Log().Debug().Err(err).Msg("Error")
		c.ResponseJson(4, err.Error(), w, r)
	} else {
		c.Cache().CacheDel("allmenu")
		c.ResponseJson(1, "", w, r)
	}
}

/**
 * @description 操作日志页面
 * @English	Oprate log page
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) OprateLog(w http.ResponseWriter, r *http.Request) {
	c.Template(w, r, nil, "./view/admin/site/oprateLog.html")
}

/**
 * @description 获取操作日志
 * @English	Get oprate log
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) GetOprateLog(w http.ResponseWriter, r *http.Request) {
	jsonLog := ""
	page, row := r.PostFormValue("page"), r.PostFormValue("rows")
	username, dateFrom, dateTo := r.PostFormValue("username"), r.PostFormValue("dateFrom"), r.PostFormValue("dateTo")
	isUsername, isDateFrom, isDateTo := c.Regexp().Username(username), c.Regexp().Date(dateFrom), c.Regexp().Date(dateTo)
	switch {
	case isUsername == false && username != "":
		jsonLog = "{\"total\":0,\"rows\":[]}"
	case isDateFrom == false && dateFrom != "":
		jsonLog = "{\"total\":0,\"rows\":[]}"
		fmt.Println("datefrom")
	case isDateTo == false && dateTo != "":
		jsonLog = "{\"total\":0,\"rows\":[]}"
		fmt.Println("dateto")
	default:
		ipage, _ := strconv.Atoi(page)
		irow, _ := strconv.Atoi(row)
		first := irow * (ipage - 1)
		timeLayout := "2006-01-02 15:04:05"  //转化所需模板
		loc, _ := time.LoadLocation("Local") //重要：获取时区
		from, _ := time.ParseInLocation(timeLayout, dateFrom, loc)
		to, _ := time.ParseInLocation(timeLayout, dateTo, loc)
		fromTime, toTime := from.Unix(), to.Unix()
		addsql := ""
		if username != "" {
			addsql = "username LIKE '%" + username + "%' AND "
		}
		if dateFrom != "" {
			addsql = addsql + fmt.Sprintf("oprate_time>='%v' AND ", fromTime)
		}
		if dateTo != "" {
			addsql = addsql + fmt.Sprintf("oprate_time<='%v' AND ", toTime)
		}
		if addsql != "" {
			addsql = "WHERE " + strings.Trim(addsql, "AND ")
		}
		if rows, found := c.Cache().CacheGet("allLoginLog"); found {
			// array = rows.([]sql.User)
			jsonLog = rows.(string)
		} else {
			logSql := "SELECT * FROM hm_oprate_log " + addsql + " ORDER BY id DESC Limit ?,?"
			totalSql := "SELECT id FROM hm_oprate_log " + addsql + " ORDER BY id DESC"
			log := []sql.OprateLog{}
			total := []sql.OprateLog{}
			err := c.Sql().Select(&log, logSql, first, row)
			if err != nil {
				c.Log().Debug().Err(err).Msg("Error")
				c.ResponseJson(4, err.Error(), w, r)
			}
			err = c.Sql().Select(&total, totalSql)
			if err != nil {
				c.Log().Debug().Err(err).Msg("Error")
				c.ResponseJson(4, err.Error(), w, r)
			}
			count := len(total)
			if count != 0 {
				jsonLog = "{\"total\":" + strconv.Itoa(count) + ",\"rows\":" + c.RowsJson(log) + "}"
			} else {
				jsonLog = "{\"total\":0,\"rows\":[]}"
			}
			// c.Cache().CacheSetAlwaysTime("allLoginLog", jsonLog)
		}
	}
	fmt.Fprint(w, jsonLog)
}
