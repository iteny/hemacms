package home

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
func (c *IndexCtrl) Skip(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/enterprise/index", http.StatusFound)
}
