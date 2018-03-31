package admin

import (
	"fmt"
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
func (c *IndexCtrl) GetMenu(w http.ResponseWriter, r *http.Request) {
	// pid := r.PostFormValue("pid")
	// intpid, _ := strconv.Atoi(pid)
	sqls := "SELECT * FROM hm_auth_rule"
	rule := []sql.AuthRule{}
	err := c.Sql().Select(&rule, sqls)
	if err != nil {
		c.Log().Debug().Err(err).Msg("错误")
		c.ResponseJson(4, err.Error(), w, r)
	}
	ar := sql.RecursiveMenu(rule, 0, 0)
	// fmt.Println(rule)
	// for k, v := range rule {
	// 	srule := []sqlm.AuthRule{}
	// 	err = c.Sql().Select(&srule, sqls, v.Id)
	// 	if err != nil {
	// 		c.Log().Debug().Err(err).Msg("错误")
	// 		c.ResponseJson(4, err.Error(), w, r)
	// 	}
	// 	for tk, _ := range srule {
	// 		rule[k].Children = append(rule[k].Children, srule[tk])
	// 	}
	// }
	// fmt.Println(ar)
	fmt.Fprint(w, c.RowsJson(ar))
}
