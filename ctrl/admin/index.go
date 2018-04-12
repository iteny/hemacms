package admin

import (
	"fmt"
	"hemacms/common"
	"hemacms/common/sql"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	cpuInfo "github.com/shirou/gopsutil/cpu"
	diskInfo "github.com/shirou/gopsutil/disk"
	hostInfo "github.com/shirou/gopsutil/host"
	memInfo "github.com/shirou/gopsutil/mem"
)

type IndexCtrl struct {
	common.BaseCtrl
}

func IndexCtrlObject() *IndexCtrl {
	return &IndexCtrl{}
}
func (c *IndexCtrl) Skip(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/intendant/login", http.StatusFound)
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
	data := make(map[string]interface{})
	session := c.Sess().Load(r)
	username, err := session.GetString("username")
	c.Log().CheckErr("Session Get Error", err)
	rolename, err := session.GetString("rolename")
	c.Log().CheckErr("Session Get Error", err)
	data["username"] = username //用户名
	data["rolename"] = rolename //角色名
	data["ajaxPolling"] = c.Config().Value("common", "ajaxPolling")
	data["arch"] = runtime.GOARCH          //系统架构
	data["serverTime"] = time.Now().Unix() //服务器时间
	host, _ := hostInfo.Info()
	// cpu, _ := cpuInfo.Info()
	// cpuUserd, _ := cpuInfo.Percent(time.Second, false)
	// mem, _ := memInfo.VirtualMemory()
	// disk, _ := diskInfo.Usage("/")
	// data["cpuUserd"] = cpuUserd
	data["uptime"] = host.BootTime - host.Uptime //正常运行时间
	// data["hostname"] = host.Hostname                        //主机名
	// data["procs"] = host.Procs                              //进程号
	// data["os"] = host.Platform + " " + host.PlatformVersion //系统版本号
	// data["kernelVersion"] = host.KernelVersion              //内核版本
	// data["cpuModelName"] = cpu[0].ModelName                 //cpu型号
	// data["serverTime"] = time.Now().Unix()                  //服务器时间
	// data["memTotal"] = mem.Total / 1024 / 1024 / 1024
	// data["memFree"] = mem.Free / 1024 / 1024 / 1024
	// data["memUserd"] = mem.Total/1024/1024/1024 - mem.Free/1024/1024/1024
	// data["memUserdPercent"] = mem.UsedPercent
	// data["diskTotal"] = disk.Total / 1024 / 1024 / 1024
	// data["diskFree"] = disk.Free / 1024 / 1024 / 1024
	// data["diskUserd"] = disk.Total/1024/1024/1024 - disk.Free/1024/1024/1024
	// data["diskUserdPercent"] = disk.UsedPercent
	c.Template(w, r, data, "./view/admin/index/home.html")
}

//轮询系统占用
func (c *IndexCtrl) AjaxPolling(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	mem, _ := memInfo.VirtualMemory()
	disk, _ := diskInfo.Usage("/")
	cpuUserd, _ := cpuInfo.Percent(time.Second, false)
	data["cpuUserd"] = cpuUserd
	data["serverTime"] = time.Now().Unix() //服务器时间
	data["memTotal"] = mem.Total / 1024 / 1024 / 1024
	data["memFree"] = mem.Free / 1024 / 1024 / 1024
	data["memUserd"] = mem.Total/1024/1024/1024 - mem.Free/1024/1024/1024
	data["memUserdPercent"] = mem.UsedPercent
	data["diskTotal"] = disk.Total / 1024 / 1024 / 1024
	data["diskFree"] = disk.Free / 1024 / 1024 / 1024
	data["diskUserd"] = disk.Total/1024/1024/1024 - disk.Free/1024/1024/1024
	data["diskUserdPercent"] = disk.UsedPercent
	data["status"], data["info"] = 1, ""
	c.ResponseData(data, w, r)
}

//tab权限
func (c *IndexCtrl) TabNoAuth(w http.ResponseWriter, r *http.Request) {
	uri := r.PostFormValue("uri")
	session := c.Sess().Load(r)
	userId, err := session.GetString("uid")
	c.Log().CheckErr("Session Get Error", err)
	if userId == "" {
		// fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":67}")
		return
	} else if userId == "1" {
		fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":88}")
		return
	} else {
		//获取所有菜单
		rule := []sql.AuthRule{}
		if rows, found := c.Cache().Get("allmenu"); found {
			rule = rows.([]sql.AuthRule)
		} else {
			sqls := "SELECT * FROM hm_auth_rule ORDER BY sort ASC"
			err := c.Sql().Select(&rule, sqls)
			c.Log().CheckErr("Sql Error", err)
			c.Cache().SetAlwaysTime("allmenu", rule)
		}
		//获取用户的用户组id
		roleId := ""
		if roleIdRow, roleIdFound := c.Cache().Get("roleId" + userId); roleIdFound {
			roleId = roleIdRow.(string)
		} else {
			access := sql.AuthRoleAccess{}
			accessSql := "SELECT * FROM hm_auth_role_access WHERE uid = ?"
			err := c.Sql().Get(&access, accessSql, userId)
			c.Log().CheckErr("Sql Error", err)
			roleId = strconv.Itoa(access.RoleId)
			c.Cache().SetAlwaysTime("roleId"+userId, roleId)
		}
		rulerz := "false"
		formatUri := c.FormatUrl(uri)
		for _, v := range rule {
			if formatUri == "/intendant/"+v.Url {
				rulerz = "true"
				//获取用户权限id组
				auths := ""
				if row, fd := c.Cache().Get("auth" + roleId); fd {
					auths = row.(string)
				} else {
					role := sql.AuthRole{}
					roleSql := "SELECT * FROM hm_auth_role WHERE id = ?"
					err = c.Sql().Get(&role, roleSql, roleId)
					c.Log().CheckErr("Sql Error", err)
					auths = role.Rules
					c.Cache().SetAlwaysTime("auth"+roleId, auths)
				}
				rules := strings.Split(auths, ",")
				renzheng := "false"
				for _, rv := range rules {
					intv := strconv.Itoa(v.Id)
					if rv == intv {
						renzheng = "true"
					}
				}
				if renzheng == "true" {
					fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":88}")
					return
				} else {
					fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":66}")
					return
				}
			}
		}
		if rulerz == "true" {
			fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":88}")
			return
		} else {
			fmt.Fprint(w, "{\"total\":0,\"rows\":[],\"status\":66}")
			return
		}
	}
}
