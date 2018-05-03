package enterprise

import (
	"hemacms/common"
	"net/http"
)

type IndexCtrl struct {
	common.BaseCtrl
}

func IndexCtrlObject() *IndexCtrl {
	return &IndexCtrl{}
}
func (c *IndexCtrl) Index(w http.ResponseWriter, r *http.Request) {
	c.Template(w, r, nil, "./template/enterprise/default/index/index.html")
}
