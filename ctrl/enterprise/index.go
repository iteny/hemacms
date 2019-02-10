package enterprise

import (
	"hemacms/common"
	"hemacms/common/sql"
	"net/http"
)

type IndexCtrl struct {
	common.BaseCtrl
}

func IndexCtrlObject() *IndexCtrl {
	return &IndexCtrl{}
}
func (c *IndexCtrl) Index(w http.ResponseWriter, r *http.Request) {
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
	data["json"] = ar
	c.Template(w, r, data, "./template/enterprise/default/index/index.html")
}
