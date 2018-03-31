/**
 * @description 配置文件INI
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
package common

import (
	"fmt"
	"time"

	"github.com/go-ini/ini"
)

type configCtrl struct {
}

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("./ini/hemacms.ini")
	if err != nil {
		fmt.Println(err.Error())
	}

}

//regexp return a regexpCtrl对象
func Config() *configCtrl {
	return &configCtrl{}
}

//regexp return a regexpCtrl对象
func (c *BaseCtrl) Config() *configCtrl {
	return &configCtrl{}
}

/**
 * @description 获取缓存，返回一个string
 * @param section 选择器，选择一个配置分组
 * @param key 获取缓存的键名
 * @return string类型
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *configCtrl) Value(section string, key string) string {
	val := cfg.Section(section).Key(key).Value()
	return val
}

/**
 * @description 获取缓存，返回一个int
 * @param section 选择器，选择一个配置分组
 * @param key 获取缓存的键名
 * @return int类型和错误
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *configCtrl) Int(section string, key string) int {
	val, err := cfg.Section(section).Key(key).Int()
	if err != nil {
		Log().CheckErrSkip("Config", 2, err)
	}
	return val
}

/**
 * @description 获取缓存，返回一个int64
 * @param section 选择器，选择一个配置分组
 * @param key 获取缓存的键名
 * @return int64类型和错误
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *configCtrl) Int64(section string, key string) int64 {
	val, err := cfg.Section(section).Key(key).Int64()
	if err != nil {
		Log().CheckErrSkip("Config", 2, err)
	}
	return val
}

/**
 * @description 获取缓存，返回一个time.Duration
 * @param section 选择器，选择一个配置分组
 * @param key 获取缓存的键名
 * @return time.Duration类型和错误
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *configCtrl) Duration(section string, key string) time.Duration {
	val, err := cfg.Section(section).Key(key).Duration()
	if err != nil {
		Log().CheckErrSkip("Config", 2, err)
	}
	return val
}

/**
 * @description 获取缓存，返回一个time.Duration
 * @param section 选择器，选择一个配置分组
 * @param key 获取缓存的键名
 * @return time.Duration类型
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *configCtrl) MustDuration(section string, key string) time.Duration {
	val := cfg.Section(section).Key(key).MustDuration()
	return val
}
func (c *configCtrl) SetValue(section string, key string, value string) {
	cfg.Section(section).Key(key).SetValue(value)
}

/**
 * @description 保存配置文件
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *configCtrl) Save(filename string) error {
	err := cfg.SaveTo(filename)
	return err
}

/**
 * @description 重新加载配置文件
 * @homepage    http://www.hemacms.com/
 * @author Nicholas Mars
 * @date 2018-02-18
 */
func (c *configCtrl) Reload() (err error) {
	err = cfg.Reload()
	return err
}
