/**
 * @description 正则数据验证
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
package common

import "regexp"

//regexpCtrl is local regexp object.
type regexpCtrl struct {
}

//regexp return a regexpCtrl对象
func Regexp() *regexpCtrl {
	return &regexpCtrl{}
}

//regexp return a regexpCtrl对象
func (c *BaseCtrl) Regexp() *regexpCtrl {
	return &regexpCtrl{}
}

/**
 * @description 验证用户名必须字母开头，5-16字符，字母数字下划线，验证通过返回true,失败false
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) Username(string string) bool {
	r := regexp.MustCompile(`(^[a-zA-Z][a-zA-Z0-9_]{4,15}$)`)
	return r.MatchString(string)
}

/**
 * @description 验证密码必须字母开头，5-16字符，字母数字下划线，验证通过返回true,失败false
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) Password(string string) bool {
	r := regexp.MustCompile(`(^[a-zA-Z][a-zA-Z0-9_]{4,15}$)`)
	return r.MatchString(string)
}

/**
 * @description 昵称只允许汉字、英文字母、数字、减号及下划线、2-80位字符，验证通过返回true,失败false
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) Nickname(string string) bool {
	r := regexp.MustCompile(`(^[\x{4e00}-\x{9fa5}A-Za-z0-9-_]{2,80}$)`)
	return r.MatchString(string)
}

/**
 * @description 名称只允许汉字、英文字母、数字、减号、空格及下划线、1-80位字符，验证通过返回true,失败false
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) Name(string string) bool {
	r := regexp.MustCompile(`(^[\x{4e00}-\x{9fa5}A-Za-z0-9-_ ]{1,80}$)`)
	return r.MatchString(string)
}

/**
 * @description 验证邮箱
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) Email(string string) bool {
	r := regexp.MustCompile(`(^([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,3}$)`)
	return r.MatchString(string)
}

/**
 * @description 备注只允许汉字、英文字母、数字、减号、逗号、句号、感叹号及下划线、2-255位字符，验证通过返回true,失败false
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) Remark(string string) bool {
	r := regexp.MustCompile(`(^[\x{4e00}-\x{9fa5}A-Za-z0-9-_,.!，。！ ]{2,255}$)`)
	return r.MatchString(string)
}

/**
 * @description 验证是否为1或0，只要不是1和0都返回false
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) Status(string string) bool {
	switch string {
	case "1":
		return true
	case "0":
		return true
	default:
		return false
	}
}

/**
 * @description 匹配数据库类型配置
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) SqlType(str string) bool {
	switch str {
	case "sqlite":
		return true
	case "mysql":
		return true
	default:
		return false
	}
}

/**
 * @description 端口号必须是1-5位正整数
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) Port(string string) bool {
	r := regexp.MustCompile(`(^[0-9]{0,5}$)`)
	return r.MatchString(string)
}

/**
 * @description 读取超时必须是1-3位正整数
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) ReadTimeout(string string) bool {
	r := regexp.MustCompile(`(^[0-9]{0,3}$)`)
	return r.MatchString(string)
}

/**
 * @description 写入超时必须是1-3位正整数
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) WriteTimeout(string string) bool {
	r := regexp.MustCompile(`(^[0-9]{0,3}$)`)
	return r.MatchString(string)
}

/**
 * @description 验证id数组
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) IdGroup(string string) bool {
	r := regexp.MustCompile(`(^[\d,]{1,255}$)`)
	return r.MatchString(string)
}

/**
 * @description 验证id
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) Id(string string) bool {
	r := regexp.MustCompile(`(^[\d]{1,8}$)`)
	return r.MatchString(string)
}

/**
 * @description 排序号验证
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) Sort(string string) bool {
	r := regexp.MustCompile(`(^[\d]{1,3}$)`)
	return r.MatchString(string)
}

/**
 * @description 菜单名称只允许汉字、英文字母及空格、2-80位字符，验证通过返回true,失败false
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) MenuName(string string) bool {
	r := regexp.MustCompile(`(^[\x{4e00}-\x{9fa5}A-Za-z0-9 ]{2,80}$)`)
	return r.MatchString(string)
}

/**
 * @description 验证英文
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) English(string string) bool {
	r := regexp.MustCompile(`(^[A-Za-z0-9 ]{2,80}$)`)
	return r.MatchString(string)
}

/**
 * @description 验证图标格式
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) Icon(string string) bool {
	r := regexp.MustCompile(`(^[A-Za-z0-9-_]{2,80}$)`)
	return r.MatchString(string)
}

/**
 * @description 验证后台URL
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) UrlAdmin(string string) bool {
	r := regexp.MustCompile(`(^[A-Za-z0-9-/]{2,80}$)`)
	return r.MatchString(string)
}

/**
 * @description 验证时间
 * @param string 需要验证的字符串
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *regexpCtrl) Date(string string) bool {
	r := regexp.MustCompile(`(^\d{4}[\-](0?[1-9]|1[012])[\-](0?[1-9]|[12][0-9]|3[01])(\s+(0?[0-9]|1[0-9]|2[0-3])\:(0?[0-9]|[1-5][0-9])\:(0?[0-9]|[1-5][0-9]))?$)`)
	return r.MatchString(string)
}

/**
 * @description 验证不通过就返回json信息
 * @param bool 接收MatchString返回的验证结果bool类型,这里可以选择这个包内封装的正则验证函数
 * @param string 需要验证的字符串
 * @param status 返回
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
// func (c *regexpCtrl) Stop(bool bool, string string, status interface{}, w io.Writer, r *http.Request) {
// 	bc := &BaseCtrl{}
// 	if bool {
// 		fmt.Println("验证成功")
// 	} else {
// 		bc.ResponseJson(status, string, w, r)
// 		return
// 	}
// }
