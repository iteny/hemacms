var HmObj = function() {
    this.config = {
        language: 'cn', //js设置国际化
        animateRandom: false, //是否开启动画随机，默认true开启，false关闭
    };
};
/**
 * @description 设置语言，默认为cn
 * @param lg 什么语言，如果在下列选项中没有，那么将设置成cn
 */
HmObj.prototype.languageSet = function(lg) {
    var that = this;
    if (lg == 'cn') {
        that.config.language = 'cn';
    } else if (lg == 'en') {
        that.config.language = 'en';
    } else {
        that.config.language = 'cn';
    }
    return that;
};
/**
 * @description 设置语言，默认为cn
 * @param time cookie设定个时间，不填则为3天
 */
HmObj.prototype.languageRun = function(time) {
    var that = this;
    if ($.cookie('stage-language') == null) {
        $.cookie('stage-language', that.config.language, {
            expires: time ? time : 3, //有限日期，可以是一个整数或一个日期(单位：天)。这个地方也要注意，如果不设置这个东西，浏览器关闭之后此cookie就失效了
            // path: "/", //cookie值保存的路径，默认与创建页路径一致。
            // domin: , //cookie域名属性，默认与创建页域名一样。这个地方要相当注意，跨域的概念，如果要主域名二级域名有效则要设置".xxx.com"
            // secure: true //一个布尔值，表示传输cookie值时，是否需要一个安全协议。
        })
    }
};

//创建HmMenuTab对象
var hm = new HmObj();
//如果后台没有设置语言，JS来设置
hm.languageSet('en');
hm.languageRun();