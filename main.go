package main

import (
	"hemacms/common"
	"hemacms/ctrl/home"
	"hemacms/middle"
	"hemacms/router"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type server struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func main() {
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
	index := home.IndexCtrlObject()
	//路由设置11
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.NotFound(middle.NoFound)
	// r.Mount("/debug", middleware.Profiler())
	r.Get("/", index.Skip) //跳转到登录页面
	r.Mount("/enterprise", router.HomeRoutes())
	r.Mount("/intendant", router.AdminRoutes())
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
