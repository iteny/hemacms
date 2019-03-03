package router

import (
	"hemacms/ctrl/enterprise"
	"net/http"

	"github.com/go-chi/chi"
)

/**
 * @description 企业路由器
 * @English	Enterprise router
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func HomeRoutes() http.Handler {
	index := enterprise.IndexCtrlObject()
	r := chi.NewRouter()
	r.Get("/index", index.Index) //首页
	// r.Use(middle.LoginVerify)
	r.Get("/{articleSlug:[a-z-]+}", index.Index)
	return r
}
