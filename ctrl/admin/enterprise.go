package admin

import (
	"fmt"
	"hemacms/common"
	"hemacms/common/sql"
	"net/http"

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
	an := sql.RecursiveNav(nav, 0, 0)
	fmt.Fprint(w, c.RowsJson(an))
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
	ar := sql.RecursiveNav(nav, 0, 0)
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
			sqls := "INSERT INTO hm_enterprise_nav(url,name,pid,isshow,sort,icon,level,en) VALUES(?,?,?,?,?,?,?,?)"
			tx := c.Sql().MustBegin()
			tx.MustExec(sqls, url, name, pid, isshow, sort, icon, 1, en)
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
