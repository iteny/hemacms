package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hemacms/common"
	"hemacms/common/sql"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	vali "github.com/iteny/hmgo/validator"
)

type EnterpriseCtrl struct {
	common.BaseCtrl
}

func EnterpriseCtrlObject() *EnterpriseCtrl {
	return &EnterpriseCtrl{}
}

/**
 * @description 导航列表
 * @English	nav management page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) NavList(w http.ResponseWriter, r *http.Request) {
	c.Template(w, r, nil, "./view/admin/enterprise/navList.html")
}

/**
 * @description 获取导航
 * @English	get nav
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) GetNav(w http.ResponseWriter, r *http.Request) {
	nav := []sql.EnterpriseNav{}
	if rows, found := c.Cache().Get("allnav"); found {
		nav = rows.([]sql.EnterpriseNav)
	} else {
		sqls := "SELECT * FROM hm_enterprise_nav ORDER BY sort ASC"
		err := c.Sql().Select(&nav, sqls)
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		}
		c.Cache().SetAlwaysTime("allnav", nav)
	}
	fmt.Println(nav)
	an := sql.RecursiveNavLevel(nav, 0, 0)
	fmt.Fprint(w, c.RowsJson(an))
}

/**
 * @description 导航批量排序
 * @English	batch sorting navs
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) SortNav(w http.ResponseWriter, r *http.Request) {
	sortNav := make(map[string]string, 0)
	defer r.Body.Close()
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.ResponseJson(4, ""+err.Error(), w, r)
	}
	json.Unmarshal(result, &sortNav)
	var menuk []string
	for k := range sortNav {
		menuk = append(menuk, k)
	}
	var bf bytes.Buffer
	bf.WriteString("UPDATE hm_enterprise_nav SET sort = CASE id ")
	// nav := make(map[string]string, 0)
	for _, v := range menuk {
		mid := vali.NumericNoHeadZero(v) && vali.Length(v, 1, 8)
		msort := vali.NumericNoHeadZero(sortNav[v]) && vali.Length(v, 1, 3)
		switch false {
		case mid:
			c.ResponseJson(11, "Invalid nav id", w, r)
			return
		case msort:
			c.ResponseJson(12, "Invalid sort", w, r)
			return
		default:
			bf.WriteString(fmt.Sprintf("WHEN %v THEN %v ", v, sortNav[v]))
		}
	}
	// fmt.Println(nav)
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
		c.Cache().Del("allnav")
		c.ResponseJson(1, "", w, r)
	}
}

/**
 * @description 添加导航页面
 * @English	add nav page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) AddNav(w http.ResponseWriter, r *http.Request) {
	nav := []sql.EnterpriseNav{}
	if rows, found := c.Cache().Get("allnav"); found {
		nav = rows.([]sql.EnterpriseNav)
	} else {
		sqls := "SELECT * FROM hm_enterprise_nav"
		err := c.Sql().Select(&nav, sqls)
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		}
		c.Cache().SetAlwaysTime("allnav", nav)
	}
	data := make(map[string]interface{})
	ar := sql.RecursiveNavLevel(nav, 0, 0)
	data["json"] = c.RowsJson(ar)
	pid := chi.URLParam(r, "navPid")
	data["pid"] = pid
	c.Template(w, r, data, "./view/admin/enterprise/addNav.html")
}

/**
 * @description 添加导航提交
 * @English	add nav submit
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) AddNavSubmit(w http.ResponseWriter, r *http.Request) {
	pid, name, url, en := r.PostFormValue("pid"), r.PostFormValue("name"), r.PostFormValue("url"), r.PostFormValue("en")
	sort, icon, isshow := r.PostFormValue("sort"), r.PostFormValue("icon"), r.PostFormValue("isshow")
	switch false {
	case vali.Numeric(pid) && vali.Length(pid, 1, 8):
		c.ResponseJson(11, "Invalid nav pid", w, r)
		return
	case vali.Chinese(name) && vali.Length(name, 2, 50):
		c.ResponseJson(12, "Invalid chinese nav name", w, r)
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
		c.ResponseJson(16, "Invalid nav url", w, r)
		return
	case vali.EnglishSpace(en) && vali.Length(en, 2, 50):
		c.ResponseJson(17, "Invalid english nav name", w, r)
		return
	default:
		nav := sql.EnterpriseNav{}
		err := c.Sql().Get(&nav, "SELECT * FROM hm_enterprise_nav WHERE name = ? OR url = ? OR en = ?", name, url, en)
		if err != nil {
			sqls := "INSERT INTO hm_enterprise_nav(url,name,external_link,dir,type,template,seo_title,seo_keyword,seo_describe,pid,isshow,sort,icon,level,en) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
			tx := c.Sql().MustBegin()
			tx.MustExec(sqls, url, name, "", "", "", "", "", "", "", pid, isshow, sort, icon, 1, en)
			err = tx.Commit()
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			} else {
				c.Cache().Del("allnav")
				c.ResponseJson(1, "", w, r)
			}
		} else {
			if nav.Name == name {
				c.ResponseJson(21, "nav chinese name already exist", w, r)
				return
			} else if nav.Url == url {
				c.ResponseJson(22, "nav url already exist", w, r)
				return
			} else if nav.En == en {
				c.ResponseJson(23, "nav english name already exist", w, r)
				return
			} else {
				c.ResponseJson(44, "Unknown error", w, r)
				return
			}
		}
	}
}

/**
 * @description 修改导航页面
 * @English	edit nav page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) EditNav(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	id := chi.URLParam(r, "navId")
	navSingle := sql.EnterpriseNav{}
	err := c.Sql().Get(&navSingle, "SELECT * FROM hm_enterprise_nav WHERE id = ?", id)
	if err != nil {
		c.ResponseJson(4, err.Error(), w, r)
	}
	data["nav"] = navSingle
	allnav := []sql.EnterpriseNav{}
	if rows, found := c.Cache().Get("allnav"); found {
		allnav = rows.([]sql.EnterpriseNav)
	} else {
		sqls := "SELECT * FROM hm_enterprise_nav"
		err := c.Sql().Select(&allnav, sqls)
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		}
		c.Cache().SetAlwaysTime("allnav", allnav)
	}
	ar := sql.RecursiveNavLevel(allnav, 0, 0)
	data["json"] = c.RowsJson(ar)
	c.Template(w, r, data, "./view/admin/enterprise/editNav.html")
}

/**
 * @description 修改导航提交
 * @English	edit nav submit
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) EditNavSubmit(w http.ResponseWriter, r *http.Request) {
	pid, name, url, en := r.PostFormValue("pid"), r.PostFormValue("name"), r.PostFormValue("url"), r.PostFormValue("en")
	sort, icon, isshow, id := r.PostFormValue("sort"), r.PostFormValue("icon"), r.PostFormValue("isshow"), r.PostFormValue("id")
	switch false {
	case vali.Numeric(pid) && vali.Length(pid, 1, 8):
		c.ResponseJson(11, "Invalid nav pid", w, r)
		return
	case vali.Chinese(name) && vali.Length(name, 2, 50):
		c.ResponseJson(12, "Invalid chinese nav name", w, r)
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
		c.ResponseJson(16, "Invalid nav url", w, r)
		return
	case vali.EnglishSpace(en) && vali.Length(en, 2, 50):
		c.ResponseJson(17, "Invalid english nav name", w, r)
		return
	case vali.NumericNoHeadZero(id) && vali.Length(id, 1, 8):
		c.ResponseJson(18, "Invalid nav id", w, r)
		return
	default:
		nav := sql.EnterpriseNav{}
		err := c.Sql().Get(&nav, "SELECT * FROM hm_enterprise_nav WHERE id <> ? AND (name = ? OR url = ? OR en = ?)", id, name, url, en)
		if err != nil {
			sqls := "UPDATE hm_enterprise_nav SET url=?,name=?,pid=?,isshow=?,sort=?,icon=?,level=?,en=? WHERE id = ?"
			tx := c.Sql().MustBegin()
			tx.MustExec(sqls, url, name, pid, isshow, sort, icon, 1, en, id)
			err := tx.Commit()
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			} else {
				c.Cache().Del("allnav")
				c.ResponseJson(1, "", w, r)
			}
		} else {
			if nav.Name == name {
				c.ResponseJson(21, "Nav chinese name already exist", w, r)
				return
			} else if nav.Url == url {
				c.ResponseJson(22, "Nav url already exist", w, r)
				return
			} else if nav.En == en {
				c.ResponseJson(23, "Nav english name already exist", w, r)
				return
			} else {
				c.ResponseJson(44, "Unknown error", w, r)
				return
			}
		}
	}
}

/**
 * @description 删除导航提交
 * @English	delete nav submit
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) DelNavSubmit(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	if vali.NumericNoHeadZero(id) && vali.Length(id, 1, 8) {
		nav := []sql.EnterpriseNav{}
		if rows, found := c.Cache().Get("allnav"); found {
			nav = rows.([]sql.EnterpriseNav)
		} else {
			sqls := "SELECT * FROM hm_enterprise_nav"
			err := c.Sql().Select(&nav, sqls)
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			}
			c.Cache().SetAlwaysTime("allnav", nav)
		}
		ids := sql.RecursiveNavId(nav, id)
		delSql := fmt.Sprintf("DELETE FROM hm_enterprise_nav WHERE id IN(%v)", ids)
		tx := c.Sql().MustBegin()
		tx.MustExec(delSql)
		err := tx.Commit()
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		} else {
			c.Cache().Del("allnav")
			c.ResponseJson(1, "", w, r)
		}
	} else {
		c.ResponseJson(11, "Invalid menu id", w, r)
		return
	}
}

/**
 * @description 内容管理
 * @English	content management page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) ContentManage(w http.ResponseWriter, r *http.Request) {
	nav := []sql.EnterpriseNav{}
	if rows, found := c.Cache().Get("allnav"); found {
		nav = rows.([]sql.EnterpriseNav)
	} else {
		sqls := "SELECT * FROM hm_enterprise_nav"
		err := c.Sql().Select(&nav, sqls)
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		}
		c.Cache().SetAlwaysTime("allnav", nav)
	}
	data := make(map[string]interface{})
	ar := sql.RecursiveNavLevel(nav, 0, 0)
	data["json"] = c.RowsJson(ar)
	c.Template(w, r, data, "./view/admin/enterprise/contentManage.html")
}
