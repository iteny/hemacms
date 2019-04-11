package router

import (
	"hemacms/ctrl/admin"
	"hemacms/middle"
	"net/http"

	"github.com/go-chi/chi"
)

/**
 * @description 后台路由器
 * @English	backstage router
 * @homepage http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-03-24
 */
func AdminRoutes() http.Handler {
	index := admin.IndexCtrlObject()
	site := admin.SiteCtrlObject()
	content := admin.ContentCtrlObject()
	enterprise := admin.EnterpriseCtrlObject()
	login := admin.LoginCtrlObject()
	r := chi.NewRouter()
	r.Use(middle.LoginVerify)
	r.Get("/", index.Skip)       //跳转到登录页面
	r.Get("/login", login.Login) //登录页
	r.Post("/login", login.LoginSubmit)
	r.Post("/loginOut", login.LoginOut)           //退出
	r.Get("/index", index.Index)                  //后台主页
	r.Get("/home", index.Home)                    //首页
	r.Post("/ajaxPolling", index.AjaxPolling)     //轮询
	r.Post("/ajaxUpdateLog", index.AjaxUpdateLog) //ajax更新日志
	r.Post("/tabNoAuth", index.TabNoAuth)         //tab权限
	r.Route("/site", func(r chi.Router) {
		r.Post("/uploadImage", site.UploadImage) //图片上传
		r.Post("/delImage", site.DelImage)       //删除图片
		r.Get("/system", site.System)            //系统设置
		r.Post("/editSystem", site.EditSystem)   //系统设置修改
		r.Get("/menu", site.Menu)                //菜单设置
		r.Post("/getMenu", site.GetMenu)         //获取菜单
		r.Post("/iconsCls", site.IconsCls)       //获取图标
		r.Post("/sortMenu", site.SortMenu)       //菜单排序
		r.Route("/addMenu", func(r chi.Router) { //添加菜单
			r.Get("/{menuPid}", site.AddMenu) //添加菜单页面
		})
		r.Route("/editMenu", func(r chi.Router) {
			r.Get("/{menuId}", site.EditMenu) //修改菜单页面
		})
		r.Post("/addMenuSubmit", site.AddMenuSubmit)   //添加菜单提交
		r.Post("/editMenuSubmit", site.EditMenuSubmit) //修改菜单提交
		r.Post("/delMenuSubmit", site.DelMenuSubmit)   //删除菜单提交
		r.Get("/user", site.User)                      //用户管理页面
		r.Post("/getUser", site.GetUser)               //获取用户数据
		r.Get("/addUser", site.AddUser)                //添加用户页面
		r.Post("/addUserSubmit", site.AddUserSubmit)   //添加用户提交
		r.Post("/delUser", site.DelUser)               //删除单个用户
		r.Post("/batchDelUser", site.BatchDelUser)     //批量删除用户
		r.Route("/editUser", func(r chi.Router) {
			r.Get("/{userId}", site.EditUser) //修改用户页面
		})
		r.Post("/editUserSubmit", site.EditUserSubmit) //修改用户提交
		r.Get("/role", site.Role)                      //角色管理页面
		r.Post("/getRole", site.GetRole)               //获取角色数据
		r.Route("/setRole", func(r chi.Router) {
			r.Get("/{roleId}", site.SetRole) //设置角色权限页面
		})
		r.Post("/setRoleSubmit", site.SetRoleSubmit) //设置角色权限提交
		r.Post("/delRole", site.DelRole)             //删除单个角色
		r.Post("/batchDelRole", site.BatchDelRole)   //批量删除角色
		r.Get("/addRole", site.AddRole)              //添加角色页面
		r.Post("/addRoleSubmit", site.AddRoleSubmit) //添加角色提交
		r.Route("/editRole", func(r chi.Router) {
			r.Get("/{roleId}", site.EditRole) //修改角色页面
		})
		r.Post("/editRoleSubmit", site.EditRoleSubmit) //修改角色提交

		r.Get("/colorSchemes", site.ColorSchemes) //配色方案

		r.Get("/loginLog", site.LoginLog)        //登录日志页面
		r.Post("/getLoginLog", site.GetLoginLog) //获取登录日志
		r.Post("/delLoginLog", site.DelLoginLog) //删除一个月前的登录日志

		r.Get("/oprateLog", site.OprateLog)              //操作日志页面
		r.Post("/getOprateLog", site.GetOprateLog)       //获取操作日志
		r.Post("/delOprateLog", site.DelOprateLog)       //删除一个月前的操作日志
		r.Post("/delAllOprateLog", site.DelAllOprateLog) //删除全部操作日志

		r.Get("/updateLog", site.UpdateLog)                    //更新日志页面
		r.Post("/getUpdateLog", site.GetUpdateLog)             //获取更新日志
		r.Get("/addUpdateLog", site.AddUpdateLog)              //添加更新日志页面
		r.Post("/addUpdateLogSubmit", site.AddUpdateLogSubmit) //添加更新日志提交
		r.Route("/editUpdateLog", func(r chi.Router) {
			r.Get("/{updateLogId}", site.EditUpdateLog) //修改更新日志页面
		})
		r.Post("/editUpdateLogSubmit", site.EditUpdateLogSubmit) //修改更新日志提交
		r.Post("/delUpdateLog", site.DelUpdateLog)               //删除单个更新日志
		r.Post("/batchDelUpdateLog", site.BatchDelUpdateLog)     //批量删除更新日志
	})
	r.Route("/content", func(r chi.Router) {
		r.Get("/columnList", content.ColumnList) //栏目列表页面
	})
	r.Route("/enterprise", func(r chi.Router) {
		r.Get("/navList", enterprise.NavList)   //导航列表页面
		r.Post("/getNav", enterprise.GetNav)    //获取导航
		r.Route("/addNav", func(r chi.Router) { //添加导航
			r.Get("/{navPid}", enterprise.AddNav) //添加导航页面
		})
		r.Route("/editNav", func(r chi.Router) {
			r.Get("/{navId}", enterprise.EditNav) //修改导航页面
		})
		r.Post("/sortNav", enterprise.SortNav)                 //导航排序
		r.Post("/addNavSubmit", enterprise.AddNavSubmit)       //添加导航提交
		r.Post("/editNavSubmit", enterprise.EditNavSubmit)     //修改导航提交
		r.Post("/delNavSubmit", enterprise.DelNavSubmit)       //删除导航提交
		r.Get("/contentManage", enterprise.ContentManage)      //内容管理页面
		r.Get("/baseInfo", enterprise.BaseInfo)                //基础信息
		r.Post("/editBaseInfo", enterprise.EditBaseInfo)       //编辑基础信息
		r.Get("/sliderManage", enterprise.SliderManage)        //幻灯片管理页面
		r.Post("/getSlider", enterprise.GetSlider)             //获取幻灯片
		r.Get("/addSlider", enterprise.AddSlider)              //添加幻灯片
		r.Post("/addSliderSubmit", enterprise.AddSliderSubmit) //添加幻灯片提交
		r.Route("/editSlider", func(r chi.Router) {
			r.Get("/{sliderId}", enterprise.EditSlider) //修改幻灯片页面
		})
		r.Post("/editSliderSubmit", enterprise.EditSliderSubmit) //修改幻灯片提交
		r.Post("/delSlider", enterprise.DelSlider)               //删除单个幻灯片
		r.Post("/batchDeLSlider", enterprise.BatchDelSlider)     //批量删除幻灯片
		r.Post("/sortSlider", enterprise.SortSlider)             //幻灯片排序
	})
	return r
}
