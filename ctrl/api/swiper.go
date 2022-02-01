package api

import (
	"hemacms/common"
	"hemacms/common/sql"
	"net/http"
	"strings"

	vali "github.com/iteny/hmgo/validator"
)

type ApiCtrl struct {
	common.BaseCtrl
}

func ApiCtrlObject() *ApiCtrl {
	return &ApiCtrl{}
}
func (c *ApiCtrl) GetSlider(w http.ResponseWriter, r *http.Request) {
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
