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
 * @description 企业模板类型
 * @English	enterprise template type page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) TemplateType(w http.ResponseWriter, r *http.Request) {
	c.Template(w, r, nil, "./view/admin/enterprise/templateType.html")
}

/**
 * @description 获取模板类型
 * @English	get template type
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) GetTemplateType(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{}, 4)
	name := r.PostFormValue("name")
	switch {
	case ((vali.EnglishSpace(name) || vali.Chinese(name)) && vali.Length(name, 2, 30) == false) && name != "":
		data["status"], data["info"] = 11, "Invalid Template Type Name"
	default:
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
				addSql = c.MultiLanguage(cookie.Value, name, "en")
			}
		}
		if addSql != "" {
			addSql = "WHERE " + strings.Trim(addSql, "AND ")
		}
		fmt.Println(addSql)
		tptypes := []sql.EnterpriseTptype{}
		if rows, found := c.Cache().Get("allTemplateType" + name); found {
			tptypes = rows.([]sql.EnterpriseTptype)
		} else {
			slideSql := "SELECT * FROM hm_enterprise_tptype " + addSql + " ORDER BY sort ASC "
			err := c.Sql().Select(&tptypes, slideSql)
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			}
			c.Cache().SetAlwaysTime("allTemplateType"+name, tptypes)
		}
		data["status"] = 1
		data["rows"] = tptypes
	}
	c.ResponseData(data, w, r)
}

/**
 * @description 添加模板类型页面
 * @English	add template type page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) AddTemplateType(w http.ResponseWriter, r *http.Request) {
	c.Template(w, r, nil, "./view/admin/enterprise/addTemplateType.html")
}

/**
 * @description 添加模板类型提交
 * @English	add template type submit
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) AddTemplateTypeSubmit(w http.ResponseWriter, r *http.Request) {
	cn, en, sort := r.PostFormValue("cn"), r.PostFormValue("en"), r.PostFormValue("sort")
	switch false {
	case vali.Chinese(cn) && vali.Length(cn, 2, 30):
		c.ResponseJson(11, "invalid chinese content", w, r)
		return
	case vali.EnglishSpace(en) && vali.Length(en, 2, 30):
		c.ResponseJson(12, "invalid english content", w, r)
		return
	default:
		sqls := "INSERT INTO hm_enterprise_tptype(cn,en,sort) VALUES(?,?,?)"
		tx := c.Sql().MustBegin()
		tx.MustExec(sqls, cn, en, sort)
		err := tx.Commit()
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		} else {
			c.Cache().ScanDel("allTemplateType")
			c.ResponseJson(1, "", w, r)
		}
	}
}

/**
 * @description 修改模板类型页面
 * @English	edit template type page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) EditTemplateType(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{}, 2)
	id := chi.URLParam(r, "tptypeId")
	tptypeSingle := sql.EnterpriseTptype{}
	err := c.Sql().Get(&tptypeSingle, "SELECT * FROM hm_enterprise_tptype WHERE id = ?", id)
	if err != nil {
		c.ResponseJson(4, err.Error(), w, r)
	}
	data["tptype"] = tptypeSingle
	c.Template(w, r, data, "./view/admin/enterprise/editTemplateType.html")
}

/**
 * @description 修改模板类型提交
 * @English	edit template type submit
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) EditTemplateTypeSubmit(w http.ResponseWriter, r *http.Request) {
	id, cn, en, sort := r.PostFormValue("id"), r.PostFormValue("cn"), r.PostFormValue("en"), r.PostFormValue("sort")
	switch false {
	case vali.Chinese(cn) && vali.Length(cn, 2, 30):
		c.ResponseJson(11, "invalid chinese content", w, r)
		return
	case vali.EnglishSpace(en) && vali.Length(en, 2, 30):
		c.ResponseJson(12, "invalid english content", w, r)
		return
	case vali.NumericNoHeadZero(id) && vali.Length(id, 1, 8):
		c.ResponseJson(13, "invalid update log id", w, r)
		return
	default:
		sqls := "UPDATE hm_enterprise_tptype SET cn=?,en=?,sort=? WHERE id = ?"
		tx := c.Sql().MustBegin()
		tx.MustExec(sqls, cn, en, sort, id)
		err := tx.Commit()
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		} else {
			c.Cache().ScanDel("allTemplateType")
			c.ResponseJson(1, "", w, r)
		}
	}
}

/**
 * @description 删除单个模板类型
 * @English	delete single template type
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) DelTemplateType(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	if vali.NumericNoHeadZero(id) && vali.Length(id, 1, 8) {
		tptypeSql := "DELETE FROM hm_enterprise_tptype WHERE id = ?"
		tx := c.Sql().MustBegin()
		tx.MustExec(tptypeSql, id)
		err := tx.Commit()
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		} else {
			c.Cache().ScanDel("allTemplateType")
			c.ResponseJson(1, "", w, r)
		}
	} else {
		c.ResponseJson("11", "Invalid template type id", w, r)
	}
}

/**
 * @description 批量删除模板类型
 * @English	batch delete template type
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) BatchDelTemplateType(w http.ResponseWriter, r *http.Request) {
	ids := r.PostFormValue("ids")
	idSplit := strings.Split(ids, ",")
	for _, v := range idSplit {
		if (vali.NumericNoHeadZero(v) && vali.Length(v, 1, 8)) == false {
			c.ResponseJson(11, "Invalid template type id", w, r)
			return
		}
	}
	tptypeSql := fmt.Sprintf("DELETE FROM hm_enterprise_tptype WHERE id IN (%v)", ids)
	tx := c.Sql().MustBegin()
	tx.MustExec(tptypeSql)
	err := tx.Commit()
	if err != nil {
		c.ResponseJson(4, err.Error(), w, r)
	} else {
		c.Cache().ScanDel("allTemplateType")
		c.ResponseJson(1, "", w, r)
	}
}

/**
 * @description 模板类型批量排序
 * @English	batch sorting template type
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) SortTemplateType(w http.ResponseWriter, r *http.Request) {
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
	bf.WriteString("UPDATE hm_enterprise_tptype SET sort = CASE id ")
	// nav := make(map[string]string, 0)
	for _, v := range menuk {
		mid := vali.NumericNoHeadZero(v) && vali.Length(v, 1, 8)
		msort := vali.NumericNoHeadZero(sortNav[v]) && vali.Length(v, 1, 3)
		switch false {
		case mid:
			c.ResponseJson(11, "Invalid template type id", w, r)
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
		c.Cache().ScanDel("allTemplateType")
		c.ResponseJson(1, "", w, r)
	}
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
	tptype := []sql.EnterpriseTptype{}
	if rows, found := c.Cache().Get("allTemplateType"); found {
		tptype = rows.([]sql.EnterpriseTptype)
	} else {
		tptypeSqls := "SELECT * FROM hm_enterprise_tptype"
		err := c.Sql().Select(&tptype, tptypeSqls)
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		}
		c.Cache().SetAlwaysTime("allTemplateType", tptype)
	}
	data := make(map[string]interface{})
	ar := sql.RecursiveNavLevel(nav, 0, 0)
	data["json"] = c.RowsJson(ar)
	data["tptype"] = tptype
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
	image, externalLink, dir, typ, template := r.PostFormValue("image"), r.PostFormValue("externalLink"), r.PostFormValue("dir"), r.PostFormValue("templateType"), r.PostFormValue("template")
	seoTitle, seoKeyword, seoDescribe := r.PostFormValue("seoTitle"), r.PostFormValue("seoKeyword"), r.PostFormValue("seoDescribe")
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
	case vali.Url(url) && vali.Length(url, 0, 50):
		c.ResponseJson(16, "Invalid nav url", w, r)
		return
	case vali.EnglishSpace(en) && vali.Length(en, 2, 50):
		c.ResponseJson(17, "Invalid english nav name", w, r)
		return
	default:
		nav := sql.EnterpriseNav{}
		err := c.Sql().Get(&nav, "SELECT * FROM hm_enterprise_nav WHERE name = ? OR url = ? OR en = ?", name, url, en)
		if err != nil {
			sqls := "INSERT INTO hm_enterprise_nav(url,name,image,external_link,dir,type,template,seo_title,seo_keyword,seo_describe,pid,isshow,sort,icon,level,en) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
			tx := c.Sql().MustBegin()
			tx.MustExec(sqls, url, name, image, externalLink, dir, typ, template, seoTitle, seoKeyword, seoDescribe, pid, isshow, sort, icon, 1, en)
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
	tptype := []sql.EnterpriseTptype{}
	if rows, found := c.Cache().Get("allTemplateType"); found {
		tptype = rows.([]sql.EnterpriseTptype)
	} else {
		tptypeSqls := "SELECT * FROM hm_enterprise_tptype"
		err := c.Sql().Select(&tptype, tptypeSqls)
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		}
		c.Cache().SetAlwaysTime("allTemplateType", tptype)
	}
	ar := sql.RecursiveNavLevel(allnav, 0, 0)
	data["json"] = c.RowsJson(ar)
	data["tptype"] = tptype
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
	image, externalLink, dir, typ, template := r.PostFormValue("image"), r.PostFormValue("externalLink"), r.PostFormValue("dir"), r.PostFormValue("templateType"), r.PostFormValue("template")
	seoTitle, seoKeyword, seoDescribe := r.PostFormValue("seoTitle"), r.PostFormValue("seoKeyword"), r.PostFormValue("seoDescribe")
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
	case vali.Url(url) && vali.Length(url, 0, 50):
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
			sqls := "UPDATE hm_enterprise_nav SET url=?,name=?,image=?,external_link=?,dir=?,type=?,template=?,seo_title=?,seo_keyword=?,seo_describe=?,pid=?,isshow=?,sort=?,icon=?,level=?,en=? WHERE id = ?"
			tx := c.Sql().MustBegin()
			tx.MustExec(sqls, url, name, image, externalLink, dir, typ, template, seoTitle, seoKeyword, seoDescribe, pid, isshow, sort, icon, 1, en, id)
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

/**
 * @description 文本模型页面
 * @English	text model page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) TextModel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "navId")
	data := make(map[string]interface{})
	data["nid"] = id
	c.Template(w, r, data, "./view/admin/enterprise/textModel.html")
}

/**
 * @description 添加文本数据
 * @English	add text data
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) AddTextData(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "navId")
	data := make(map[string]interface{})
	data["nid"] = id
	c.Template(w, r, data, "./view/admin/enterprise/textModel/addTextData.html")
}

/**
 * @description 添加文本数据提交
 * @English	add text data submit
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) AddTextDataSubmit(w http.ResponseWriter, r *http.Request) {
	cnTitle, cnText := r.PostFormValue("cn_title"), r.PostFormValue("cn_text")
	enTitle, enText := r.PostFormValue("en_title"), r.PostFormValue("en_text")
	icon, nid, sort := r.PostFormValue("icon"), r.PostFormValue("nid"), r.PostFormValue("sort")
	switch false {
	case vali.Article(cnTitle) && vali.Length(cnTitle, 2, 255):
		c.ResponseJson(11, "invalid chinese title", w, r)
		return
	case vali.Article(cnText) && vali.Length(cnText, 2, 255):
		c.ResponseJson(12, "invalid chinese text", w, r)
		return
	case vali.Article(enTitle) && vali.Length(enTitle, 2, 255):
		c.ResponseJson(13, "invalid english title", w, r)
		return
	case vali.Article(enText) && vali.Length(enText, 2, 255):

		c.ResponseJson(14, "invalid english text", w, r)
		return
	case vali.NumericNoHeadZero(nid) && vali.Length(nid, 1, 6):
		c.ResponseJson(13, "invalid nav id", w, r)
		return
	default:
		sqls := "INSERT INTO hm_enterprise_textmodel(cn_title,cn_text,en_title,en_text,icon,nid,sort) VALUES(?,?,?,?,?,?,?)"
		tx := c.Sql().MustBegin()
		tx.MustExec(sqls, cnTitle, cnText, enTitle, enText, icon, nid, sort)
		err := tx.Commit()
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		} else {
			c.Cache().ScanDel("alltextmodel")
			c.ResponseJson(1, "", w, r)
		}
	}
}

/**
 * @description 获取文本数据
 * @English	get text data
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) GetTextData(w http.ResponseWriter, r *http.Request) {
	// nav := []sql.EnterpriseTextmodel{}
	// if rows, found := c.Cache().Get("alltextmodel"); found {
	// 	nav = rows.([]sql.EnterpriseTextmodel)
	// } else {
	// 	sqls := "SELECT * FROM hm_enterprise_textmodel ORDER BY sort ASC"
	// 	err := c.Sql().Select(&nav, sqls)
	// 	if err != nil {
	// 		c.ResponseJson(4, err.Error(), w, r)
	// 	}
	// 	c.Cache().SetAlwaysTime("alltextmodel", nav)
	// }
	// fmt.Println(nav)
	// // an := sql.RecursiveNavLevel(nav, 0, 0)
	// fmt.Fprint(w, c.RowsJson(an))
	data := make(map[string]interface{}, 4)
	name := r.PostFormValue("name")
	switch {
	case ((vali.Article(name) && vali.Length(name, 2, 20)) == false) && name != "":
		data["status"], data["info"] = 11, "Invalid Slider Name"
	default:
		addsql := ""
		if name != "" {
			addsql = addsql + "cn_title LIKE '%" + name + "%' AND "
		}
		if addsql != "" {
			addsql = "WHERE " + strings.Trim(addsql, "AND ")
		}
		slides := []sql.EnterpriseTextmodel{}
		if rows, found := c.Cache().Get("alltextmodel" + name); found {
			slides = rows.([]sql.EnterpriseTextmodel)
		} else {
			slideSql := "SELECT * FROM hm_enterprise_textmodel " + addsql + " ORDER BY sort ASC "
			err := c.Sql().Select(&slides, slideSql)
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			}
			c.Cache().SetAlwaysTime("alltextmodel"+name, slides)
		}
		data["status"] = 1
		data["rows"] = slides
	}
	c.ResponseData(data, w, r)
}

/**
 * @description 基础信息
 * @English	base info page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) BaseInfo(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["websiteName"] = c.Config().Value("enterpriseSetup", "websiteName")
	data["websiteLogo"] = c.Config().Value("enterpriseSetup", "websiteLogo")
	data["urlIco"] = c.Config().Value("enterpriseSetup", "urlIco")
	data["websiteAddress"] = c.Config().Value("enterpriseSetup", "websiteAddress")
	data["websiteTitle"] = c.Config().Value("enterpriseSetup", "websiteTitle")
	data["websiteKeyword"] = c.Config().Value("enterpriseSetup", "websiteKeyword")
	data["websiteDescribe"] = c.Config().Value("enterpriseSetup", "websiteDescribe")
	data["copyrightInfo"] = c.Config().Value("enterpriseSetup", "copyrightInfo")
	data["filingNumber"] = c.Config().Value("enterpriseSetup", "filingNumber")
	data["qrcodeDownloadApp"] = c.Config().Value("enterpriseSetup", "qrcodeDownloadApp")
	c.Template(w, r, data, "./view/admin/enterprise/baseInfo.html")
}

/**
 * @description 修改基础信息
 * @English	edit base info
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) EditBaseInfo(w http.ResponseWriter, r *http.Request) {
	websiteName, websiteLogo := r.PostFormValue("websiteName"), r.PostFormValue("websiteLogo")
	urlIco, websiteAddress := r.PostFormValue("urlIco"), r.PostFormValue("websiteAddress")
	websiteTitle, websiteKeyword := r.PostFormValue("websiteTitle"), r.PostFormValue("websiteKeyword")
	websiteDescribe, copyrightInfo := r.PostFormValue("websiteDescribe"), r.PostFormValue("copyrightInfo")
	filingNumber, qrcodeDownloadApp := r.PostFormValue("filingNumber"), r.PostFormValue("qrcodeDownloadApp")
	c.Config().SetValue("enterpriseSetup", "websiteName", websiteName)
	c.Config().SetValue("enterpriseSetup", "websiteLogo", websiteLogo)
	c.Config().SetValue("enterpriseSetup", "urlIco", urlIco)
	c.Config().SetValue("enterpriseSetup", "websiteAddress", websiteAddress)
	c.Config().SetValue("enterpriseSetup", "websiteTitle", websiteTitle)
	c.Config().SetValue("enterpriseSetup", "websiteKeyword", websiteKeyword)
	c.Config().SetValue("enterpriseSetup", "websiteDescribe", websiteDescribe)
	c.Config().SetValue("enterpriseSetup", "copyrightInfo", copyrightInfo)
	c.Config().SetValue("enterpriseSetup", "filingNumber", filingNumber)
	c.Config().SetValue("enterpriseSetup", "qrcodeDownloadApp", qrcodeDownloadApp)
	err := c.Config().Save("./ini/hemacms.ini")
	if err != nil {
		c.ResponseJson(4, err.Error(), w, r)
	} else {
		err = c.Config().Reload()
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		} else {
			c.ResponseJson(1, "", w, r)
		}
	}
}

/**
 * @description 幻灯片管理页面
 * @English	slider manage page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) SliderManage(w http.ResponseWriter, r *http.Request) {
	c.Template(w, r, nil, "./view/admin/enterprise/sliderManage.html")
}

/**
 * @description 获取幻灯片
 * @English	get slider
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) GetSlider(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{}, 4)
	name := r.PostFormValue("name")
	switch {
	case ((vali.Article(name) && vali.Length(name, 2, 20)) == false) && name != "":
		data["status"], data["info"] = 11, "Invalid Slider Name"
	default:
		addsql := ""
		if name != "" {
			addsql = addsql + "name LIKE '%" + name + "%' AND "
		}
		if addsql != "" {
			addsql = "WHERE " + strings.Trim(addsql, "AND ")
		}
		slides := []sql.EnterpriseSlider{}
		if rows, found := c.Cache().Get("enterpriseSlider" + name); found {
			slides = rows.([]sql.EnterpriseSlider)
		} else {
			slideSql := "SELECT * FROM hm_enterprise_slider " + addsql + " ORDER BY sort ASC "
			err := c.Sql().Select(&slides, slideSql)
			if err != nil {
				c.ResponseJson(4, err.Error(), w, r)
			}
			c.Cache().SetAlwaysTime("enterpriseSlider"+name, slides)
		}
		data["status"] = 1
		data["rows"] = slides
	}
	c.ResponseData(data, w, r)
}

/**
 * @description 添加幻灯片
 * @English	add slider
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) AddSlider(w http.ResponseWriter, r *http.Request) {
	c.Template(w, r, nil, "./view/admin/enterprise/addSlider.html")
}

/**
 * @description 添加幻灯片提交
 * @English	add slider submit
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) AddSliderSubmit(w http.ResponseWriter, r *http.Request) {
	name, url := r.PostFormValue("name"), r.PostFormValue("url")
	sort, image := r.PostFormValue("sort"), r.PostFormValue("image")
	// switch false {
	// case vali.Article(cn) && vali.Length(cn, 2, 255):
	// 	c.ResponseJson(11, "invalid chinese content", w, r)
	// 	return
	// case vali.Article(en) && vali.Length(en, 2, 255):
	// 	c.ResponseJson(12, "invalid english content", w, r)
	// 	return
	// default:
	sqls := "INSERT INTO hm_enterprise_slider(name,url,sort,image) VALUES(?,?,?,?)"
	tx := c.Sql().MustBegin()
	tx.MustExec(sqls, name, url, sort, image)
	err := tx.Commit()
	if err != nil {
		c.ResponseJson(4, err.Error(), w, r)
	} else {
		c.Cache().ScanDel("enterpriseSlider")
		c.ResponseJson(1, "", w, r)
	}
	// }
}

/**
 * @description 修改幻灯片
 * @English	edit slider
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) EditSlider(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{}, 2)
	id := chi.URLParam(r, "sliderId")
	sliderSingle := sql.EnterpriseSlider{}
	err := c.Sql().Get(&sliderSingle, "SELECT * FROM hm_enterprise_slider WHERE id = ?", id)
	if err != nil {
		c.ResponseJson(4, err.Error(), w, r)
	}
	data["slider"] = sliderSingle
	c.Template(w, r, data, "./view/admin/enterprise/editSlider.html")
}

/**
 * @description 修改幻灯片提交
 * @English	edit slider submit
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) EditSliderSubmit(w http.ResponseWriter, r *http.Request) {
	id, name, url := r.PostFormValue("id"), r.PostFormValue("name"), r.PostFormValue("url")
	sort, image := r.PostFormValue("sort"), r.PostFormValue("image")
	// switch false {
	// case vali.Article(cn) && vali.Length(cn, 2, 255):
	// 	c.ResponseJson(11, "invalid chinese content", w, r)
	// 	return
	// case vali.Article(en) && vali.Length(en, 2, 255):
	// 	c.ResponseJson(12, "invalid english content", w, r)
	// 	return
	// case vali.NumericNoHeadZero(id) && vali.Length(id, 1, 8):
	// 	c.ResponseJson(13, "invalid update log id", w, r)
	// 	return
	// default:
	sqls := "UPDATE hm_enterprise_slider SET name=?,url=?,sort=?,image=? WHERE id = ?"
	tx := c.Sql().MustBegin()
	tx.MustExec(sqls, name, url, sort, image, id)
	err := tx.Commit()
	if err != nil {
		c.ResponseJson(4, err.Error(), w, r)
	} else {
		c.Cache().ScanDel("enterpriseSlider")
		c.ResponseJson(1, "", w, r)
	}
	// }
}

/**
 * @description 删除单个幻灯
 * @English	delete single slider
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) DelSlider(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	if vali.NumericNoHeadZero(id) && vali.Length(id, 1, 8) {
		sliderSql := "DELETE FROM hm_enterprise_slider WHERE id = ?"
		tx := c.Sql().MustBegin()
		tx.MustExec(sliderSql, id)
		err := tx.Commit()
		if err != nil {
			c.ResponseJson(4, err.Error(), w, r)
		} else {
			c.Cache().ScanDel("enterpriseSlider")
			c.ResponseJson(1, "", w, r)
		}
	} else {
		c.ResponseJson("11", "Invalid slider id", w, r)
	}
}

/**
 * @description 批量删除幻灯
 * @English	batch delete slider
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) BatchDelSlider(w http.ResponseWriter, r *http.Request) {
	ids := r.PostFormValue("ids")
	idSplit := strings.Split(ids, ",")
	for _, v := range idSplit {
		if (vali.NumericNoHeadZero(v) && vali.Length(v, 1, 8)) == false {
			c.ResponseJson(11, "Invalid slider id", w, r)
			return
		}
	}
	sliderSql := fmt.Sprintf("DELETE FROM hm_enterprise_slider WHERE id IN (%v)", ids)
	tx := c.Sql().MustBegin()
	tx.MustExec(sliderSql)
	err := tx.Commit()
	if err != nil {
		c.ResponseJson(4, err.Error(), w, r)
	} else {
		c.Cache().ScanDel("enterpriseSlider")
		c.ResponseJson(1, "", w, r)
	}
}

/**
 * @description 幻灯片批量排序
 * @English	batch sorting slider
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *EnterpriseCtrl) SortSlider(w http.ResponseWriter, r *http.Request) {
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
	bf.WriteString("UPDATE hm_enterprise_slider SET sort = CASE id ")
	// nav := make(map[string]string, 0)
	for _, v := range menuk {
		mid := vali.NumericNoHeadZero(v) && vali.Length(v, 1, 8)
		msort := vali.NumericNoHeadZero(sortNav[v]) && vali.Length(v, 1, 3)
		switch false {
		case mid:
			c.ResponseJson(11, "Invalid slider id", w, r)
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
		c.Cache().ScanDel("enterpriseSlider")
		c.ResponseJson(1, "", w, r)
	}
}
