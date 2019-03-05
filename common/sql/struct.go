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

//图片附件表
type Image struct {
	Id   int    `json:"id"`
	Url  string `json:"url"`
	Time string `json:"time"`
	Uid  string `json:"uid"`
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

//企业站导航表 hm_enterprise_nav
type EnterpriseNav struct {
	Id           int             `json:"id"`
	Url          string          `json:"url"`                              //内部链接
	Name         string          `json:"text"`                             //中文名称
	Image        string          `json:"image"`                            //图片
	ExternalLink string          `json:"external_link" db:"external_link"` //外部链接
	Dir          string          `json:"dir"`                              //模板目录
	Type         string          `json:"type"`                             //模板类型
	Template     string          `json:"template"`                         //模板文件
	SeoTitle     string          `json:"seo_title" db:"seo_title"`         //seo标题
	SeoKeyword   string          `json:"seo_keyword" db:"seo_keyword"`     //seo关键字
	SeoDescribe  string          `json:"seo_describe" db:"seo_describe"`   //seo描述
	Pid          int             `json:"pid"`                              //父id
	Isshow       int             `json:"isshow"`                           //是否在首页显示
	Sort         int             `json:"sort"`                             //排序
	Icon         string          `json:"iconCls"`                          //后台图标
	Level        int             `json:"level"`                            //等级
	En           string          `json:"en"`                               //英文名称
	Checked      int             `json:"checked"`
	Children     []EnterpriseNav `json:"children"`
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
		for sk := range rm {
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
		for sk := range rm {
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

//递归重新排序无限极分类
func RecursiveNav(arr []EnterpriseNav, pid int, level int) (ar []EnterpriseNav) {
	array := make([]EnterpriseNav, 0)
	for k, v := range arr {
		if pid == v.Pid {
			arr[k].Level = level + 1
			array = append(array, arr[k])
		}
	}
	for tk, tv := range array {
		rm := RecursiveNav(arr, tv.Id, level+1)
		for sk := range rm {
			array[tk].Children = append(array[tk].Children, rm[sk])
		}
	}
	return array
}
func RecursiveNavLevel(arr []EnterpriseNav, pid int, level int) (ar []EnterpriseNav) {
	array := make([]EnterpriseNav, 0)
	for k, v := range arr {
		if pid == v.Pid {
			arr[k].Level = level + 1
			array = append(array, arr[k])
		}
	}
	for tk, tv := range array {
		rm := RecursiveNavLevel(arr, tv.Id, level+1)
		for sk := range rm {
			if rm[sk].Level < 3 {
				array[tk].Children = append(array[tk].Children, rm[sk])
			}
		}
	}
	return array
}

//递归菜单ID
func RecursiveNavId(arr []EnterpriseNav, id string) string {
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
