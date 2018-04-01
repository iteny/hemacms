/**
 * @description 抛出异常错误信息
 * @homepage    http://www.hemacms.com/
 * @copyright   hemacms
 * @author Nicholas Mars
 * @date 2017-09-18
 */
$.extend($.fn.tabs.methods, {
    /**
     * 绑定双击事件
     * @param {Object} jq
     * @param {Object} caller 绑定的事件处理程序
     */
    bindDblclick: function(jq, caller) {
        return jq.each(function() {
            var that = this;
            $(this).children("div.tabs-header").find("ul.tabs").undelegate('li', 'dblclick.tabs').delegate('li', 'dblclick.tabs', function(e) {
                if (caller && typeof(caller) == 'function') {
                    var title = $(this).text();
                    var index = $(that).tabs('getTabIndex', $(that).tabs('getTab', title));
                    caller(index, title);
                }
            });
        });
    },
    /**
     * 解除绑定双击事件
     * @param {Object} jq
     */
    unbindDblclick: function(jq) {
        return jq.each(function() {
            $(this).children("div.tabs-header").find("ul.tabs").undelegate('li', 'dblclick.tabs');
        });
    }
});
//给jQuery增加一个方法animateCss,用于给dom元素增加动画效果
$.fn.extend({
    animateCss: function(animationName) {
        var animationEnd = 'webkitAnimationEnd mozAnimationEnd MSAnimationEnd oanimationend animationend';
        this.addClass('animated ' + animationName).one(animationEnd, function() {
            $(this).removeClass('animated ' + animationName);
        });
        return this;
    }
});
var HmObj = function() {
    this.config = {
        language: 'cn', //js设置国际化
        animateRandom: false, //是否开启动画随机，默认true开启，false关闭
    };
    this.menuCfg = {
        url: undefined, //数据源地址
        type: 'GET', //读取方式
    };
    this.menuData = {}; //菜单数据
    this.language = {}; //语言数据
    //过渡动画
    this.animateIn = {
        1: 'bounce',
        2: 'flash',
        3: 'pulse',
        4: 'rubberBand',
        5: 'shake',
        6: 'headShake',
        7: 'swing',
        8: 'tada',
        9: 'wobble',
        10: 'jello',
        11: 'bounceIn',
        12: 'bounceInDown',
        13: 'bounceInLeft',
        14: 'bounceInRight',
        15: 'bounceInUp',
        16: 'fadeIn',
        17: 'fadeInDown',
        18: 'fadeInDownBig',
        19: 'fadeInLeft',
        20: 'fadeInLeftBig',
        21: 'fadeInRight',
        22: 'fadeInRightBig',
        23: 'fadeInUp',
        24: 'fadeInUpBig',
        25: 'flipInX',
        26: 'flipInY',
        27: 'flipOutX',
        28: 'flipOutY',
        29: 'lightSpeedIn',
        30: 'rotateIn',
        31: 'rotateInDownLeft',
        32: 'rotateInDownRight',
        33: 'rotateInUpLeft',
        34: 'rotateInUpRight',
        35: 'hinge',
        36: 'jackInTheBox',
        37: 'rollIn',
        38: 'zoomIn',
        39: 'zoomInDown',
        40: 'zoomInLeft',
        41: 'zoomInRight',
        42: 'zoomInUp',
        43: 'slideInDown',
        44: 'slideInLeft',
        45: 'slideInRight',
        46: 'slideInUp',
    };
    //消失动画
    this.animateOut = {
        1: 'bounceOut',
        2: 'bounceOutDown',
        3: 'bounceOutLeft',
        4: 'bounceOutRight',
        5: 'bounceOutUp',
        6: 'fadeOut',
        7: 'fadeOutDown',
        8: 'fadeOutDownBig',
        9: 'fadeOutLeft',
        10: 'fadeOutLeftBig',
        11: 'fadeOutRight',
        12: 'fadeOutRightBig',
        13: 'fadeOutUp',
        14: 'fadeOutUpBig',
        15: 'lightSpeedOut',
        16: 'rotateOut',
        17: 'rotateOutDownLeft',
        18: 'rotateOutDownRight',
        19: 'rotateOutUpLeft',
        20: 'rotateOutUpRight',
        21: 'rollOut',
        22: 'zoomOut',
        23: 'zoomOutDown',
        24: 'zoomOutLeft',
        25: 'zoomOutRight',
        26: 'zoomOutUp',
        27: 'slideOutDown',
        28: 'slideOutLeft',
        29: 'slideOutRight',
        30: 'slideOutUp',
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
    if ($.cookie('back-language') == null) {
        $.cookie('back-language', that.config.language, {
            expires: time ? time : 3, //有限日期，可以是一个整数或一个日期(单位：天)。这个地方也要注意，如果不设置这个东西，浏览器关闭之后此cookie就失效了
            // path: "/", //cookie值保存的路径，默认与创建页路径一致。
            // domin: , //cookie域名属性，默认与创建页域名一样。这个地方要相当注意，跨域的概念，如果要主域名二级域名有效则要设置".xxx.com"
            // secure: true //一个布尔值，表示传输cookie值时，是否需要一个安全协议。
        })
    }
};

/**
 * @description 国际化
 * @author Nicholas Mars
 */
HmObj.prototype.readJsonFile = function() {
    var that = this;
    switch ($.cookie('back-language')) {
        case 'cn':
            that.language = language.cn;
            break;
        case 'en':
            that.language = language.en;
            break;
        default:
            that.language = language.cn;
            break;
    }
};
//切换语言
HmObj.prototype.changeLanguage = function() {
    var lang = '',
        that = this;
    switch (that.config.language) {
        case 'cn':
            return 'cn';
            break;
        case 'en':
            return 'en';
            break;
        default:
            return 'en';
            break;
    }
};
//切换后台风格
HmObj.prototype.changeStyle = function(type) {
    switch (type) {
        case 'default':
            $('#hm-app').removeClass('hm_style_mac');
            $('#hm-app').addClass('hm_style_default');
            $.cookie('back-style', 'hm_style_default', {
                expires: 30, //有限日期，可以是一个整数或一个日期(单位：天)。这个地方也要注意，如果不设置这个东西，浏览器关闭之后此cookie就失效了
                // path: "/", //cookie值保存的路径，默认与创建页路径一致。
                // domin: , //cookie域名属性，默认与创建页域名一样。这个地方要相当注意，跨域的概念，如果要主域名二级域名有效则要设置".xxx.com"
                // secure: true //一个布尔值，表示传输cookie值时，是否需要一个安全协议。
            });
            this.notice('success', this.language.styleSwitch + this.language.styleDefault);
            break;
        case 'mac':
            $('#hm-app').removeClass('hm_style_default');
            $('#hm-app').addClass('hm_style_mac');
            $.cookie('back-style', 'hm_style_mac', {
                expires: 30, //有限日期，可以是一个整数或一个日期(单位：天)。这个地方也要注意，如果不设置这个东西，浏览器关闭之后此cookie就失效了
                // path: "/", //cookie值保存的路径，默认与创建页路径一致。
                // domin: , //cookie域名属性，默认与创建页域名一样。这个地方要相当注意，跨域的概念，如果要主域名二级域名有效则要设置".xxx.com"
                // secure: true //一个布尔值，表示传输cookie值时，是否需要一个安全协议。
            });
            this.notice('success', this.language.styleSwitch + this.language.styleMac);
            break;
        default:

    }
};
/**
 * @description 处理easyui国家化问题
 * @author Nicholas Mars
 */
HmObj.prototype.easyuiLanguage = function() {
    var that = this;
    if ($.fn.pagination) {
        $.fn.pagination.defaults.beforePageText = that.language.easyuiBeforePageText ? that.language.easyuiBeforePageText : 'Page';
        $.fn.pagination.defaults.afterPageText = that.language.easyuiAfterPageText ? that.language.easyuiAfterPageText : 'of {pages}';
        $.fn.pagination.defaults.displayMsg = that.language.easyuiDisplayMsg ? that.language.easyuiDisplayMsg : 'Displaying {from} to {to} of {total} items';
    }
    if ($.fn.datagrid) {
        $.fn.datagrid.defaults.loadMsg = that.language.easyuiLoadMsg ? that.language.easyuiLoadMsg : 'Processing, please wait ...';
    }
    if ($.fn.treegrid && $.fn.datagrid) {
        $.fn.treegrid.defaults.loadMsg = $.fn.datagrid.defaults.loadMsg;
    }
    if ($.messager) {
        $.messager.defaults.ok = that.language.easyuiOk ? that.language.easyuiOk : 'Ok';
        $.messager.defaults.cancel = that.language.easyuiCancel ? that.language.easyuiCancel : 'Cancel';
    }
    $.map(['validatebox', 'textbox', 'passwordbox', 'filebox', 'searchbox',
        'combo', 'combobox', 'combogrid', 'combotree',
        'datebox', 'datetimebox', 'numberbox',
        'spinner', 'numberspinner', 'timespinner', 'datetimespinner'
    ], function(plugin) {
        if ($.fn[plugin]) {
            $.fn[plugin].defaults.missingMessage = that.language.easyuiMissingMessage ? that.language.easyuiMissingMessage : 'This field is required.';
        }
    });
    if ($.fn.validatebox) {
        $.fn.validatebox.defaults.rules.email.message = that.language.easyuiEmailMessage ? that.language.easyuiEmailMessage : 'Please enter a valid email address.';
        $.fn.validatebox.defaults.rules.url.message = that.language.easyuiUrlMessage ? that.language.easyuiUrlMessage : 'Please enter a valid URL.';
        $.fn.validatebox.defaults.rules.length.message = that.language.easyuiLengthMessage ? that.language.easyuiLengthMessage : 'Please enter a value between {0} and {1}.';
        $.fn.validatebox.defaults.rules.remote.message = that.language.easyuiRemoteMessage ? that.language.easyuiRemoteMessage : 'Please fix this field.';
    }
    if ($.fn.calendar) {
        $.fn.calendar.defaults.weeks = that.language.easyuiWeeks ? that.language.easyuiWeeks : ['S', 'M', 'T', 'W', 'T', 'F', 'S'];
        $.fn.calendar.defaults.months = that.language.easyuiMonths ? that.language.easyuiMonths : ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
    }
    if ($.fn.datebox) {
        $.fn.datebox.defaults.currentText = that.language.easyuiCurrentText ? that.language.easyuiCurrentText : 'Today';
        $.fn.datebox.defaults.closeText = that.language.easyuiCloseText ? that.language.easyuiCloseText : 'Close';
        $.fn.datebox.defaults.okText = that.language.easyuiOkText ? that.language.easyuiOkText : 'Ok';
        $.fn.datebox.defaults.formatter = function(date) {
            var y = date.getFullYear();
            var m = date.getMonth() + 1;
            var d = date.getDate();
            return y + '-' + (m < 10 ? ('0' + m) : m) + '-' + (d < 10 ? ('0' + d) : d);
        };
        $.fn.datebox.defaults.parser = function(s) {
            if (!s) return new Date();
            var ss = s.split('-');
            var y = parseInt(ss[0], 10);
            var m = parseInt(ss[1], 10);
            var d = parseInt(ss[2], 10);
            if (!isNaN(y) && !isNaN(m) && !isNaN(d)) {
                return new Date(y, m - 1, d);
            } else {
                return new Date();
            }
        };
    }
    if ($.fn.datetimebox && $.fn.datebox) {
        $.extend($.fn.datetimebox.defaults, {
            currentText: $.fn.datebox.defaults.currentText,
            closeText: $.fn.datebox.defaults.closeText,
            okText: $.fn.datebox.defaults.okText
        });
    }
    if ($.fn.datetimespinner) {
        $.fn.datetimespinner.defaults.selections = [
            [0, 4],
            [5, 7],
            [8, 10],
            [11, 13],
            [14, 16],
            [17, 19]
        ]
    }
    $.extend($.fn.validatebox.defaults.rules, {
        username: {
            validator: function(value, param) {
                return /^[a-zA-Z][a-zA-Z0-9_]{4,15}$/.test(value);
            },
            message: that.language.usernameVali ? that.language.usernameVali : "Account must begin with a letter, 5-16 characters, alphanumeric underscore!",
        },
        password: {
            validator: function(value, param) {
                return /^[a-zA-Z][a-zA-Z0-9_]{4,15}$/.test(value);
            },
            message: that.language.passwordVali ? that.language.passwordVali : "Password must begin with a letter, 5-16 characters, alphanumeric underscore!",
        },
        passworded: {
            validator: function(value, param) {
                return $(param[0]).val() == value;
            },
            message: that.language.passwordEdVali ? that.language.passwordEdVali : "Entered passwords differ!",
        },
        nickname: {
            validator: function(value, param) {
                return /^[\u4e00-\u9fa5A-Za-z0-9-_]{2,80}$/.test(value);
            },
            message: that.language.nicknameVali ? that.language.nicknameVali : "Nicknames only allow Chinese characters, English letters, numbers, minus and underlines, and 2-80 bit characters",
        },
        remark: {
            validator: function(value, param) {
                return /^[\u4e00-\u9fa5A-Za-z0-9-_,.!，。！ ]{2,255}$/.test(value);
            },
            message: that.language.remarkVali ? that.language.remarkVali : "Remark only letters, numbers, English Chinese characters, minus, comma, period, exclamation and underline, 2-255 characters!",
        },
        id: {
            validator: function(value, param) {
                return /^[0-9]{0,8}$/.test(value);
            },
            message: that.language.idVali ? that.language.idVali : "Id number must be 1-8 bits positive integers!",
        },
        port: {
            validator: function(value, param) {
                return /^[0-9]{0,5}$/.test(value);
            },
            message: that.language.portVali ? that.language.portVali : "The port number must be 1-5 bits positive integers!",
        },
        readTimeout: {
            validator: function(value, param) {
                return /^[0-9]{0,3}$/.test(value);
            },
            message: that.language.readTimeoutVali ? that.language.readTimeoutVali : "The readTimeout must be 1-3 bits positive integers!",
        },
        writeTimeout: {
            validator: function(value, param) {
                return /^[0-9]{0,3}$/.test(value);
            },
            message: that.language.writeTimeoutVali ? that.language.writeTimeoutVali : "The writeTimeout must be 1-3 bits positive integers!",
        },
        sort: {
            validator: function(value, param) {
                return /^[0-9]{0,3}$/.test(value);
            },
            message: that.language.sortVali ? that.language.sortVali : "Sort must be 1-3 positive integers!",
        },
        menuName: {
            validator: function(value, param) {
                return /^[\u4e00-\u9fa5A-Za-z0-9 ]{1,80}$/.test(value);
            },
            message: that.language.menuTestName ? that.language.menuTestName : "Chinese name only allows Chinese characters, English letters and spaces, 1-80 characters!",
        },
        roleName: {
            validator: function(value, param) {
                return /^[\u4e00-\u9fa5A-Za-z0-9 ]{1,80}$/.test(value);
            },
            message: that.language.roleName + that.language.name ? that.language.roleName + that.language.name : "Role name only allows Chinese characters, English letters and spaces, 1-80 characters!",
        },
        english: { // 验证英语
            validator: function(value) {
                return /^[A-Za-z ]{2,80}$/i.test(value);
            },
            message: that.language.menuTestEn ? that.language.menuTestEn : "English name only allows English and spaces,2-80 characters!",
        },
        menuUrl: { // 验证后台地址
            validator: function(value) {
                return /^[A-Za-z/]{2,80}$/i.test(value);
            },
            message: that.language.menuTestUrl ? that.language.menuTestUrl : "Menu URL only allows English and left slashes,2-80 characters!",
        },
        ip: { // 验证IP地址
            validator: function(value) {
                return /\d+\.\d+\.\d+\.\d+/.test(value);
            },
            message: 'IP地址格式不正确'
        },
        ZIP: {
            validator: function(value, param) {
                return /^[0-9]\d{5}$/.test(value);
            },
            message: '邮政编码不存在'
        },
        QQ: {
            validator: function(value, param) {
                return /^[1-9]\d{4,10}$/.test(value);
            },
            message: 'QQ号码不正确'
        },
        mobile: {
            validator: function(value, param) {
                return /^(?:13\d|15\d|18\d)-?\d{5}(\d{3}|\*{3})$/.test(value);
            },
            message: '手机号码不正确'
        },
        tel: {
            validator: function(value, param) {
                return /^(\d{3}-|\d{4}-)?(\d{8}|\d{7})?(-\d{1,6})?$/.test(value);
            },
            message: '电话号码不正确'
        },
        mobileAndTel: {
            validator: function(value, param) {
                return /(^([0\+]\d{2,3})\d{3,4}\-\d{3,8}$)|(^([0\+]\d{2,3})\d{3,4}\d{3,8}$)|(^([0\+]\d{2,3}){0,1}13\d{9}$)|(^\d{3,4}\d{3,8}$)|(^\d{3,4}\-\d{3,8}$)/.test(value);
            },
            message: '请正确输入电话号码'
        },
        number: {
            validator: function(value, param) {
                return /^[0-9]+.?[0-9]*$/.test(value);
            },
            message: '请输入数字'
        },
        money: {
            validator: function(value, param) {
                return (/^(([1-9]\d*)|\d)(\.\d{1,2})?$/).test(value);
            },
            message: '请输入正确的金额'

        },
        mone: {
            validator: function(value, param) {
                return (/^(([1-9]\d*)|\d)(\.\d{1,2})?$/).test(value);
            },
            message: '请输入整数或小数'

        },
        integer: {
            validator: function(value, param) {
                return /^[+]?[1-9]\d*$/.test(value);
            },
            message: '请输入最小为1的整数'
        },
        integ: {
            validator: function(value, param) {
                return /^[+]?[0-9]\d*$/.test(value);
            },
            message: '请输入整数'
        },
        range: {
            validator: function(value, param) {
                if (/^[1-9]\d*$/.test(value)) {
                    return value >= param[0] && value <= param[1]
                } else {
                    return false;
                }
            },
            message: '输入的数字在{0}到{1}之间'
        },
        minLength: {
            validator: function(value, param) {
                return value.length >= param[0]
            },
            message: '至少输入{0}个字'
        },
        maxLength: {
            validator: function(value, param) {
                return value.length <= param[0]
            },
            message: '最多{0}个字'
        },
        //select即选择框的验证
        selectValid: {
            validator: function(value, param) {
                if (value == param[0]) {
                    return false;
                } else {
                    return true;
                }
            },
            message: '请选择'
        },
        idCode: {
            validator: function(value, param) {
                return /(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)/.test(value);
            },
            message: '请输入正确的身份证号'
        },
        loginName: {
            validator: function(value, param) {
                return /^[\u0391-\uFFE5\w]+$/.test(value);
            },
            message: '登录名称只允许汉字、英文字母、数字及下划线。'
        },
        equalTo: {
            validator: function(value, param) {
                return value == $(param[0]).val();
            },
            message: '两次输入的字符不一至'
        },
        englishOrNum: { // 只能输入英文和数字
            validator: function(value) {
                return /^[a-zA-Z0-9_ ]{1,}$/.test(value);
            },
            message: '请输入英文、数字、下划线或者空格'
        },
        xiaoshu: {
            validator: function(value) {
                return /^(([1-9]+)|([0-9]+\.[0-9]{1,2}))$/.test(value);
            },
            message: '最多保留两位小数！'
        },
        ddPrice: {
            validator: function(value, param) {
                if (/^[1-9]\d*$/.test(value)) {
                    return value >= param[0] && value <= param[1];
                } else {
                    return false;
                }
            },
            message: '请输入1到100之间正整数'
        },
        jretailUpperLimit: {
            validator: function(value, param) {
                if (/^[0-9]+([.]{1}[0-9]{1,2})?$/.test(value)) {
                    return parseFloat(value) > parseFloat(param[0]) && parseFloat(value) <= parseFloat(param[1]);
                } else {
                    return false;
                }
            },
            message: '请输入0到100之间的最多俩位小数的数字'
        },
        rateCheck: {
            validator: function(value, param) {
                if (/^[0-9]+([.]{1}[0-9]{1,2})?$/.test(value)) {
                    return parseFloat(value) > parseFloat(param[0]) && parseFloat(value) <= parseFloat(param[1]);
                } else {
                    return false;
                }
            },
            message: '请输入0到1000之间的最多俩位小数的数字'
        }
    });
}
/**
 * @description 随机过渡动画
 * @author Nicholas Mars
 */
HmObj.prototype.noticeAnimateIn = function() {
    var sj = Math.ceil(Math.random() * 46),
        that = this,
        amin = "";
    for (var k in that.animateIn) {
        if (sj == k) {
            amin = that.animateIn[k];
        }
    }
    return amin;
};
/**
 * @description 随机消失动画
 * @author Nicholas Mars
 */
HmObj.prototype.noticeAnimateOut = function() {
    var sj = Math.ceil(Math.random() * 30),
        that = this,
        amout = "";
    for (var k in that.animateOut) {
        if (sj == k) {
            amout = that.animateOut[k];
        }
    }
    return amout;
};
HmObj.prototype.menuCfgReset = function(options) {
    var that = this;
    if (!options.hasOwnProperty('url')) {
        that.notice('error', '没有传入数据源:url参数，菜单项无法正常初始化！');
    }
    var allow = ['url', 'type'];
    var option = that.configFilter(options, allow);
    // 传入参数进行配置
    $.extend(that.menuCfg, options);
    return that;
};
/**
 * config传入参数过滤处理
 * @param  {[type]}
 * @param  {[type]}
 * @return {[type]}
 */
HmObj.prototype.configFilter = function(obj, allow) {
    var newO = {};
    for (var o in obj) {
        if ($.inArray(o, allow)) {
            newO[o] = obj[o];
        }
    }
    return newO;
};
/**
 * @description 弹出一个消息框
 * @author Nicholas Mars
 * @param type  消息框类型，只允许4种，info<warn<error<success，如果没有设置或错误设置，那么默认为info
 * @param msg   消息框的信息,如果不设置自动调用语言文件里相应的内容
 * @param title 消息框的标题,如果不设置自动调用语言文件里相应的内容
 * @param time  消息框持续时间，如果不设置，默认为3秒
 */
HmObj.prototype.notice = function(type, msg, title, time) {
    PNotify.prototype.options.styling = "bootstrap3";
    var icon, that = this,
        amin = that.config.animateRandom ? that.noticeAnimateIn() : false,
        amout = that.config.animateRandom ? that.noticeAnimateOut() : false;
    //判断错误信息类型，更改背景和图标
    switch (type) {
        case "info": //普通消息
            icon = "fa-info-circle";
            break;
        case "warn": //警告
            icon = "fa-exclamation-circle";
            type = "notice";
            break;
        case "error": //错误
            icon = "fa-times-circle";
            break;
        case "success": //成功
            icon = "fa-check-circle";
            break;
        default:
            icon = "fa-info-circle";
            break;
    }
    //由于依赖语言文件加载为先，所以需要延迟0.1秒，否则读取不了语言文件

    new PNotify({
        title: title ? title : that.language.noticeTitle, //标题
        text: msg ? msg : that.language.noticeMsg, //内容
        animate: { //动画效果
            animate: true,
            in_class: amin ? amin : 'bounceInRight',
            out_class: amout ? amout : 'bounceOut'
        },
        // styling: "fontawesome", //选择样式,"brighttheme", "bootstrap3", "fontawesome"
        addclass: "hm-custom", //增加class用以自定义样式
        cornerclass: "hm-custom-content", //增加消息框边框样式
        width: "300px", //宽度
        // min_height: "16px", //最小高度
        icon: 'fa ' + icon, //图标
        type: type ? type : "info", //类型notice,info,success,error
        shadow: true, //阴影
        delay: time ? time : 3000, //多少毫秒后消息被删除
        hide: true, //是否自动关闭
        mouse_reset: true, //鼠标悬浮的时候，时间重置
        nonblock: {
            nonblock: false, //无阻塞消息
        },
    });
    //设置关闭与暂停的高度
    var noticeHeight = $('.hm-custom.ui-pnotify').innerHeight() / 2;
    $('.hm-custom .ui-pnotify-sticker').attr('style', 'cursor: pointer; visibility: visible;height:' + noticeHeight + 'px;line-height:' + (noticeHeight / 2) + 'px');
    $('.hm-custom .ui-pnotify-closer').attr('style', 'cursor: pointer; visibility: visible;height:' + noticeHeight + 'px;line-height:' + (noticeHeight / 2) + 'px');
    //设置关闭与暂停多语言
    $('.hm-custom .ui-pnotify-closer>span').html(that.language.noticeClose ? that.language.noticeClose : "关闭");
    $('.hm-custom .ui-pnotify-sticker>span').html(that.language.noticePause ? that.language.noticePause : "暂停");
    //自定义具体窗口高度的位置
    $('.hm-custom.ui-pnotify').attr('style', 'display:none;top:120px;width:300px;right:16px;');
};
// 获取菜单数据并插入一级菜单
HmObj.prototype.getMenuData = function() {
    var that = this,
        config = that.menuCfg;
    if (config.url === undefined) {
        that.notice('error', hm.language.menuDataUrl ? hm.language.menuDataUrl : 'Please configure the data source for the menu item【Url】');
    } else {
        //若未传入data参数 通过url方式获取
        $.ajax({
            type: config.type,
            url: config.url,
            // async: false,
            dataType: 'json',
            success: function(result, status, xhr) {
                // 取得数据
                if (result == null) {
                    alert('1111');
                    that.notice('error', hm.language.getMenuData ? hm.language.getMenuData : 'Getting menu data failure!');
                } else if (result.status == 66) {
                    that.notice('error', hm.language.noAuthGetMenu ? hm.language.noAuthGetMenu : 'You do not have permission to get the menu!');
                    that.closeProgress(1);
                } else {
                    if (result != null) {
                        window.localStorage.setItem('hm_menu_data', JSON.stringify(result));
                        var str = '',
                            count = result.length,
                            lang = 'text';
                        switch ($.cookie('back-language')) {
                            case 'cn':
                                lang = 'text';
                                break;
                            case 'en':
                                lang = 'en';
                                break;
                            default:
                                lang = 'text';
                                break;
                        }
                        for (var i = 0; i < count; i++) {
                            str += '<a href="#" data-id="' + result[i]["id"] + '" class="easyui-linkbutton oneji" iconCls="fa ' + result[i]["iconCls"] + '" iconAlign="top" plain="true">' + result[i][lang] + '</a>';
                        }
                        $('#hm-north-content').html(str);
                        $('.oneji').linkbutton();
                    } else {
                        that.notice('error', hm.language.localStorageMenu ? hm.language.localStorageMenu : 'No menu data in the local cache');
                        return false;
                    }

                }
            },
            error: function(xhr, status, error) {
                that.notice('error', 'hemacms Error: ' + error);
            },
            complete: function() {
                // _that.init();
            }
        });
    }

};
HmObj.prototype.menuSet = function() {
    var data = JSON.parse(window.localStorage.getItem('hm_menu_data'));

    if (data != null) {
        var str = '',
            count = data.length;
        for (var i = 0; i < count; i++) {
            // console.log(data[i]);
            str += '<a href="#" data-id="' + data[i]["id"] + '" class="easyui-linkbutton oneji" iconCls="fa ' + data[i]["iconCls"] + '" iconAlign="top" plain="true">' + data[i]['text'] + '</a>';
        }
        $('#hm-north-content').html(str);
    } else {
        this.notice('error', hm.language.localStorageMenu ? hm.language.localStorageMenu : 'No menu data in the local cache');
        return false;
    }
};
//打开模态进度条
HmObj.prototype.openProgress = function(msg) {
    var text = this.language.easyuiLoadMsg ? this.language.easyuiLoadMsg : 'Processing, please wait ...';
    var op = '';
    if (msg == null) {
        op = text;
    } else {
        op = msg;
    }
    $.messager.progress({
        text: op,
    });
};
//检测页面是否加载完毕
HmObj.prototype.closeProgress = function(handwork) {
    if (handwork == null) {
        if (document.readyState == "complete") { //当页面加载状态为完全结束时进入
            $.messager.progress('close');
        }
    } else {
        $.messager.progress('close');
    }
};
//刷新当前tab页面
HmObj.prototype.refreshTab = function() {
    var currentTab = $('#hm-tabs').tabs('getSelected');
    var url = $(currentTab.panel('options')).attr('href');
    $('#hm-tabs').tabs('update', {
        tab: currentTab,
        options: {
            href: url
        }
    });
    currentTab.panel('refresh');
}
//刷新当前页面
HmObj.prototype.reloadPage = function() {
    window.location.reload();
};
//数据表格宽度自动适应
HmObj.prototype.datagridWidthResize = function(domid, type) {
    $(window).resize(function() {
        if (type == 'treegrid') {
            $(domid).treegrid({
                fitColumns: true,
                autoRowHeight: true,
            });
        } else {
            $(domid).datagrid({
                fitColumns: true,
                autoRowHeight: true,
            });
        }
    });
};
//数据表格重载数据
HmObj.prototype.datagridReload = function(domid, type) {
    if (type == 'treegrid') {
        $(domid).treegrid("reload");
    } else {
        $(domid).datagrid({
            onLoadSuccess: function(data) {
                if (data.status == 66) {
                    parent.window.location = '/intendant/login';
                } else {
                    $('.easyui-linkbutton').linkbutton();
                    $('.easyui-tooltip').tooltip({
                        position: 'bottom',
                        onShow: function() {
                            $(this).tooltip('tip').css({
                                backgroundColor: '#f39c12',
                                borderColor: '#f39c12',
                                color: '#fff',
                                size: '18px',
                            });
                        }
                    });
                }
            }
        });
    }
};
//快捷按键
HmObj.prototype.shortcutKey = function() {
    if (event.keyCode == 13) {
        document.getElementById("hm_login_submit").click(); //回车登录
    }
    if (event.keyCode == 18) {
        $('#username').val("");
        $('#password').val("");
    }
};
//查询快捷键
HmObj.prototype.searchKey = function() {
    if (event.keyCode == 13) {
        $('#hm_search_submit').click(); //回车登录
    }
}
//点击选择图标
HmObj.prototype.selectMenuIcons = function(icon) {
    $('iframe').contents().find("#hm_menu_icons").html('<i class="' + icon + '"></i>');
    $('iframe').contents().find("#hm_menu_icons_input").val(icon);
    $('.hm_dialog').dialog('destroy');
}
//打开选择图标框
HmObj.prototype.showMenuIcons = function(title) {
    var iconPach = '/intendant/site/iconsCls';
    $('<div class="hm_dialog" style="position:relative"/>').dialog({
        title: '&nbsp;&nbsp;&nbsp;&nbsp;' + title,
        width: 780,
        height: 500,
        closable: false, //去除右上角的X关闭按钮
        draggable: false, //禁止拖动窗口
        cache: false,
        href: '',
        modal: true,
        buttons: [{
            text: this.language.close ? this.language.close : "Close",
            handler: function() {
                $('.hm_dialog').dialog('destroy');
            }
        }]
    });
    $('.hm_dialog').html('正在加载图标中...');
    $.post(iconPach, "", function(data) {
        if (typeof data == 'object') {
            var content = [];
            for (x in data) {
                content[x] = "<a title='点击选择' class=' " + data[x] + "' onclick=\"parent.hm.selectMenuIcons('" + data[x] + "')\"></a>";
            }
            $('.hm_dialog').addClass('hm_icons');
            $('.hm_dialog').html("<div style='padding:15px;'>" + content.join(" ") + "</div>");
        } else {
            $('.layui-layer-content').html("<div style='padding:10px;color: #FFD700;'>图标加载失败，请联系管理员！</div>");
        }
    }, 'json').error(function() {
        $('.layui-layer-content').html("<div style='padding:10px;color: #FFD700;'>图标加载失败，请联系管理员！3秒后自动关闭...</div>");
    });

};
//关闭tab
HmObj.prototype.closeTab = function(menu, type) {
    var currentTabTitle = $(menu).data("tabTitle");
    var currentTab = $('#hm-tabs').tabs('getSelected');
    switch (type) {
        case 1:
            $("#hm-tabs").tabs("close", currentTabTitle);
            return false;
            break;
        case 2:
            $(".tabs li").each(function(index, obj) {
                //获取所有可关闭的选项卡
                var tab = $(".tabs-closable", this).text();
                if (tab !== currentTabTitle) {
                    $("#hm-tabs").tabs('close', tab);
                }
            });
            return false;
            break;
        case 3:
            $(".tabs li").each(function(index, obj) {
                //获取所有可关闭的选项卡
                var tab = $(".tabs-closable", this).text();
                $("#hm-tabs").tabs('close', tab);
            });
            return false;
            break;
        default:
    }
};
//将时间戳转换成时间
HmObj.prototype.dateFormat = function(time) {
    var newDate = new Date();
    newDate.setTime(time * 1000);
    switch ($.cookie('back-language')) {
        case "cn":
            return newDate.toLocaleString();
            break;
        case "en":
            return newDate.toGMTString();
            break;
        default:
            return newDate.toGMTString();
            break;
    }
    // // Wed Jun 18 2014
    // console.log(newDate.toDateString());
    // // Wed, 18 Jun 2014 02:33:24 GMT
    // console.log(newDate.toGMTString());
    // // 2014-06-18T02:33:24.000Z
    // console.log(newDate.toISOString());
    // // 2014-06-18T02:33:24.000Z
    // console.log(newDate.toJSON());
    // // 2014年6月18日
    // console.log(newDate.toLocaleDateString());
    // // 2014年6月18日 上午10:33:24
    // console.log(newDate.toLocaleString());
    // // 上午10:33:24
    // console.log(newDate.toLocaleTimeString());
    // // Wed Jun 18 2014 10:33:24 GMT+0800 (中国标准时间)
    // console.log(newDate.toString());
    // // 10:33:24 GMT+0800 (中国标准时间)
    // console.log(newDate.toTimeString());
    // // Wed, 18 Jun 2014 02:33:24 GMT
    // console.log(newDate.toUTCString());
};
//序列化元素
HmObj.prototype.serializeObject = function(params) { /*将form表单内的元素序列化为对象，扩展Jquery的一个方法*/
    var o = {};
    $.each(params, function(index) {
        if (o[this['name']]) {
            o[this['name']] = o[this['name']] + "," + this['value'];
        } else {
            o[this['name']] = this['value'];
        }
    });
    return o;
};
//异步查询
HmObj.prototype.ajaxSearch = function(externalFunc, me) {
    var that = $(me),
        form = that.parents('form'),
        params = form.serializeArray();
    $('#hm_data').datagrid({
        queryParams: this.serializeObject(params),
        onLoadSuccess: function(data) {
            if (data.status == 66) {
                parent.window.location = '/intendant/login';
            } else {
                $('.easyui-linkbutton').linkbutton();
                $('.easyui-tooltip').tooltip({
                    position: 'bottom',
                    onShow: function() {
                        $(this).tooltip('tip').css({
                            backgroundColor: '#f39c12',
                            borderColor: '#f39c12',
                            color: '#fff',
                            size: '18px',
                        });
                    }
                });
            }
        }
    });
}
//清空查询
HmObj.prototype.clearSearch = function(me) {
    var that = $(me);
    // $("#hm_data").datagrid("load", {}); //重新加载数据，无填写数据，向后台传递值则为空
    $('#hm_data').datagrid({
        onLoadSuccess: function(data) {
            if (data.status == 66) {
                parent.window.location = '/intendant/login';
            } else {
                $('.easyui-linkbutton').linkbutton();
                $('.easyui-tooltip').tooltip({
                    position: 'bottom',
                    onShow: function() {
                        $(this).tooltip('tip').css({
                            backgroundColor: '#f39c12',
                            borderColor: '#f39c12',
                            color: '#fff',
                            size: '18px',
                        });
                    }
                });
            }
        }
    });
    that.parents("form[name=datagird_tools]").find("input").val(""); //找到form表单下的所有input标签并清空
};
//批量排序
HmObj.prototype.ajaxSort = function(externalFunc, me) {
    var that = $(me),
        form = that.parents('form'),
        vali = form.form('enableValidation').form('validate');
    if (that.linkbutton('options').disabled == true) {
        parent.hm.notice('warn', hm.language.easyuiResubmit ? hm.language.easyuiResubmit : 'Don‘t repeat the submit');
    } else {
        that.linkbutton('disable');
        if (vali) {
            var param = {};
            var sorts = form.serializeArray();
            $.each(sorts, function() {
                if (param[this.name] !== undefined) {
                    if (!param[this.name].push) {
                        param[this.name] = [param[this.name]];
                    }
                    param[this.name].push(this.value || '');
                } else {
                    param[this.name] = this.value || '';
                }
            });
            $.ajax({
                url: form.attr('action'),
                dataType: 'json',
                type: 'POST',
                data: JSON.stringify(param),
                beforeSend: function() {
                    parent.hm.openProgress();
                },
                success: function(data) {
                    parent.hm.closeProgress();
                    externalFunc(data)
                    that.linkbutton('enable');
                },
                error: function(XMLHttpRequest, textStatus, errorThrown) {
                    parent.hm.notice("error", 'status:' + XMLHttpRequest.status + ',readyState:' + XMLHttpRequest.readyState + ',textStatus:' + textStatus);
                    that.linkbutton('enable');
                    parent.hm.closeProgress();
                }
            });
        } else {
            parent.hm.notice('warn', hm.language.easyuiRemoteMessage ? hm.language.easyuiRemoteMessage : 'Please fix this field.');
            setTimeout(function() {
                that.linkbutton('enable');
            }, 2000);
        }
    }
};
//修改propertygrid
HmObj.prototype.ajaxPgrid = function(externalFunc, me) {
    var that = $(me),
        form = that.parents('form'),
        vali = form.form('enableValidation').form('validate'),
        failed = hm.language.editFailed ? hm.language.editFailed : "Edit failed!",
        success = hm.language.editSuccess ? hm.language.editSuccess : "Edit success!",
        arr = $('#hm-pgrid').propertygrid('getRows'),
        s = '',
        data = {};
    for (var i = 0; i < arr.length; i++) {
        s += arr[i].key + '=' + encodeURIComponent(arr[i].value); //url键值对
        data[arr[i].key] = arr[i].value; //JSON对象
    }
    if (that.linkbutton('options').disabled == true) {
        parent.hm.notice('warn', hm.language.easyuiResubmit ? hm.language.easyuiResubmit : 'Don‘t repeat the submit');
    } else {
        that.linkbutton('disable');
        if (vali) {
            $.ajax({
                url: form.attr('action'),
                dataType: 'json',
                type: 'POST',
                data: data,
                beforeSend: function() {
                    parent.hm.openProgress();
                },
                success: function(data) {
                    parent.hm.closeProgress(); //当页面加载完毕执行
                    externalFunc(data)
                    that.linkbutton('enable');
                },
                error: function(XMLHttpRequest, textStatus, errorThrown) {
                    parent.hm.notice("error", 'status:' + XMLHttpRequest.status + ',readyState:' + XMLHttpRequest.readyState + ',textStatus:' + textStatus);
                    that.linkbutton('enable');
                    parent.hm.closeProgress();
                }
            });
        } else {
            parent.hm.notice('warn', hm.language.easyuiRemoteMessage ? hm.language.easyuiRemoteMessage : 'Please fix this field.');
            setTimeout(function() {
                that.linkbutton('enable');
            }, 2000);
        }
    }

};
//获取tree节点ids
HmObj.prototype.getTreeNodes = function(externalFunc, me) {
    var that = $(me),
        url = that.attr("data-url");
    if (that.linkbutton('options').disabled == true) {
        parent.hm.notice('warn', hm.language.easyuiResubmit ? hm.language.easyuiResubmit : 'Don‘t repeat the submit');
    } else {
        that.linkbutton('disable');
        var nodes = $('#hm_tree').tree('getChecked', ['checked', 'indeterminate']),
            ids = '',
            id = that.attr("data-id");
        for (var i = 0; i < nodes.length; i++) {
            if (ids != '') ids += ',';
            ids += nodes[i].id;
        }
        $.ajax({
            url: url,
            dataType: 'json',
            type: 'POST',
            data: {
                'ids': ids,
                'id': id
            },
            beforeSend: function() {
                parent.hm.openProgress();
            },
            success: function(data) {
                parent.hm.closeProgress();
                externalFunc(data)
                that.linkbutton('enable');
            },
            error: function(XMLHttpRequest, textStatus, errorThrown) {
                parent.hm.notice("error", 'status:' + XMLHttpRequest.status + ',readyState:' + XMLHttpRequest.readyState + ',textStatus:' + textStatus);
                that.linkbutton('enable');
                parent.hm.closeProgress();
            }
        });

    }
};
//添加或修改
HmObj.prototype.ajaxAddEdit = function(externalFunc, me) {
    var that = $(me),
        form = that.parents('form'),
        vali = form.form('enableValidation').form('validate');
    console.info(form.serialize());
    if (that.linkbutton('options').disabled == true) {
        parent.hm.notice('warn', hm.language.easyuiResubmit ? hm.language.easyuiResubmit : 'Don‘t repeat the submit');
    } else {
        that.linkbutton('disable');
        if (vali) {
            $.ajax({
                url: form.attr('action'),
                dataType: 'json',
                type: 'POST',
                data: form.serialize(),
                beforeSend: function() {
                    parent.hm.openProgress();
                },
                success: function(data) {
                    parent.hm.closeProgress();
                    externalFunc(data)
                    that.linkbutton('enable');
                },
                error: function(XMLHttpRequest, textStatus, errorThrown) {
                    parent.hm.notice("error", 'status:' + XMLHttpRequest.status + ',readyState:' + XMLHttpRequest.readyState + ',textStatus:' + textStatus);
                    that.linkbutton('enable');
                    parent.hm.closeProgress();
                }
            });
        } else {
            parent.hm.notice('warn', hm.language.easyuiRemoteMessage ? hm.language.easyuiRemoteMessage : 'Please fix this field.');
            setTimeout(function() {
                that.linkbutton('enable');
            }, 2000);
        }
    }
};
//根据id删除单个
HmObj.prototype.ajaxDel = function(externalFunc, me) {
    var that = $(me),
        url = that.attr('data-url'),
        id = that.attr('data-id'),
        title = hm.language.noticeTitle ? hm.language.noticeTitle : "Notice Message",
        failed = hm.language.delFailed ? hm.language.delFailed : "Delete failed!",
        success = hm.language.delSuccess ? hm.language.delSuccess : "Delete success!",
        msg = hm.language.confirmDel ? hm.language.confirmDel : "You confirm that you want to delete ";
    parent.$.messager.confirm({
        title: '&nbsp;&nbsp;&nbsp;&nbsp;' + title,
        msg: msg + 'ID&nbsp;<span style="color:red;">[' + id + ']</span>&nbsp;?',
        closable: false, //去除右上角的X关闭按钮
        draggable: false, //禁止拖动窗口
        fn: function(r) {
            if (r) {
                $.ajax({
                    url: url,
                    dataType: 'json',
                    type: 'POST',
                    data: {
                        id: id
                    },
                    beforeSend: function() {
                        parent.hm.openProgress();
                    },
                    success: function(data) {
                        parent.hm.closeProgress();
                        externalFunc(data)
                        that.linkbutton('enable');
                    },
                    error: function(XMLHttpRequest, textStatus, errorThrown) {
                        parent.hm.notice("error", 'status:' + XMLHttpRequest.status + ',readyState:' + XMLHttpRequest.readyState + ',textStatus:' + textStatus);
                        that.linkbutton('enable');
                        parent.hm.closeProgress();
                    }
                });
            }
        }
    });
};
//根据id批量删除
HmObj.prototype.ajaxBatchDel = function(externalFunc, me) {
    var that = $(me),
        title = hm.language.noticeTitle ? hm.language.noticeTitle : "Notice Message",
        msg = hm.language.confirmDel ? hm.language.confirmDel : "You confirm that you want to delete ";
    if (that.linkbutton('options').disabled == true) {
        parent.hm.notice('warn', hm.language.easyuiResubmit ? hm.language.easyuiResubmit : 'Don‘t repeat the submit');
    } else {
        that.linkbutton('disable');
        var checkedItems = $('#hm_data').datagrid('getChecked'),
            param = [],
            name = [];
        $.each(checkedItems, function(k, v) {
            param.push(v.id);
        });
        var ids = param.join(",");
        if (param.length == 0) {
            hm.notice('warn', "你没有选择任何");
            that.linkbutton('enable');
            return false;
        }
        parent.$.messager.confirm({
            title: '&nbsp;&nbsp;&nbsp;&nbsp;' + title,
            msg: msg + 'ID&nbsp;<span style="color:red;">[' + ids + ']</span>&nbsp;?',
            closable: false, //去除右上角的X关闭按钮
            draggable: false, //禁止拖动窗口
            fn: function(r) {
                if (r) {
                    $.ajax({
                        url: that.attr('action'),
                        dataType: 'json',
                        type: 'POST',
                        data: {
                            'ids': ids
                        },
                        beforeSend: function() {
                            parent.hm.openProgress();
                        },
                        success: function(data) {
                            parent.hm.closeProgress();
                            externalFunc(data)
                            that.linkbutton('enable');
                        },
                        error: function(XMLHttpRequest, textStatus, errorThrown) {
                            parent.hm.notice("error", 'status:' + XMLHttpRequest.status + ',readyState:' + XMLHttpRequest.readyState + ',textStatus:' + textStatus);
                            that.linkbutton('enable');
                            parent.hm.closeProgress();
                        }
                    });
                } else {
                    that.linkbutton('enable');
                }
            }
        });
    }
};
//创建HmMenuTab对象
var hm = new HmObj();
//如果后台没有设置语言，JS来设置
hm.languageSet('en');
hm.languageRun();
//启动国际化
hm.readJsonFile();
//启动easui国际化
hm.easyuiLanguage();
$(function() {
    //排序
    // $('.ajax_sort').on('click', function(e) {
    //     e.preventDefault();
    //     var that = $(this),
    //         form = that.parents('form[name=hm_form_sort]'),
    //         vali = form.form('enableValidation').form('validate');
    //     if (that.linkbutton('options').disabled == true) {
    //         parent.hm.notice('warn', hm.language.easyuiResubmit ? hm.language.easyuiResubmit : 'Don‘t repeat the submit');
    //     } else {
    //         that.linkbutton('disable');
    //         if (vali) {
    //             var param = {};
    //             var sorts = form.serializeArray();
    //             $.each(sorts, function() {
    //                 if (param[this.name] !== undefined) {
    //                     if (!param[this.name].push) {
    //                         param[this.name] = [param[this.name]];
    //                     }
    //                     param[this.name].push(this.value || '');
    //                 } else {
    //                     param[this.name] = this.value || '';
    //                 }
    //             });
    //             $.ajax({
    //                 url: form.attr('action'),
    //                 dataType: 'json',
    //                 type: 'POST',
    //                 data: JSON.stringify(param),
    //                 beforeSend: function() {
    //                     parent.hm.openProgress();
    //                 },
    //                 success: function(data) {
    //                     if (!data.status) {
    //                         parent.hm.notice("error", hm.language.menuSortInfoFailed + data.info);
    //                     } else {
    //                         parent.hm.notice("success", hm.language.menuSortInfoSuccess + data.info);
    //                         parent.hm.refreshTab();
    //                     }
    //                     that.linkbutton('enable');
    //                     parent.hm.closeProgress(); //当页面加载完毕执行
    //                 },
    //                 error: function(XMLHttpRequest, textStatus, errorThrown) {
    //                     parent.hm.notice("error", 'status:' + XMLHttpRequest.status + ',readyState:' + XMLHttpRequest.readyState + ',textStatus:' + textStatus);
    //                     that.linkbutton('enable');
    //                     parent.hm.closeProgress(); //当页面加载完毕执行
    //                 }
    //             });
    //         } else {
    //             parent.hm.notice('warn', hm.language.easyuiRemoteMessage ? hm.language.easyuiRemoteMessage : 'Please fix this field.');
    //             setTimeout(function() {
    //                 that.linkbutton('enable');
    //             }, 2000);
    //         }
    //     }
    // });
    //添加或修改
    // $('.ajax_add_edit').on('click', function(e) {
    //     e.preventDefault();
    //     var that = $(this),
    //         form = that.parents('form[name=hm_form_add_edit]'),
    //         vali = form.form('enableValidation').form('validate');
    //     switch (form.attr('type')) {
    //         case "add":
    //             var failed = hm.language.addFailed ? hm.language.addFailed : "Add failed!",
    //                 success = hm.language.addSuccess ? hm.language.addSuccess : "Add success!";
    //             break;
    //         case "edit":
    //             var failed = hm.language.editFailed ? hm.language.editFailed : "Edit failed!",
    //                 success = hm.language.editSuccess ? hm.language.editSuccess : "Edit success!";
    //             break;
    //         default:
    //             hm.notice('warn', hm.language.unableAddEditType ? hm.language.unableAddEditType : "Parameter error,Unable to add or edit!");
    //     }
    //     if (that.linkbutton('options').disabled == true) {
    //         parent.hm.notice('warn', hm.language.easyuiResubmit ? hm.language.easyuiResubmit : 'Don‘t repeat the submit');
    //     } else {
    //         that.linkbutton('disable');
    //         if (vali) {
    //             $.ajax({
    //                 url: form.attr('action'),
    //                 dataType: 'json',
    //                 type: 'POST',
    //                 data: form.serialize(),
    //                 beforeSend: function() {
    //                     parent.hm.openProgress();
    //                 },
    //                 success: function(data) {
    //                     if (!data.status) {
    //                         parent.hm.notice("error", failed + data.info);
    //                     } else {
    //                         parent.hm.notice("success", success + data.info);
    //                         parent.hm.refreshTab();
    //                     }
    //                     that.linkbutton('enable');
    //                     parent.hm.closeProgress(); //当页面加载完毕执行
    //                 },
    //                 error: function(XMLHttpRequest, textStatus, errorThrown) {
    //                     parent.hm.notice("error", 'status:' + XMLHttpRequest.status + ',readyState:' + XMLHttpRequest.readyState + ',textStatus:' + textStatus);
    //                     that.linkbutton('enable');
    //                     parent.hm.closeProgress(); //当页面加载完毕执行
    //                 }
    //             });
    //         } else {
    //             parent.hm.notice('warn', hm.language.easyuiRemoteMessage ? hm.language.easyuiRemoteMessage : 'Please fix this field.');
    //             setTimeout(function() {
    //                 that.linkbutton('enable');
    //             }, 2000);
    //         }
    //     }
    // });
    //删除单个
    $('form').on('click', '.ajax_del', function(e) {
        e.preventDefault();
        var that = $(this),
            name = that.attr('data-name'),
            id = that.attr('data-id'),
            title = hm.language.noticeTitle ? hm.language.noticeTitle : "Notice Message",
            failed = hm.language.delFailed ? hm.language.delFailed : "Delete failed!",
            success = hm.language.delSuccess ? hm.language.delSuccess : "Delete success!";
        msg = hm.language.confirmDel ? hm.language.confirmDel : "You confirm that you want to delete ";
        parent.$.messager.confirm({
            title: '&nbsp;&nbsp;&nbsp;&nbsp;' + title,
            msg: msg + '[' + name + ']?',
            closable: false, //去除右上角的X关闭按钮
            draggable: false, //禁止拖动窗口
            fn: function(r) {
                if (r) {
                    $.ajax({
                        url: that.attr('href'),
                        dataType: 'json',
                        type: 'POST',
                        data: {
                            id: id
                        },
                        beforeSend: function() {
                            parent.hm.openProgress();
                        },
                        success: function(data) {
                            if (!data.status) {
                                parent.hm.notice("error", failed + data.info);
                            } else {
                                parent.hm.notice("success", success + data.info);
                                parent.hm.refreshTab();
                            }
                            that.linkbutton('enable');
                            parent.hm.closeProgress(); //当页面加载完毕执行
                        },
                        error: function(XMLHttpRequest, textStatus, errorThrown) {
                            parent.hm.notice("error", 'status:' + XMLHttpRequest.status + ',readyState:' + XMLHttpRequest.readyState + ',textStatus:' + textStatus);
                            that.linkbutton('enable');
                            parent.hm.closeProgress(); //当页面加载完毕执行
                        }
                    });
                }
            }
        });
    });
    //propertygrid编辑
    // $('.ajax_pgrid').on('click', function(e) {
    //     e.preventDefault();
    //     var that = $(this),
    //         form = that.parents('form[name=hm_form_pgrid]'),
    //         vali = form.form('enableValidation').form('validate'),
    //         failed = hm.language.editFailed ? hm.language.editFailed : "Edit failed!",
    //         success = hm.language.editSuccess ? hm.language.editSuccess : "Edit success!",
    //         arr = $('#hm-pgrid').propertygrid('getRows'),
    //         s = '',
    //         data = {};
    //     for (var i = 0; i < arr.length; i++) {
    //         s += arr[i].key + '=' + encodeURIComponent(arr[i].value); //url键值对
    //         data[arr[i].key] = arr[i].value; //JSON对象
    //     }
    //     if (that.linkbutton('options').disabled == true) {
    //         parent.hm.notice('warn', hm.language.easyuiResubmit ? hm.language.easyuiResubmit : 'Don‘t repeat the submit');
    //     } else {
    //         that.linkbutton('disable');
    //         if (vali) {
    //             $.ajax({
    //                 url: form.attr('action'),
    //                 dataType: 'json',
    //                 type: 'POST',
    //                 data: data,
    //                 beforeSend: function() {
    //                     parent.hm.openProgress();
    //                 },
    //                 success: function(data) {
    //                     if (data.status !== 1) {
    //                         that.linkbutton('enable');
    //                         parent.hm.closeProgress(); //当页面加载完毕执行
    //                         parent.hm.notice("error", failed + data.info);
    //                     } else {
    //                         that.linkbutton('enable');
    //                         parent.hm.closeProgress(); //当页面加载完毕执行
    //                         parent.hm.notice("success", success + data.info);
    //                         parent.hm.refreshTab();
    //                     }
    //                 },
    //                 error: function(XMLHttpRequest, textStatus, errorThrown) {
    //                     parent.hm.notice("error", 'status:' + XMLHttpRequest.status + ',readyState:' + XMLHttpRequest.readyState + ',textStatus:' + textStatus);
    //                     that.linkbutton('enable');
    //                     parent.hm.closeProgress(); //当页面加载完毕执行
    //                 }
    //             });
    //         } else {
    //             parent.hm.notice('warn', hm.language.easyuiRemoteMessage ? hm.language.easyuiRemoteMessage : 'Please fix this field.');
    //             setTimeout(function() {
    //                 that.linkbutton('enable');
    //             }, 2000);
    //         }
    //     }
    // });
    //批量根据id删除
    // $('.ajax_batch_del').on('click', function(e) {
    //     e.preventDefault();
    //     var that = $(this),
    //         title = hm.language.noticeTitle ? hm.language.noticeTitle : "Notice Message",
    //         msg = hm.language.confirmDel ? hm.language.confirmDel : "You confirm that you want to delete ";
    //     if (that.linkbutton('options').disabled == true) {
    //         parent.hm.notice('warn', hm.language.easyuiResubmit ? hm.language.easyuiResubmit : 'Don‘t repeat the submit');
    //     } else {
    //         that.linkbutton('disable');
    //         var checkedItems = $('#hm_data').datagrid('getChecked'),
    //             param = [],
    //             name = [];
    //         $.each(checkedItems, function(k, v) {
    //             param.push(v.id);
    //             name.push(v.username);
    //         });
    //         if (param.length == 0) {
    //             hm.notice('warn', "你没有选择任何");
    //             that.linkbutton('enable');
    //             return false;
    //         }
    //         parent.$.messager.confirm({
    //             title: '&nbsp;&nbsp;&nbsp;&nbsp;' + title,
    //             msg: msg + '[' + name.join(",") + ']?',
    //             closable: false, //去除右上角的X关闭按钮
    //             draggable: false, //禁止拖动窗口
    //             fn: function(r) {
    //                 if (r) {
    //                     $.ajax({
    //                         url: that.attr('action'),
    //                         dataType: 'json',
    //                         type: 'POST',
    //                         data: {
    //                             'id': param.join(",")
    //                         },
    //                         beforeSend: function() {
    //                             parent.hm.openProgress();
    //                         },
    //                         success: function(data) {
    //                             if (!data.status) {
    //                                 parent.hm.notice("error", hm.language.menuSortInfoFailed + data.info);
    //                             } else {
    //                                 parent.hm.notice("success", hm.language.menuSortInfoSuccess + data.info);
    //                                 parent.hm.refreshTab();
    //                             }
    //                             that.linkbutton('enable');
    //                             parent.hm.closeProgress(); //当页面加载完毕执行
    //                         },
    //                         error: function(XMLHttpRequest, textStatus, errorThrown) {
    //                             parent.hm.notice("error", 'status:' + XMLHttpRequest.status + ',readyState:' + XMLHttpRequest.readyState + ',textStatus:' + textStatus);
    //                             that.linkbutton('enable');
    //                             parent.hm.closeProgress(); //当页面加载完毕执行
    //                         }
    //                     });
    //                 } else {
    //                     that.linkbutton('enable');
    //                 }
    //             }
    //         });
    //     }
    // });
});