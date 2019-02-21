package enterprise

import (
	"hemacms/common"
	"hemacms/common/sql"
	"net/http"
	"strings"
)

type IndexCtrl struct {
	common.BaseCtrl
}

func IndexCtrlObject() *IndexCtrl {
	return &IndexCtrl{}
}
func (c *IndexCtrl) Index(w http.ResponseWriter, r *http.Request) {
	addr := strings.Split(r.RequestURI, "/")
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
	an := sql.RecursiveNavLevel(nav, 0, 0)
	data["json"] = an
	data["active"] = addr[2]
	cookie, err := r.Cookie("stage-language")
	if err != nil {
		c.Log().CheckErr("cookie error", err)
	} else {
		data["cookie"] = cookie.Value
	}
	c.Template(w, r, data, "./template/enterprise/default/index/index.html")
}
