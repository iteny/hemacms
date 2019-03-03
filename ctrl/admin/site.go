package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hemacms/common"
	"hemacms/common/sql"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
	vali "github.com/iteny/hmgo/validator"
)

const (
	upload_path string = "./static/upload/"
)

type SiteCtrl struct {
	common.BaseCtrl
}

func SiteCtrlObject() *SiteCtrl {
	return &SiteCtrl{}
}

/**
 * @description 上传图片管理
 * @English	uploader images management page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) UploadImage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	//接收客户端传来的文件 uploadfile 与客户端保持一致
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//上传的文件保存在ppp路径下
	ext := path.Ext(handler.Filename) //获取文件后缀
	fileNewName := string(time.Now().Format("20060102150405")) + strconv.Itoa(time.Now().Nanosecond()) + ext

	f, err := os.OpenFile(upload_path+fileNewName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		c.ResponseJson(4, err.Error(), w, r)
		return
	} else {
		session := c.Sess().Load(r)
		uid, err := session.GetString("uid")
		c.Log().CheckErr("Session Get Error", err)
		pic := sql.Picture{}
		err = c.Sql().Get(&pic, "SELECT * FROM hm_picture WHERE url = ?", fileNewName)
		if err != nil {
			sqls := "INSERT INTO hm_picture(url,time,uid) VALUES(?,?,?)"
			tx := c.Sql().MustBegin()
			tx.MustExec(sqls, fileNewName, time.Now().Unix(), uid)
			err = tx.Commit()
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			} else {
				c.Cache().Del("allpic")
				c.ResponseJson(1, fileNewName, w, r)
			}
		} else {
			c.ResponseJson(21, "picture name already exist", w, r)
			return
		}
	}
	defer f.Close()
	io.Copy(f, file)
	// fmt.Fprintln(w, "upload ok!"+fileNewName)
}

/**
 * @description 删除图片
 * @English	delete images file
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) DelImage(w http.ResponseWriter, r *http.Request) {
	pic := r.PostFormValue("pic")
	picSql := "DELETE FROM hm_picture WHERE url = ?"
	tx := c.Sql().MustBegin()
	tx.MustExec(picSql, pic)
	err := tx.Commit()
	if err != nil {
		c.ResponseJson(4, err.Error(), w, r)
	} else {
		c.Cache().ScanDel("allpic")
		c.ResponseJson(1, "", w, r)
	}
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
	if rows, found := c.Cache().Get("allmenu"); found {
		rule = rows.([]sql.AuthRule)
	} else {
		sqls := "SELECT * FROM hm_auth_rule ORDER BY sort ASC"
		err := c.Sql().Select(&rule, sqls)
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		}
		c.Cache().SetAlwaysTime("allmenu", rule)
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
	}
	json.Unmarshal(result, &sortMenu)
	var menuk []string
	for k := range sortMenu {
		menuk = append(menuk, k)
	}
	var bf bytes.Buffer
	bf.WriteString("UPDATE hm_auth_rule SET sort = CASE id ")
	// menu := make(map[string]string, 0)
	for _, v := range menuk {
		mid := vali.NumericNoHeadZero(v) && vali.Length(v, 1, 8)
		msort := vali.NumericNoHeadZero(sortMenu[v]) && vali.Length(v, 1, 3)
		switch false {
		case mid:
			c.ResponseJson(11, "Invalid menu id", w, r)
			return
		case msort:
			c.ResponseJson(12, "Invalid sort", w, r)
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
		c.Cache().Del("allmenu")
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
	for _, v := range strings.Split(s, "\n") {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		ss = append(ss, v)
	}
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
	if rows, found := c.Cache().Get("allmenu"); found {
		rule = rows.([]sql.AuthRule)
	} else {
		sqls := "SELECT * FROM hm_auth_rule"
		err := c.Sql().Select(&rule, sqls)
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		}
		c.Cache().SetAlwaysTime("allmenu", rule)
	}
	data := make(map[string]interface{})
	ar := sql.RecursiveMenuLevel(rule, 0, 0)
	data["json"] = c.RowsJson(ar)
	pid := chi.URLParam(r, "menuPid")
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
	switch false {
	case vali.Numeric(pid) && vali.Length(pid, 1, 8):
		c.ResponseJson(11, "Invalid menu pid", w, r)
		return
	case vali.Chinese(name) && vali.Length(name, 2, 50):
		c.ResponseJson(12, "Invalid chinese menu name", w, r)
		return
	case vali.NumericNoHeadZero(sort) && vali.Length(sort, 1, 3):
		c.ResponseJson(13, "Invalid sort", w, r)
		return
	case vali.IconCss(icon) && vali.Length(icon, 2, 50):
		c.ResponseJson(14, "Invalid icon css", w, r)
		return
	case vali.Status(isshow):
		c.ResponseJson(15, "Invalid status", w, r)
		return
	case vali.MenuUrl(url) && vali.Length(url, 2, 50):
		c.ResponseJson(16, "Invalid menu url", w, r)
		return
	case vali.EnglishSpace(en) && vali.Length(en, 2, 50):
		c.ResponseJson(17, "Invalid english menu name", w, r)
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
				c.ResponseJson(4, err.Error(), w, r)
			} else {
				c.Cache().Del("allmenu")
				c.ResponseJson(1, "", w, r)
			}
		} else {
			if rule.Name == name {
				c.ResponseJson(21, "Menu chinese name already exist", w, r)
				return
			} else if rule.Url == url {
				c.ResponseJson(22, "Menu url already exist", w, r)
				return
			} else if rule.En == en {
				c.ResponseJson(23, "Menu english name already exist", w, r)
				return
			} else {
				c.ResponseJson(44, "Unknown error", w, r)
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
		c.ResponseJson(4, err.Error(), w, r)
	}
	data["menu"] = ruleSingle
	allrule := []sql.AuthRule{}
	if rows, found := c.Cache().Get("allmenu"); found {
		allrule = rows.([]sql.AuthRule)
	} else {
		sqls := "SELECT * FROM hm_auth_rule"
		err := c.Sql().Select(&allrule, sqls)
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		}
		c.Cache().SetAlwaysTime("allmenu", allrule)
	}
	ar := sql.RecursiveMenuLevel(allrule, 0, 0)
	data["json"] = c.RowsJson(ar)
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
	switch false {
	case vali.NumericNoHeadZero(pid) && vali.Length(pid, 1, 8):
		c.ResponseJson(11, "Invalid menu pid", w, r)
		return
	case vali.Chinese(name) && vali.Length(name, 2, 50):
		c.ResponseJson(12, "Invalid chinese menu name", w, r)
		return
	case vali.NumericNoHeadZero(sort) && vali.Length(sort, 1, 3):
		c.ResponseJson(13, "Invalid sort", w, r)
		return
	case vali.IconCss(icon) && vali.Length(icon, 2, 50):
		c.ResponseJson(14, "Invalid icon css", w, r)
		return
	case vali.Status(isshow):
		c.ResponseJson(15, "Invalid status", w, r)
		return
	case vali.MenuUrl(url) && vali.Length(url, 2, 50):
		c.ResponseJson(16, "Invalid menu url", w, r)
		return
	case vali.EnglishSpace(en) && vali.Length(en, 2, 50):
		c.ResponseJson(17, "Invalid english menu name", w, r)
		return
	case vali.NumericNoHeadZero(id) && vali.Length(id, 1, 8):
		c.ResponseJson(18, "Invalid menu id", w, r)
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
				c.ResponseJson(4, err.Error(), w, r)
			} else {
				c.Cache().Del("allmenu")
				c.ResponseJson(1, "", w, r)
			}
		} else {
			if rule.Name == name {
				c.ResponseJson(21, "Menu chinese name already exist", w, r)
				return
			} else if rule.Url == url {
				c.ResponseJson(22, "Menu url already exist", w, r)
				return
			} else if rule.En == en {
				c.ResponseJson(23, "Menu english name already exist", w, r)
				return
			} else {
				c.ResponseJson(44, "Unknown error", w, r)
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
	if vali.NumericNoHeadZero(id) && vali.Length(id, 1, 8) {
		rule := []sql.AuthRule{}
		if rows, found := c.Cache().Get("allmenu"); found {
			rule = rows.([]sql.AuthRule)
		} else {
			sqls := "SELECT * FROM hm_auth_rule"
			err := c.Sql().Select(&rule, sqls)
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			}
			c.Cache().SetAlwaysTime("allmenu", rule)
		}
		ids := sql.RecursiveMenuId(rule, id)
		delSql := fmt.Sprintf("DELETE FROM hm_auth_rule WHERE id IN(%v)", ids)
		tx := c.Sql().MustBegin()
		tx.MustExec(delSql)
		err := tx.Commit()
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		} else {
			c.Cache().Del("allmenu")
			c.ResponseJson(1, "", w, r)
		}
	} else {
		c.ResponseJson(11, "Invalid menu id", w, r)
		return
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
	data["port"] = c.Config().Int("servSet", "port")
	data["readTimeout"] = c.Config().Int("servSet", "readTimeout")
	data["writeTimeout"] = c.Config().Int("servSet", "writeTimeout")
	data["language"] = c.Config().Value("cookie", "language")
	data["terminalLog"] = c.Config().Value("common", "terminalLog")
	data["sqlLog"] = c.Config().Value("common", "sqlLog")
	data["ajaxPolling"] = c.Config().Value("common", "ajaxPolling")
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
	port, language := r.PostFormValue("port"), r.PostFormValue("language")
	readTimeout, writeTimeout := r.PostFormValue("readTimeout"), r.PostFormValue("writeTimeout")
	terminalLog, sqlLog := r.PostFormValue("terminalLog"), r.PostFormValue("sqlLog")
	ajaxPolling := r.PostFormValue("ajaxPolling")
	switch false {
	case vali.NumericNoHeadZero(port) && vali.Length(port, 1, 4):
		c.ResponseJson(12, "Invalid http port", w, r)
		return
	case vali.NumericNoHeadZero(readTimeout) && vali.Length(readTimeout, 1, 5):
		c.ResponseJson(13, "Invalid read timeout", w, r)
		return
	case vali.NumericNoHeadZero(writeTimeout) && vali.Length(writeTimeout, 1, 5):
		c.ResponseJson(14, "Invalid write timeout", w, r)
		return
	case vali.MultiLanguage(language):
		c.ResponseJson(15, "Invalid default language", w, r)
		return
	case vali.Reg(terminalLog, `on|off`):
		c.ResponseJson(16, "Invalid terminal log", w, r)
		return
	case vali.Reg(sqlLog, `on|off`):
		c.ResponseJson(17, "Invalid sql log", w, r)
		return
	case vali.NumericNoHeadZero(ajaxPolling) && vali.Length(ajaxPolling, 1, 3):
		c.ResponseJson(18, "", w, r)
		return
	default:
		c.Config().SetValue("servSet", "port", port)
		c.Config().SetValue("servSet", "readTimeout", readTimeout)
		c.Config().SetValue("servSet", "writeTimeout", writeTimeout)
		c.Config().SetValue("cookie", "language", language)
		c.Config().SetValue("common", "terminalLog", terminalLog)
		c.Config().SetValue("common", "sqlLog", sqlLog)
		c.Config().SetValue("common", "ajaxPolling", ajaxPolling)
		err := c.Config().Save("./ini/hemacms.ini")
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		} else {
			c.SetCookie("back-language", language, 0, w)
			err = c.Config().Reload()
			if err != nil {
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
	data := make(map[string]interface{}, 4)
	page, row := r.PostFormValue("page"), r.PostFormValue("rows")
	username, dateFrom, dateTo := r.PostFormValue("username"), r.PostFormValue("dateFrom"), r.PostFormValue("dateTo")
	data["total"], data["rows"] = 0, []string{}
	switch {
	case ((vali.English(username) && vali.Length(username, 5, 20)) == false) && username != "":
		data["status"], data["info"] = 11, "Invalid account"
	case vali.Time(dateFrom, vali.RF3339Js) == false && dateFrom != "":
		data["status"], data["info"] = 12, "Invalid dateFrom"
	case vali.Time(dateTo, vali.RF3339Js) == false && dateTo != "":
		data["status"], data["info"] = 13, "Invalid dateTo"
	case (vali.NumericNoHeadZero(page) && vali.Length(page, 1, 5) && vali.NumericNoHeadZero(row) && vali.Length(row, 2, 2)) == false:
		data["status"], data["info"] = 14, "Invalid pagination"
	default:
		pageNum := c.Pagination(page, row)
		fromTime, toTime := c.DateToTimestamp(dateFrom, vali.RF3339Js), c.DateToTimestamp(dateTo, vali.RF3339Js)
		addsql := ""
		if username != "" {
			addsql = addsql + "username LIKE '%" + username + "%' AND "
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
		count := 0
		if rows, found := c.Cache().Get("userCount" + username + dateFrom + dateTo + page + row); found {
			count = rows.(int)
		} else {
			totalSql := "SELECT id FROM hm_user " + addsql
			total := []sql.User{}
			err := c.Sql().Select(&total, totalSql)
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			}
			count = len(total)
			c.Cache().SetAlwaysTime("userCount"+username+dateFrom+dateTo+page+row, count)
		}
		// log := []sql.LoginLog{}
		if rows, found := c.Cache().Get("user" + username + dateFrom + dateTo + page + row); found {
			array = rows.([]sql.User)
		} else {
			userSql := "SELECT * FROM hm_user " + addsql + " ORDER BY id ASC Limit ?,?"
			roleSql := "SELECT * FROM hm_auth_role"
			accessSql := "SELECT uid,role_id FROM hm_auth_role_access"
			user := []sql.User{}
			role := []sql.AuthRole{}
			access := []sql.AuthRoleAccess{}
			err := c.Sql().Select(&user, userSql, pageNum, row)
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			}
			err = c.Sql().Select(&role, roleSql)
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			}
			err = c.Sql().Select(&access, accessSql)
			if err != nil {
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
			c.Cache().SetAlwaysTime("user"+username+dateFrom+dateTo+page+row, array)
		}
		if count != 0 {
			data["total"], data["rows"] = strconv.Itoa(count), array
		} else {
			data["status"], data["info"], data["rows"] = 54, "no found data", array
		}
		c.Cache().ScanKey("user")
	}
	c.ResponseData(data, w, r)
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
	if vali.NumericNoHeadZero(id) && vali.Length(id, 1, 8) {
		if id == "1" {
			c.ResponseJson(12, "Super admin don't allow deletion", w, r)
			return
		} else {
			userSql := "DELETE FROM hm_user WHERE id = ?"
			accessSql := "DELETE FROM hm_auth_role_access WHERE uid = ?"
			tx := c.Sql().MustBegin()
			tx.MustExec(userSql, id)
			tx.MustExec(accessSql, id)
			err := tx.Commit()
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			} else {
				c.Cache().ScanDel("user")
				c.ResponseJson(1, "", w, r)
			}
		}
	} else {
		c.ResponseJson(11, "Invalid user id", w, r)
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
	idSplit := strings.Split(ids, ",")
	for _, v := range idSplit {
		if v == "1" {
			c.ResponseJson(12, "Super admin don't allow deletion", w, r)
			return
		}
		if (vali.NumericNoHeadZero(v) && vali.Length(v, 1, 8)) == false {
			c.ResponseJson(11, "Invalid user id", w, r)
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
		c.ResponseJson(4, err.Error(), w, r)
	} else {
		c.Cache().ScanDel("user")
		c.ResponseJson(1, "", w, r)
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
	data := make(map[string]interface{}, 2)
	roleSql := "select * from hm_auth_role"
	role := []sql.AuthRole{}
	err := c.Sql().Select(&role, roleSql)
	if err != nil {
		c.ResponseJson(4, err.Error(), w, r)
	}
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
	switch false {
	case vali.EnglishNumeric(username) && vali.Length(username, 5, 20):
		c.ResponseJson(11, "Invalid account name", w, r)
		return
	case vali.EnglishNumeric(password) && vali.Length(password, 5, 15):
		c.ResponseJson(12, "Invalid password", w, r)
		return
	case ((vali.Chinese(nickname) || vali.EnglishSpace(nickname)) && vali.Length(nickname, 2, 50)) == true:
		c.ResponseJson(13, "Invalid nickname", w, r)
		return
	case vali.Email(email) && vali.Length(email, 4, 50):
		c.ResponseJson(14, "Invalid email", w, r)
		return
	case vali.Article(remark) && vali.Length(remark, 2, 255):
		c.ResponseJson(15, "Invalid remark", w, r)
		return
	case vali.NumericNoHeadZero(role) && vali.Length(role, 1, 8):
		c.ResponseJson(16, "Invalid role id", w, r)
		return
	case vali.Status(status):
		c.ResponseJson(17, "Invalid status", w, r)
		return
	case vali.EqualTo(password, passworded):
		c.ResponseJson(18, "Two input password mismatch", w, r)
		return
	default:
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
				c.ResponseJson(4, err.Error(), w, r)
			}
			tx.MustExec(accessSql, id, role)
			err = tx.Commit()
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			} else {
				c.Cache().ScanDel("user")
				c.ResponseJson(1, "", w, r)
			}
		} else {
			if user.Username == username {
				c.ResponseJson(21, "Account name already exist", w, r)
				return
			} else if user.Nickname == nickname {
				c.ResponseJson(22, "Nickname already exist", w, r)
				return
			} else if user.Email == email {
				c.ResponseJson(23, "Email already exist", w, r)
				return
			} else {
				c.ResponseJson(44, "Unknown error", w, r)
				return
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
		c.ResponseJson(4, err.Error(), w, r)
	}
	err = c.Sql().Get(&accessSingle, "SELECT * FROM hm_auth_role_access WHERE uid = ?", id)
	if err != nil {
		c.ResponseJson(4, err.Error(), w, r)
	}
	data["user"] = userSingle
	data["roleme"] = accessSingle
	roleSql := "select * from hm_auth_role"
	role := []sql.AuthRole{}
	err = c.Sql().Select(&role, roleSql)
	if err != nil {
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
	switch false {
	case vali.EnglishNumeric(username) && vali.Length(username, 5, 20):
		c.ResponseJson(11, "Invalid account name", w, r)
		return
	case vali.EnglishNumeric(password) && vali.Length(password, 5, 15):
		c.ResponseJson(12, "Invalid password", w, r)
		return
	case ((vali.Chinese(nickname) || vali.EnglishSpace(nickname)) && vali.Length(nickname, 2, 50)) == true:
		c.ResponseJson(13, "Invalid nickname", w, r)
		return
	case vali.Email(email) && vali.Length(email, 4, 50):
		c.ResponseJson(14, "Invalid email", w, r)
		return
	case vali.Article(remark) && vali.Length(remark, 2, 255):
		c.ResponseJson(15, "Invalid remark", w, r)
		return
	case vali.NumericNoHeadZero(role) && vali.Length(role, 1, 8):
		c.ResponseJson(16, "Invalid role id", w, r)
		return
	case vali.Status(status):
		c.ResponseJson(17, "Invalid status", w, r)
		return
	case vali.EqualTo(password, passworded):
		c.ResponseJson(18, "Two input password mismatch", w, r)
		return
	case vali.NumericNoHeadZero(id) && vali.Length(id, 1, 8):
		c.ResponseJson(19, "Invalid account id", w, r)
		return
	default:
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
				c.ResponseJson(4, err.Error(), w, r)
			} else {
				c.Cache().ScanDel("user")
				c.ResponseJson(1, "", w, r)
			}
		} else {
			if user.Username == username {
				c.ResponseJson(21, "Account name already exist", w, r)
				return
			} else if user.Nickname == nickname {
				c.ResponseJson(22, "Nickname already exist", w, r)
				return
			} else if user.Email == email {
				c.ResponseJson(23, "Email already exist", w, r)
				return
			} else {
				c.ResponseJson(44, "Unknown error", w, r)
				return
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
	data := make(map[string]interface{}, 4)
	page, row, name := r.PostFormValue("page"), r.PostFormValue("rows"), r.PostFormValue("name")
	data["total"], data["rows"] = 0, []string{}
	switch {
	case ((vali.EnglishSpace(name) || vali.Chinese(name)) && vali.Length(name, 2, 20) == false) && name != "":
		data["status"], data["info"] = 11, "Invalid role name"
	default:
		pageNum := c.Pagination(page, row)
		addSql := ""
		if name != "" {
			cookie, err := r.Cookie("back-language")
			if err != nil {
				c.Log().CheckErr("cookie error", err)
				value := c.Config().Value("cookie", "language")
				if vali.MultiLanguage(value) {
					c.SetCookie("back-language", value, 0, w)
				} else {
					c.SetCookie("back-language", "en", 0, w)
				}
			} else {
				addSql = c.MultiLanguageAddSql(cookie.Value, name, "en")
			}
		}
		if addSql != "" {
			addSql = "WHERE " + strings.TrimRight(addSql, "AND ")
		}
		count := 0
		if rows, found := c.Cache().Get("roleCount" + name + page + row); found {
			count = rows.(int)
		} else {
			totalSql := "SELECT id FROM hm_auth_role " + addSql
			total := []sql.AuthRole{}
			err := c.Sql().Select(&total, totalSql)
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			}
			count = len(total)
			c.Cache().SetAlwaysTime("roleCount"+name+page+row, count)
		}
		role := []sql.AuthRole{}
		if rows, found := c.Cache().Get("role" + name + page + row); found {
			role = rows.([]sql.AuthRole)
		} else {
			roleSql := "SELECT * FROM hm_auth_role " + addSql + " ORDER BY id ASC Limit ?,?"
			err := c.Sql().Select(&role, roleSql, pageNum, row)
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			}
			c.Cache().SetAlwaysTime("role"+name+page+row, role)
		}
		if count != 0 && role != nil {
			data["total"], data["rows"] = strconv.Itoa(count), role
		} else {
			data["status"], data["info"], data["rows"] = 54, "no found data", role
		}
	}
	c.ResponseData(data, w, r)
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
	if rows, found := c.Cache().Get("role" + id); found {
		role = rows.(sql.AuthRole)
	} else {
		sqls := "SELECT * FROM hm_auth_role WHERE id = ?"
		err := c.Sql().Get(&role, sqls, id)
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		}
		c.Cache().SetAlwaysTime("role"+id, role)
	}
	rule := []sql.AuthRule{}
	if rows, found := c.Cache().Get("allmenu"); found {
		rule = rows.([]sql.AuthRule)
	} else {
		sqls := "SELECT * FROM hm_auth_rule ORDER BY sort ASC"
		err := c.Sql().Select(&rule, sqls)
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		}
		c.Cache().SetAlwaysTime("allmenu", rule)
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
	if (vali.NumericNoHeadZero(id) && vali.Length(id, 1, 8)) == false {
		c.ResponseJson(11, "Invalid role id", w, r)
		return
	}
	idSplit := strings.Split(ids, ",")
	for _, v := range idSplit {
		if (vali.NumericNoHeadZero(v) && vali.Length(v, 1, 8)) == false {
			c.ResponseJson(12, "Invalid rule id", w, r)
			return
		}
	}
	sqls := "UPDATE hm_auth_role SET rules=? WHERE id = ?"
	tx := c.Sql().MustBegin()
	tx.MustExec(sqls, ids, id)
	err := tx.Commit()
	if err != nil {
		c.ResponseJson(4, err.Error(), w, r)
	} else {
		c.Cache().Del("role" + id)
		c.Cache().Del("auth" + id)
		c.ResponseJson(1, "", w, r)
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
	if vali.NumericNoHeadZero(id) && vali.Length(id, 1, 8) {
		if id == "1" {
			c.ResponseJson(12, "Super admin don't allow deletion", w, r)
			return
		} else {
			roleSql := "DELETE FROM hm_auth_role WHERE id = ?"
			tx := c.Sql().MustBegin()
			tx.MustExec(roleSql, id)
			err := tx.Commit()
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			} else {
				c.Cache().ScanDel("role")
				c.ResponseJson(1, "", w, r)
			}
		}
	} else {
		c.ResponseJson(11, "Invalid role id", w, r)
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
	idSplit := strings.Split(ids, ",")
	for _, v := range idSplit {
		if v == "1" {
			c.ResponseJson(12, "Super admin don't allow deletion", w, r)
			return
		}
		if (vali.NumericNoHeadZero(v) && vali.Length(v, 1, 8)) == false {
			c.ResponseJson(11, "Invalid role id", w, r)
			return
		}
	}
	roleSql := fmt.Sprintf("DELETE FROM hm_auth_role WHERE id IN (%v)", ids)
	tx := c.Sql().MustBegin()
	tx.MustExec(roleSql)
	err := tx.Commit()
	if err != nil {
		c.ResponseJson(4, err.Error(), w, r)
	} else {
		c.Cache().ScanDel("role")
		c.ResponseJson(1, "", w, r)
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
 * @description 添加����色提交
 * @English	add role submit
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) AddRoleSubmit(w http.ResponseWriter, r *http.Request) {
	name, sort, remark, en := r.PostFormValue("name"), r.PostFormValue("sort"), r.PostFormValue("remark"), r.PostFormValue("en")
	switch false {
	case vali.Chinese(name) && vali.Length(name, 2, 50):
		c.ResponseJson(11, "invalid chinese role name", w, r)
		return
	case vali.NumericNoHeadZero(sort) && vali.Length(sort, 1, 3):
		c.ResponseJson(12, "invalid sort", w, r)
		return
	case vali.Article(remark) && vali.Length(remark, 2, 255):
		c.ResponseJson(13, "invalid remark", w, r)
		return
	case vali.EnglishSpace(en) && vali.Length(en, 2, 100):
		c.ResponseJson(14, "invalid english role name ", w, r)
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
				c.ResponseJson(4, err.Error(), w, r)
			} else {
				c.Cache().Del("allrole")
				c.ResponseJson(1, "", w, r)
			}
		} else {
			if role.Name == name {
				c.ResponseJson(21, "Chinese role name already exist", w, r)
				return
			} else if role.En == en {
				c.ResponseJson(22, "English role name already exist", w, r)
				return
			} else {
				c.ResponseJson(44, "Unknown error", w, r)
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
	role := sql.AuthRole{}
	sqls := "SELECT * FROM hm_auth_role WHERE id = ?"
	err := c.Sql().Get(&role, sqls, id)
	if err != nil {
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
	switch false {
	case vali.Chinese(name) && vali.Length(name, 2, 50):
		c.ResponseJson(11, "invalid chinese role name", w, r)
		return
	case vali.NumericNoHeadZero(sort) && vali.Length(sort, 1, 3):
		c.ResponseJson(12, "invalid sort", w, r)
		return
	case vali.Article(remark) && vali.Length(remark, 2, 255):
		c.ResponseJson(13, "invalid remark", w, r)
		return
	case vali.EnglishSpace(en) && vali.Length(en, 2, 100):
		c.ResponseJson(14, "invalid english role name ", w, r)
		return
	case vali.NumericNoHeadZero(id) && vali.Length(id, 1, 8):
		c.ResponseJson(15, "invalid role id", w, r)
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
				c.ResponseJson(4, err.Error(), w, r)
			} else {
				c.Cache().Del("allrole")
				c.ResponseJson(1, "", w, r)
			}
		} else {
			if role.Name == name {
				c.ResponseJson(21, "Chinese role name already exist", w, r)
				return
			} else if role.En == en {
				c.ResponseJson(22, "English role name already exist", w, r)
				return
			} else {
				c.ResponseJson(44, "Unknown error", w, r)
				return
			}
		}
	}
}

/**
 * @description 配色方案
 * @English	color schemes
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
	data := make(map[string]interface{}, 4)
	page, row := r.PostFormValue("page"), r.PostFormValue("rows")
	username, dateFrom, dateTo := r.PostFormValue("username"), r.PostFormValue("dateFrom"), r.PostFormValue("dateTo")
	data["total"], data["rows"] = 0, []string{}
	switch {
	case ((vali.English(username) && vali.Length(username, 5, 20)) == false) && username != "":
		data["status"], data["info"] = 11, "Invalid account"
	case vali.Time(dateFrom, vali.RF3339Js) == false && dateFrom != "":
		data["status"], data["info"] = 12, "Invalid dateFrom"
	case vali.Time(dateTo, vali.RF3339Js) == false && dateTo != "":
		data["status"], data["info"] = 13, "Invalid dateTo"
	case (vali.NumericNoHeadZero(page) && vali.Length(page, 1, 5) && vali.NumericNoHeadZero(row) && vali.Length(row, 2, 2)) == false:
		data["status"], data["info"] = 14, "Invalid pagination"
	default:
		pageNum := c.Pagination(page, row)
		fromTime, toTime := c.DateToTimestamp(dateFrom, vali.RF3339Js), c.DateToTimestamp(dateTo, vali.RF3339Js)
		addsql := ""
		if username != "" {
			addsql = addsql + "username LIKE '%" + username + "%' AND "
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
		count := 0
		if rows, found := c.Cache().Get("loginLogCount" + username + dateFrom + dateTo + page + row); found {
			count = rows.(int)
		} else {
			totalSql := "SELECT id FROM hm_login_log " + addsql
			total := []sql.OprateLog{}
			err := c.Sql().Select(&total, totalSql)
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			}
			count = len(total)
			c.Cache().SetAlwaysTime("loginLogCount"+username+dateFrom+dateTo+page+row, count)
		}
		log := []sql.LoginLog{}
		if rows, found := c.Cache().Get("loginLog" + username + dateFrom + dateTo + page + row); found {
			log = rows.([]sql.LoginLog)
		} else {
			logSql := "SELECT * FROM hm_login_log " + addsql + " ORDER BY id DESC Limit ?,?"
			err := c.Sql().Select(&log, logSql, pageNum, row)
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			}
			c.Cache().SetAlwaysTime("loginLog"+username+dateFrom+dateTo+page+row, log)
		}
		if count != 0 {
			data["total"], data["rows"] = strconv.Itoa(count), log
		} else {
			data["status"], data["info"], data["rows"] = 54, "no found data", log
		}
	}
	c.ResponseData(data, w, r)
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
		c.ResponseJson(4, err.Error(), w, r)
	} else {
		c.Cache().ScanDel("loginLog")
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
 * @description 获取操作������
 * @English	Get oprate log
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) GetOprateLog(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{}, 4)
	page, row, status := r.PostFormValue("page"), r.PostFormValue("rows"), r.PostFormValue("status")
	username, dateFrom, dateTo := r.PostFormValue("username"), r.PostFormValue("dateFrom"), r.PostFormValue("dateTo")
	excuteTime := r.PostFormValue("excuteTime")
	data["total"], data["rows"] = 0, []string{}
	switch {
	case ((vali.English(username) && vali.Length(username, 5, 20)) == false) && username != "":
		data["status"], data["info"] = 11, "Invalid account"
	case ((vali.Numeric(status) && vali.Length(status, 1, 1)) == false) && status != "":
		data["status"], data["info"] = 12, "Invalid status"
	case vali.Time(dateFrom, vali.RF3339Js) == false && dateFrom != "":
		data["status"], data["info"] = 13, "Invalid dateFrom"
	case vali.Time(dateTo, vali.RF3339Js) == false && dateTo != "":
		data["status"], data["info"] = 14, "Invalid dateTo"
	case ((vali.Numeric(excuteTime) && vali.Length(excuteTime, 1, 8)) == false) && excuteTime != "":
		data["status"], data["info"] = 15, "Invalid excuteTime"
	case (vali.NumericNoHeadZero(page) && vali.Length(page, 1, 5) && vali.NumericNoHeadZero(row) && vali.Length(row, 2, 2)) == false:
		data["status"], data["info"] = 16, "Invalid pagination"
	default:
		pageNum := c.Pagination(page, row)
		fromTime, toTime := c.DateToTimestamp(dateFrom, vali.RF3339Js), c.DateToTimestamp(dateTo, vali.RF3339Js)
		addsql := ""
		if username != "" {
			addsql = addsql + "username LIKE '%" + username + "%' AND "
		}
		if status == "1" {
			addsql = addsql + "status =" + status + " AND "
		}
		if status == "0" {
			addsql = addsql + "status <> 1 AND "
		}
		if excuteTime != "" {
			addsql = addsql + "excute_time >= '" + excuteTime + "' AND "
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
		count := 0
		if rows, found := c.Cache().Get("oprateLogCount" + username + status + dateFrom + dateTo + excuteTime + page + row); found {
			count = rows.(int)
		} else {
			totalSql := "SELECT id FROM hm_oprate_log " + addsql
			total := []sql.OprateLog{}
			err := c.Sql().Select(&total, totalSql)
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			}
			count = len(total)
			c.Cache().SetAlwaysTime("oprateLogCount"+username+status+dateFrom+dateTo+excuteTime+page+row, count)
		}
		log := []sql.OprateLog{}
		if rows, found := c.Cache().Get("oprateLog" + username + status + dateFrom + dateTo + excuteTime + page + row); found {
			log = rows.([]sql.OprateLog)
		} else {
			logSql := "SELECT * FROM hm_oprate_log " + addsql + " ORDER BY id DESC Limit ?,?"
			err := c.Sql().Select(&log, logSql, pageNum, row)
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			}
			c.Cache().SetAlwaysTime("oprateLog"+username+status+dateFrom+dateTo+excuteTime+page+row, log)
		}
		if count != 0 {
			data["total"], data["rows"] = strconv.Itoa(count), log
		} else {
			data["status"], data["info"], data["rows"] = 54, "no found data", log
		}
	}
	c.ResponseData(data, w, r)
}

/**
 * @description 删除一个月前的操作日志
 * @English	Delete the oprate log a month ago
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) DelOprateLog(w http.ResponseWriter, r *http.Request) {
	addsql := fmt.Sprintf("oprate_time<='%v'", time.Now().Unix()-86400*30)
	delSql := "DELETE FROM hm_oprate_log WHERE " + addsql
	tx := c.Sql().MustBegin()
	tx.MustExec(delSql)
	err := tx.Commit()
	if err != nil {
		c.ResponseJson(4, err.Error(), w, r)
	} else {
		c.Cache().ScanDel("oprateLog")
		c.ResponseJson(1, "", w, r)
	}
}

/**
 * @description 删除全部操作日志
 * @English	Delete all oprate log
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) DelAllOprateLog(w http.ResponseWriter, r *http.Request) {
	delSql := "DELETE FROM hm_oprate_log"
	tx := c.Sql().MustBegin()
	tx.MustExec(delSql)
	err := tx.Commit()
	if err != nil {
		c.ResponseJson(4, err.Error(), w, r)
	} else {
		c.Cache().ScanDel("oprateLog")
		c.ResponseJson(1, "", w, r)
	}
}

/**
 * @description 更新日志页面
 * @English	update log page
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) UpdateLog(w http.ResponseWriter, r *http.Request) {
	c.Template(w, r, nil, "./view/admin/site/updateLog.html")
}

/**
 * @description 获取更新日志
 * @English	Get update log
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) GetUpdateLog(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{}, 4)
	page, row := r.PostFormValue("page"), r.PostFormValue("rows")
	dateFrom, dateTo := r.PostFormValue("dateFrom"), r.PostFormValue("dateTo")
	data["total"], data["rows"] = 0, []string{}
	switch {
	case vali.Time(dateFrom, vali.RF3339Js) == false && dateFrom != "":
		data["status"], data["info"] = 12, "Invalid dateFrom"
	case vali.Time(dateTo, vali.RF3339Js) == false && dateTo != "":
		data["status"], data["info"] = 13, "Invalid dateTo"
	case (vali.NumericNoHeadZero(page) && vali.Length(page, 1, 5) && vali.NumericNoHeadZero(row) && vali.Length(row, 2, 2)) == false:
		data["status"], data["info"] = 14, "Invalid pagination"
	default:
		pageNum := c.Pagination(page, row)
		fromTime, toTime := c.DateToTimestamp(dateFrom, vali.RF3339Js), c.DateToTimestamp(dateTo, vali.RF3339Js)
		addsql := ""
		if dateFrom != "" {
			addsql = addsql + fmt.Sprintf("time>='%v' AND ", fromTime)
		}
		if dateTo != "" {
			addsql = addsql + fmt.Sprintf("time<='%v' AND ", toTime)
		}
		if addsql != "" {
			addsql = "WHERE " + strings.Trim(addsql, "AND ")
		}
		count := 0
		if rows, found := c.Cache().Get("updateLogCount" + dateFrom + dateTo + page + row); found {
			count = rows.(int)
		} else {
			totalSql := "SELECT id FROM hm_update_log " + addsql
			total := []sql.UpdateLog{}
			err := c.Sql().Select(&total, totalSql)
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			}
			count = len(total)
			c.Cache().SetAlwaysTime("updateLogCount"+dateFrom+dateTo+page+row, count)
		}
		log := []sql.UpdateLog{}
		if rows, found := c.Cache().Get("updateLog" + dateFrom + dateTo + page + row); found {
			log = rows.([]sql.UpdateLog)
		} else {
			logSql := "SELECT * FROM hm_update_log " + addsql + " ORDER BY time DESC Limit ?,?"
			err := c.Sql().Select(&log, logSql, pageNum, row)
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			}
			c.Cache().SetAlwaysTime("updateLog"+dateFrom+dateTo+page+row, log)
		}
		fmt.Println(log)
		if count != 0 {
			data["total"], data["rows"] = strconv.Itoa(count), log
		} else {
			data["status"], data["info"], data["rows"] = 54, "no found data", log
		}
	}
	c.ResponseData(data, w, r)
}

/**
 * @description 添加更新日志页面
 * @English	Add update log page
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) AddUpdateLog(w http.ResponseWriter, r *http.Request) {
	c.Template(w, r, nil, "./view/admin/site/addUpdateLog.html")
}

/**
 * @description 添加更新日志提交
 * @English	Add update log submit
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) AddUpdateLogSubmit(w http.ResponseWriter, r *http.Request) {
	cn, en := r.PostFormValue("cn"), r.PostFormValue("en")
	switch false {
	case vali.Article(cn) && vali.Length(cn, 2, 255):
		c.ResponseJson(11, "invalid chinese content", w, r)
		return
	case vali.Article(en) && vali.Length(en, 2, 255):
		c.ResponseJson(12, "invalid english content", w, r)
		return
	default:
		sqls := "INSERT INTO hm_update_log(cn,en,time) VALUES(?,?,?)"
		tx := c.Sql().MustBegin()
		tx.MustExec(sqls, cn, en, time.Now().Unix())
		err := tx.Commit()
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		} else {
			c.Cache().ScanDel("updateLog")
			c.ResponseJson(1, "", w, r)
		}
	}
}

/**
 * @description 修改更新日志页面
 * @English	Edit update log page
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) EditUpdateLog(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{}, 2)
	id := chi.URLParam(r, "updateLogId")
	updateLogSingle := sql.UpdateLog{}
	err := c.Sql().Get(&updateLogSingle, "SELECT * FROM hm_update_log WHERE id = ?", id)
	if err != nil {
		c.ResponseJson(4, err.Error(), w, r)
	}
	data["updateLog"] = updateLogSingle
	c.Template(w, r, data, "./view/admin/site/editUpdateLog.html")
}

/**
 * @description 修改更新日志提交
 * @English	Edit update log submit
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) EditUpdateLogSubmit(w http.ResponseWriter, r *http.Request) {
	id, cn, en := r.PostFormValue("id"), r.PostFormValue("cn"), r.PostFormValue("en")
	switch false {
	case vali.Article(cn) && vali.Length(cn, 2, 255):
		c.ResponseJson(11, "invalid chinese content", w, r)
		return
	case vali.Article(en) && vali.Length(en, 2, 255):
		c.ResponseJson(12, "invalid english content", w, r)
		return
	case vali.NumericNoHeadZero(id) && vali.Length(id, 1, 8):
		c.ResponseJson(13, "invalid update log id", w, r)
		return
	default:
		sqls := "UPDATE hm_update_log SET cn=?,en=?,time=? WHERE id = ?"
		tx := c.Sql().MustBegin()
		tx.MustExec(sqls, cn, en, time.Now().Unix(), id)
		err := tx.Commit()
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		} else {
			c.Cache().ScanDel("updateLog")
			c.ResponseJson(1, "", w, r)
		}
	}
}

/**
 * @description 删除单个更新日志
 * @English	Delete single update log
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) DelUpdateLog(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	if vali.NumericNoHeadZero(id) && vali.Length(id, 1, 8) {
		updateLogSql := "DELETE FROM hm_update_log WHERE id = ?"
		tx := c.Sql().MustBegin()
		tx.MustExec(updateLogSql, id)
		err := tx.Commit()
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		} else {
			c.Cache().ScanDel("updateLog")
			c.ResponseJson(1, "", w, r)
		}
	} else {
		c.ResponseJson("11", "Invalid update log id", w, r)
	}
}

/**
 * @description 批量删除更新日志
 * @English	Batch delete update logs
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *SiteCtrl) BatchDelUpdateLog(w http.ResponseWriter, r *http.Request) {
	ids := r.PostFormValue("ids")
	idSplit := strings.Split(ids, ",")
	for _, v := range idSplit {
		if (vali.NumericNoHeadZero(v) && vali.Length(v, 1, 8)) == false {
			c.ResponseJson(11, "Invalid update log id", w, r)
			return
		}
	}
	updateLogSql := fmt.Sprintf("DELETE FROM hm_update_log WHERE id IN (%v)", ids)
	tx := c.Sql().MustBegin()
	tx.MustExec(updateLogSql)
	err := tx.Commit()
	if err != nil {
		c.ResponseJson(4, err.Error(), w, r)
	} else {
		c.Cache().ScanDel("updateLog")
		c.ResponseJson(1, "", w, r)
	}
}
