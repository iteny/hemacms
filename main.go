package main

import (
	"fmt"
	"hemacms/common"
	"hemacms/ctrl/admin"
	"hemacms/middle"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/shirou/gopsutil/mem"
)

type server struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func main() {
	//检测内存信息
	v, _ := mem.VirtualMemory()
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	server := &server{Addr: "80", ReadTimeout: 10, WriteTimeout: 10}
	if servPort := common.Config().Value("servSet", "port"); servPort != "" {
		server.Addr = servPort
	}
	if servReadTimeout := common.Config().Int("servSet", "readTimeout"); servReadTimeout != 0 {
		server.ReadTimeout = time.Duration(servReadTimeout)
	}
	if servWriteTimeout := common.Config().Int("servSet", "writeTimeout"); servWriteTimeout != 0 {
		server.WriteTimeout = time.Duration(servWriteTimeout)
	}
	//路由设置
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.NotFound(middle.NoFound)
	// r.Mount("/debug", middleware.Profiler())
	r.Mount("/intendant", adminRoutes())
	//Easily serve static files
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	FileServer(r, "/static", http.Dir(filesDir))
	uploadDir := filepath.Join(workDir, "upload")
	FileServer(r, "/upload", http.Dir(uploadDir))
	//http server
	s := &http.Server{
		Addr:           ":" + server.Addr,
		Handler:        r,
		ReadTimeout:    server.ReadTimeout * time.Second,
		WriteTimeout:   server.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}
	fs := http.StripPrefix(path, http.FileServer(root))
	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"
	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}

//后台路由设置
func adminRoutes() http.Handler {
	// login := admin.LoginCtrlObject()
	index := admin.IndexCtrlObject()
	site := admin.SiteCtrlObject()
	login := admin.LoginCtrlObject()
	r := chi.NewRouter()
	r.Use(middle.LoginVerify)
	r.Get("/login", login.Login) //登录页
	r.Post("/login", login.LoginSubmit)
	r.Post("/loginOut", login.LoginOut)   //退出
	r.Get("/index", index.Index)          //后台主页
	r.Get("/home", index.Home)            //首页
	r.Post("/tabNoAuth", index.TabNoAuth) //tab权限
	r.Route("/site", func(r chi.Router) {
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
	})

	// r.Post("/login", admin.Login.Login) // sign in commit page
	return r
}
