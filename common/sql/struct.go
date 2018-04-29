package sql

import (
	"bytes"
	"strconv"
	"strings"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Nickname   string `json:"nickname"`
	Email      string `json:"email"`
	CreateTime int    `json:"create_time" db:"create_time"`
	CreateIp   string `json:"create_ip" db:"create_ip"`
	Remark     string `json:"remark"`
	Status     int    `json:"status"`
	Role       string `json:"role"`
	RoleEn     string `json:"en"`
}
type AuthRole struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Rules  string `json:"rules"`
	Sort   int    `json:"sort"`
	Remark string `json:"remark"`
	En     string `json:"en"`
}
type AuthRoleAccess struct {
	Uid    int
	RoleId int `db:"role_id"`
	role   AuthRole
}

//table hm_auth_rule
type AuthRule struct {
	Id       int        `json:"id"`
	Url      string     `json:"url"`
	Name     string     `json:"text"`
	Pid      int        `json:"pid"`
	Isshow   int        `json:"isshow"`
	Sort     int        `json:"sort"`
	Icon     string     `json:"iconCls"`
	Level    int        `json:"level"`
	En       string     `json:"en"`
	Checked  int        `json:"checked"`
	Children []AuthRule `json:"children"`
}

//登录日志
type LoginLog struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	LoginTime int    `json:"login_time" db:"login_time"`
	LoginIp   string `json:"login_ip" db:"login_ip"`
	Status    int    `json:"status"`
	Area      string `json:"area"`
	Country   string `json:"country"`
	Useragent string `json:"useragent"`
	Uid       int    `json:"uid"`
}

//操作日志
type OprateLog struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	OprateTime int    `json:"oprate_time" db:"oprate_time"`
	OprateIp   string `json:"oprate_ip" db:"oprate_ip"`
	Useragent  string `json:"useragent"`
	Detail     string `json:"detail"`
	Info       string `json:"info"`
	Url        string `json:"url"`
	Method     string `json:"method"`
	ExcuteTime string `json:"excute_time" db:"excute_time"`
	Status     int    `json:"status"`
}

//更新日志
type UpdateLog struct {
	Id   int    `json:"id"`
	Time int    `json:"time"`
	Cn   string `json:"cn"`
	En   string `json:"en"`
}

//递归重新排序无限极分类
func RecursiveMenu(arr []AuthRule, pid int, level int) (ar []AuthRule) {
	array := make([]AuthRule, 0)
	for k, v := range arr {
		if pid == v.Pid {

			arr[k].Level = level + 1
			// fmt.Println(arr[k])
			array = append(array, arr[k])
			// fmt.Printf("%#v", array)

		}
	}
	for tk, tv := range array {
		rm := RecursiveMenu(arr, tv.Id, level+1)
		for sk, _ := range rm {
			array[tk].Children = append(array[tk].Children, rm[sk])
		}

		// array = append(array, rm[km])
		// array[km].Level = array[km].Level + 1
	}
	return array
}
func RecursiveMenuLevel(arr []AuthRule, pid int, level int) (ar []AuthRule) {
	array := make([]AuthRule, 0)
	for k, v := range arr {
		if pid == v.Pid {

			arr[k].Level = level + 1
			// fmt.Println(arr[k])
			array = append(array, arr[k])
			// fmt.Printf("%#v", array)

		}
	}
	for tk, tv := range array {
		rm := RecursiveMenuLevel(arr, tv.Id, level+1)
		for sk, _ := range rm {
			if rm[sk].Level < 4 {
				array[tk].Children = append(array[tk].Children, rm[sk])
			}
		}
		// array = append(array, rm[km])
		// array[km].Level = array[km].Level + 1
	}
	return array
}

//递归菜单ID
func RecursiveMenuId(arr []AuthRule, id string) string {
	var bf bytes.Buffer
	bf.WriteString(id)
	bf.WriteString(",")
	meid, _ := strconv.Atoi(id)
	for _, v := range arr {
		if meid == v.Pid {
			tid := strconv.Itoa(v.Id)
			bf.WriteString(tid)
			bf.WriteString(",")
			for _, tv := range arr {
				if v.Id == tv.Pid {
					sid := strconv.Itoa(tv.Id)
					bf.WriteString(sid)
					bf.WriteString(",")
				}
			}
		}
	}
	ids := strings.TrimRight(bf.String(), ",")
	return ids
}
