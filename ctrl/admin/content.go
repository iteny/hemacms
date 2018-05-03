package admin

import (
	"hemacms/common"
	"net/http"
)

type ContentCtrl struct {
	common.BaseCtrl
}

func ContentCtrlObject() *ContentCtrl {
	return &ContentCtrl{}
}

/**
 * @description 栏目列表
 * @English	menu management page
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func (c *ContentCtrl) ColumnList(w http.ResponseWriter, r *http.Request) {
	c.Template(w, r, nil, "./view/admin/content/columnList.html")
}
