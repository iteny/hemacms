var language = {
    //英语,english
    en: {
        //easyui国际化开始
        easyuiBeforePageText: "Page",
        easyuiAfterPageText: "of {pages}",
        easyuiDisplayMsg: "Displaying {from} to {to} of {total} items",
        easyuiLoadMsg: "Processing, please wait ...",
        easyuiOk: "Ok",
        easyuiCancel: "Cancel",
        easyuiMissingMessage: "This field is required.",
        easyuiEmailMessage: "Please enter a valid email address.",
        easyuiUrlMessage: "Please enter a valid URL.",
        easyuiLengthMessage: "Please enter a value between {0} and {1}.",
        easyuiRemoteMessage: "Please fix this field.",
        easyuiWeeks: ['S', 'M', 'T', 'W', 'T', 'F', 'S'],
        easyuiMonths: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
        easyuiCurrentText: "Today",
        easyuiCloseText: "Close",
        easyuiOkText: "Ok",
        easyuiResubmit: "Don't repeat the submit",
        //easyui国际化结束
        //通知中心开始
        parsleyEmail: "This value should be a valid email.",
        parsleyUrl: "This value should be a valid url.",
        parsleyNumber: "This value should be a valid number.",
        parsleyInteger: "This value should be a valid integer.",
        parsleyDigits: "This value should be digits.",
        parsleyAlphanum: "This value should be alphanumeric.",
        parsleyNotblank: "This value should not be blank.",
        parsleyRequired: "This value is required.",
        parsleyPattern: "This value seems to be invalid.",
        parsleyMin: "This value should be greater than or equal to %s.",
        parsleyMax: "This value should be lower than or equal to %s.",
        parsleyRange: "This value should be between %s and %s.",
        parsleyMinlength: "This value is too short. It should have %s characters or more.",
        parsleyMaxlength: "This value is too long. It should have %s characters or fewer.",
        parsleyLength: "This value length is invalid. It should be between %s and %s characters long.",
        parsleyMincheck: "You must select at least %s choices.",
        parsleyMaxcheck: "You must select %s choices or fewer.",
        parsleyCheck: "You must select between %s and %s choices.",
        parsleyEqualto: "This value should be the same.",
        noticeTitle: "Notice Message", //notice title default english
        noticeMsg: "Sorry,I didn't receive the message!", //notice message default english
        noticeClose: "Close", //notice close button default english
        noticePause: "Pause", //notice close button default english
        noticeLast: "last one notice",
        noticeAll: "last all notice",
        confirmDel: "You confirm that you want to delete ",
        close: "Close",
        //通知中心结束
        sorry: "Sorry!", //sorry
        homePage: "Home",
        noShutdown: "no shut down",
        sort: "Sort",
        superiorMenu: "Superior menu",
        menuIcon: "Menu icon",
        menuSort: "Menu sort",
        menuName: "Menu name",
        menuEnglish: "English name",
        menuChinese: "Chinese name",
        menuUrl: "Menu url",
        isshow: "Isshow",
        remark: "Remark",
        seeRemark: "See remark",
        noRemark: "No remark",
        noData: "no records data found",
        menuEn: "Menu english",
        manageOperat: "Management operation",
        menuLevelConfine: "Sorry,Do not allow the menu more than three!",
        menuSortInfoSuccess: "Menu sort successed!",
        menuSortInfoFailed: "Menu sort failed!",
        unableAddEditType: "Parameter error,Unable to add or edit!",
        addSuccess: "Add success!",
        addFailed: "Add failed!",
        editSuccess: "Edit success!",
        editFailed: "Edit failed!",
        delSuccess: "Delete success!",
        delFailed: "Delete failed!",
        searchSuccess: "Search success!",
        searchFailed: "Search failed!",
        authSuccess: "Authorization success!",
        authFailed: "Authorization failed!",
        success: "success",
        fail: "failed",
        refresh: "Refresh", //refresh
        refreshCurrent: "Refresh Current Page",
        ajaxAdd: "Add",
        ajaxEdit: "Edit",
        ajaxDel: "Delete",
        ajaxAuth: "Authorization",
        ajaxSetAuth: "Permission setting",
        search: "Search",
        clearSearch: "Clear search",
        goback: "Go back",
        selectIcons: "Select icons",
        topMenu: "Top menu",
        addTopMenu: "Add topmenu",
        addSubMenu: "Add submenu",
        sop: "Sop", //Standard Operating Procelure
        closeCurrentTab: "closed current tab",
        closeElseTab: "closed else tabs",
        closeAllTab: "closed all tabs",
        isMenu: "no menus", //no menus
        localStorageMenu: "No menu data in the local cache",
        menuPid: "Unable to obtain the specified parent ID",
        getMenuData: "Getting menu data failure",
        loadMenuData: "Loading menu data!",
        menuDataUrl: "Please configure the data source for the menu item【Url】",
        on: "on",
        off: "off",
        reloadDatagrid: "Reload data",
        //切换语言开始
        currentLanguage: "Language",
        styleSwitch: "Switch to",
        styleDefault: "Default Style",
        styleMac: "Mac Style",
        //切换语言结束
        //登录开始
        backStageLogin: "Back-stage management",
        login: "Log in",
        username: "Account",
        password: "Password",
        loginMsg: "Waiting...,Log in for you!",
        loginSuccess: "Login success, 3 seconds after you jump!",
        loginPassFail: "Account or password error!",
        loginDisabled: "This account has been disabled!",
        //登录结束
        //用户操作开始
        userLogout: "Log out",
        logoutMsg: "Waiting...,Log out for you!",
        logoutMsgFail: "Get out of Failed!",
        logoutMsgSuccess: "Get out of success!",
        unknownError: "Unknow error!",
        //用户操作结束
        //easyui validatebox扩展开始
        isUsername: "Please enter the account!",
        isPassword: "Please enter the password!",
        isNickname: "Please enter the nickname!",
        isEmail: "Please enter the email!",
        isRemark: "Please enter the remark!",
        usernameVali: "Account must begin with a letter, 5-16 bits characters, alphanumeric underscore!",
        passwordVali: "Password must begin with a letter, 5-16 bits characters, alphanumeric underscore!",
        passwordEdVali: "Entered passwords differ!",
        nicknameVali: "Nicknames only allow Chinese characters, English letters, numbers, minus and underlines, and 2-80 bit characters",
        remarkVali: "Remark only letters, numbers, English Chinese characters, minus, comma, period, exclamation and underline, 2-255 characters!",
        sqlTypeVali: "The database type can only be sqlite3, mysql!",
        portVali: "The port number must be 1-5 bits positive integers!",
        readTimeoutVali: "The readTimeout must be 1-3 bits positive integers!",
        writeTimeoutVali: "The writeTimeout must be 1-3 bits positive integers!",
        sortVali: "Sort must be 1-3 positive integers!",
        idVali: "Id must be 1-8 positive integers!",
        englishVali: "English name only allows English and spaces,2-80 characters!",
        //easyui validatebox扩展结束
        //配置文件开始
        cfgRows: "Rows",
        cfgName: "Config name",
        cfgValue: "Config value",
        cfgRemark: "Config remark",
        cfgSql: "Database configuration",
        cfgServer: "Server configuration",
        cfgUseWhatSql: "Which database to use",
        cfgPort: "Port number",
        cfgReadTimeout: "Read out of time",
        cfgWriteTimeout: "Write out of time",
        cfgCookie: "Cookie configuration",
        cfgLanguage: "Default language",
        cfgLanguageVali: "Default language only allows English and spaces,2-80 characters!",
        cfgCommon: "Common configuration",
        cfgTerminalLog: "Terminal print operation log switch",
        cfgSqlLog: "Database operation log switch",
        cfgAjaxPolling: "Get system occupancy",
        //配置文件结束
        //用户管理开始
        userAdd: "Add account",
        userAdminIsNo: "Superadmin isn't allowed to edit",
        userBatchDel: "Batch delete accounts",
        userName: "Account",
        userRoleName: "Role name",
        userCreateTime: "Create time",
        userCreateIp: "Create ip",
        userEmail: "Email",
        userPwd: "Password",
        userPwdEd: "Confirm password",
        userNickname: "Nickname",
        userRoleGroup: "Endow role",
        userStatus: "Enabled",
        userIsRole: "Role number must be 1 to 8 positive integers!",
        userIsStatus: "Enabled, it must be 1 or 0!",
        userCheckUsername: "Account has already existed!",
        userCheckNickname: "Nickname has already existed!",
        userCheckEmail: "Email has already existed!",
        userId: "Account ID is illegal!",
        userIdGroup: "Account ID array is illegal!",
        userAdminNoDel: "Superadmin does not allow deletion!",
        userSearch: "Search accounts",
        userTo: "to",
        //用户管理结束
        //后台菜单开始
        menuCheckName: "Chinese name already exists!",
        menuCheckUrl: "Menu URL already exists!",
        menuCheckEn: "English name already exists!",
        menuTestPid: "The upper menu ID input is illegal!",
        menuTestName: "Chinese name only allows Chinese characters, English letters and spaces, 2-80 characters!",
        menuTestSort: "Menu ordering is illegal!",
        menuTestIcon: "Menu icon entry is illegal!",
        menuTestIsshow: "Menu isshow input is illegal!",
        menuTestUrl: "Menu URL only allows English and left slashes,2-80 characters!",
        //后台菜单结束
        //角色管理开始
        roleName: "Role name",
        roleChinese: "Role chinese name",
        roleEnglish: "Role english name",
        roleSort: "Role sort",
        roleRemark: "Role remark",
        roleAdd: "Add role",
        roleBatchDel: "Batch delete roles",
        roleSearch: "Search role name",
        roleIds: "Authorization id group is illegal!",
        roleCheckName: "Role chinese name already exists!",
        roleCheckEn: "Role english name already exists!",
        //角色管理结束
        //通用提示开始
        name: " only allows Chinese characters, English letters and spaces, 1-80 characters!",
        noAuth: "You have no authority!",
        noFound: "Sorry! The page you visited cannot be found...",
        english: "English",
        chinese: "Chinese",
        noAuthGetMenu: "You do not have permission to get the menu!",
        noUrl: "The address of the web page was not obtained!",
        tantoTabs: "Sorry! You open too many tabs and it will cause the machine to slow down!",
        currentRole: "Current role",
        jsCopyCss: "Copy color style ",
        jsCopyRgb: "Copy color rgb code ",
        jsCopySuccess: " success!",
        jsCopyFailed: "Copy failed!",
        jsCss: "Copy css",
        jsRgb: "Copy rgb",
        status: "Status",
        backstageHome: "Backstage-HemaCms",
        //通用提示结束
        //日志开始
        loginLogTime: "Login time",
        loginLogIp: "Login ip",
        loginUseragent: "Useragent",
        loginSearchUser: "Search account",
        loginDelMonthLog: "Delete the login log a month ago",
        oprateDetail: "Oprate detail",
        oprateUrl: "Request url",
        oprateInfo: "Oprate info",
        oprateMethod: "Method",
        oprateLogTime: "Oprate time",
        oprateLogIp: "Oprate ip",
        oprateExcuteTime: "Excute time",
        oprateMillisecond: "Millisecond",
        oprateDelMonthLog: "Delete the oprate log a month ago",
        oprateDelAllLog: "Delete all operat logs",
        oprateTime: "Oprate time",
        //日志结束
        //验证数据开始
        valiEnglish: "Please enter a valid English letter A~Z,a~z!",
        //验证数据结束
    },
    //简体中文,chinese
    cn: {
        //easyui国际化开始
        easyuiBeforePageText: "第",
        easyuiAfterPageText: "共{pages}页",
        easyuiDisplayMsg: "显示{from}到{to},共{total}记录",
        easyuiLoadMsg: "正在处理，请稍待。。。",
        easyuiOk: "确定",
        easyuiCancel: "取消",
        easyuiMissingMessage: "该输入项为必输项",
        easyuiEmailMessage: "请输入有效的电子邮件地址",
        easyuiUrlMessage: "请输入有效的URL地址",
        easyuiLengthMessage: "输入内容长度必须介于{0}和{1}之间",
        easyuiRemoteMessage: "请修正该字段",
        easyuiWeeks: ['日', '一', '二', '三', '四', '五', '六'],
        easyuiMonths: ['一月', '二月', '三月', '四月', '五月', '六月', '七月', '八月', '九月', '十月', '十一月', '十二月'],
        easyuiCurrentText: "今天",
        easyuiCloseText: "关闭",
        easyuiOkText: "确定",
        easyuiResubmit: "不要重复提交",
        //easyui国际化结束
        //通知中心开始
        parsleyEmail: "请输入一个有效的电子邮箱地址.",
        parsleyUrl: "请输入一个有效的链接.",
        parsleyNumber: "请输入正确的数字.",
        parsleyInteger: "请输入正确的整数.",
        parsleyDigits: "请输入正确的号码",
        parsleyAlphanum: "请输入字母或数字.",
        parsleyNotblank: "请输入值.",
        parsleyRequired: "必填项.",
        parsleyPattern: "格式不正确.",
        parsleyMin: "输入值请大于或等于 %s.",
        parsleyMax: "输入值请小于或等于 %s.",
        parsleyRange: "输入值应该在 %s 到 %s 之间.",
        parsleyMinlength: "请输入至少 %s 个字符.",
        parsleyMaxlength: "请输入至多 %s 个字符.",
        parsleyLength: "字符长度应该在 %s 到 %s 之间.",
        parsleyMincheck: "请至少选择 %s 个选项.",
        parsleyMaxcheck: "请选择不超过 %s 个选项.",
        parsleyCheck: "请选择 %s 到 %s 个选项.",
        parsleyEqualto: "输入值不同.",
        noticeTitle: "提示信息", //消息框标题默认中文
        noticeMsg: "对不起,没有接收到消息内容！", //消息框内容信息默认中文
        noticeClose: "关闭", //消息框关闭按钮默认中文
        noticePause: "暂停", //消息框暂停按钮默认中文
        noticeLast: "最后一个通知",
        noticeAll: "所有历史通知",
        confirmDel: "你确认要删除",
        close: "关闭",
        //通知中心结束
        sorry: "对不起!", //sorry
        homePage: "首页",
        noShutdown: "不能关闭",
        sort: "排序",
        superiorMenu: "上级菜单",
        menuIcon: "菜单图标",
        menuSort: "菜单排序",
        menuName: "菜单名称",
        menuEnglish: "英语名称",
        menuChinese: "中文名称",
        menuUrl: "菜单网址",
        isshow: "是否显示",
        remark: "备注",
        seeRemark: "查看备注",
        noRemark: "没有备注",
        noData: "没有查询到数据",
        menuEn: "菜单英语",
        manageOperat: "管理操作",
        menuLevelConfine: "对不起，不允许菜单超过三级!",
        menuSortInfoSuccess: "菜单排序成功!",
        menuSortInfoFailed: "菜单排序失败!",
        unableAddEditType: "参数错误,无法识别添加还是修改!",
        addSuccess: "添加成功!",
        addFailed: "添加失败!",
        editSuccess: "修改成功!",
        editFailed: "修改失败!",
        delSuccess: "删除成功!",
        delFailed: "删除失败!",
        searchSuccess: "查询成功!",
        searchFailed: "查询失败!",
        authSuccess: "授权成功!",
        authFailed: "授权失败!",
        success: "成功",
        fail: "失败",
        refresh: "刷新", //refresh
        refreshCurrent: "刷新当前页面",
        ajaxAdd: "添加",
        ajaxEdit: "修改",
        ajaxDel: "删除",
        ajaxAuth: "授权",
        ajaxSetAuth: "权限设置",
        search: "查询",
        clearSearch: "清空查询",
        goback: "返回",
        selectIcons: "选择图标",
        topMenu: "顶级菜单",
        addTopMenu: "添加顶级菜单",
        addSubMenu: "添加子菜单",
        sop: "标准操作", //标准操作
        closeCurrentTab: "关闭当前选项卡",
        closeElseTab: "关闭其他选项卡",
        closeAllTab: "关闭所有选项卡",
        isMenu: "没有菜单", //no menus
        localStorageMenu: "本地缓存中没有菜单数据",
        menuPid: "无法获取指定父ID！",
        getMenuData: "获取菜单数据失败",
        loadMenuData: "正在加载菜单数据!",
        menuDataUrl: "请为菜单项配置数据源【Url】",
        on: "开",
        off: "关",
        reloadDatagrid: "重载数据",
        //切换语言开始
        currentLanguage: "当前语言",
        styleSwitch: "切换到",
        styleDefault: "默认风格",
        styleMac: "Mac风格",
        //切换语言结束
        //登录开始
        backStageLogin: "后台登录",
        login: "登录",
        username: "用户名",
        password: "密码",
        loginMsg: "稍等...,正在为你登录!",
        loginSuccess: "登录成功，3秒后位你跳转!",
        loginPassFail: "用户名或密码错误!",
        loginDisabled: "该账号已被封停!",
        //登录结束
        //用户操作开始
        userLogout: "退出",
        logoutMsg: "稍等...,正在为你退出!",
        logoutMsgFail: "退出失败",
        logoutMsgSuccess: "退出成功",
        unknownError: "未知错误!",
        //用户操作结束
        //easyui validatebox扩展开始
        isUsername: "请输入用户名!",
        isPassword: "请输入密码!",
        isNickname: "请输入昵称!",
        isEmail: "请输入邮箱!",
        isRemark: "请输入备注!",
        usernameVali: "用户名必须字母开头、5-16位字符，字母数字下划线!",
        passwordVali: "密码必须字母开头、5-16位字符、字母数字下划线!",
        passwordEdVali: "两次输入的密码不一致!",
        nicknameVali: "昵称只允许汉字、英文字母、数字、减号及下划线、2-80位字符!",
        remarkVali: "备注只允许汉字、英文字母、数字、减号、逗号、句号、感叹号及下划线、2-255位字符!",
        sqlTypeVali: "数据库类型只能是sqlite3,mysql!",
        portVali: "端口号必须是1-5位正整数!",
        readTimeoutVali: "读取超时必须是1-3位正整数!",
        writeTimeoutVali: "写入超时必须是1-3位正整数!",
        sortVali: "排序必须是1-3位正整数!",
        idVali: "ID必须是1-8位正整数!",
        englishVali: "英文名称只允许英文和空格,2-80位字符!",
        //easyui validatebox扩展结束
        //配置文件开始
        cfgRows: "项",
        cfgName: "配置名称",
        cfgValue: "配置值",
        cfgRemark: "配置注释",
        cfgSql: "数据库配置",
        cfgServer: "服务器配置",
        cfgUseWhatSql: "使用哪种数据库",
        cfgPort: "端口号",
        cfgReadTimeout: "读取超时",
        cfgWriteTimeout: "写入超时",
        cfgCookie: "Cookie配置",
        cfgLanguage: "默认语言",
        cfgLanguageVali: "默认语言只允许英文和空格,2-80位字符!",
        cfgCommon: "公共配置",
        cfgTerminalLog: "终端打印操作日志开关",
        cfgSqlLog: "数据库操作日志开关",
        cfgAjaxPolling: "获取系统占用",
        //配置文件结束
        //用户管理开始
        userAdd: "添加用户",
        userAdminIsNo: "超级管理员不允许编辑",
        userBatchDel: "批量删除用户",
        userName: "用户名",
        userRoleName: "角色名",
        userCreateTime: "创建时间",
        userCreateIp: "创建IP",
        userEmail: "邮箱",
        userPwd: "密码",
        userPwdEd: "确认密码",
        userNickname: "昵称",
        userRoleGroup: "所属角色",
        userStatus: "是否启用",
        userIsRole: "角色编号必须是正整数1到8位",
        userIsStatus: "是否启用必须是1或0",
        userCheckUsername: "用户名已存在!",
        userCheckNickname: "昵称已存在!",
        userCheckEmail: "邮箱已存在!",
        userId: "用户ID不合法!",
        userIdGroup: "用户ID数组不合法!",
        userAdminNoDel: "超级管理员不允许删除!",
        userSearch: "查询用户名",
        userTo: "到",
        //用户管理结束
        //后台菜单开始
        menuCheckName: "中文名称已存在!",
        menuCheckUrl: "菜单网址已存在!",
        menuCheckEn: "英文名称已存在!",
        menuTestPid: "上级菜单ID输入不合法!",
        menuTestName: "中文名称只允许汉字、英文字母及空格、2-80位字符!",
        menuTestSort: "菜单排序输入不合法!",
        menuTestIcon: "菜单图标输入不合法!",
        menuTestIsshow: "菜单显示输入不合法!",
        menuTestUrl: "菜单网址只允许英文和左斜杠,2-80位字符!",
        //后台菜单结束
        //角色管理开始
        roleName: "角色名称",
        roleChinese: "角色中文名称",
        roleEnglish: "角色英文名称",
        roleSort: "角色排序",
        roleRemark: "角色备注",
        roleAdd: "添加角色",
        roleBatchDel: "批量删除角色",
        roleSearch: "查询角色名",
        roleIds: "授权id组不合法!",
        roleCheckName: "角色中文名称已存在!",
        roleCheckEn: "角色英文名称已存在!",
        //角色管理结束
        //通用提示开始
        name: "只允许汉字、英文字母及空格、1-80位字符!",
        noAuth: "你没有相关权限!",
        noFound: "抱歉!您访问的页面找不到哦...",
        english: "英语",
        chinese: "简体中文",
        noAuthGetMenu: "您没有权限获取菜单!",
        noUrl: "没有获取到网页的地址!",
        tantoTabs: "对不起!你打开太多选项卡，会导致机器变慢', '提示信息!",
        currentRole: "当前角色",
        jsCopyCss: "复制颜色样式",
        jsCopyRgb: "复制颜色rgb编码",
        jsCopySuccess: "成功!",
        jsCopyFailed: "复制失败!",
        jsCss: "复制Css",
        jsRgb: "复制Rgb",
        status: "状态",
        backstageHome: "后台首页-HemaCms",
        //通用提示结束
        //日志开始
        loginLogTime: "登录时间",
        loginLogIp: "登录IP",
        loginUseragent: "用户浏览器信息",
        loginSearchUser: "查询用户名",
        loginDelMonthLog: "删除一个月前的登录日志",
        oprateDetail: "操作明细",
        oprateUrl: "请求地址",
        oprateInfo: "操作信息",
        oprateMethod: "请求方式",
        oprateLogTime: "操作时间",
        oprateLogIp: "操作IP",
        oprateExcuteTime: "执行时间",
        oprateMillisecond: "毫秒",
        oprateDelMonthLog: "删除一个月前的操作日志",
        oprateDelAllLog: "删除全部操作日志",
        oprateTime: "操作时间",
        //日志结束
        //验证数据开始
        valiEnglish: "请输入有效的英文字母A~Z,a~z!",
        //验证数据结束

    },
};
$('[language]').each(function() {
    var me = $(this);
    var a = me.attr('language').split(':');
    var p = a[0]; //文字放置位置
    var m = a[1]; //文字的标识

    //用户选择语言后保存在cookie中，这里读取cookie中的语言版本
    var lan = $.cookie('back-language');

    //选取语言文字
    // switch (lan) {
    //     case 'cn':
    //         var t = cn[m]; //这里cn[m]中的cn是上面定义的json字符串的变量名，m是json中的键，用此方式读取到json中的值
    //         break;
    //     case 'en':
    //         var t = en[m];
    //         break;
    //     default:
    //         var t = hk[m];
    // }
    var t = '';
    switch (lan) {
        case 'cn':
            t = language['cn'][m];
            break;
        case 'en':
            t = language['en'][m];
            break;
        default:
            t = language['cn'][m];
            break;
    }
    //如果所选语言的json中没有此内容就选取其他语言显示
    if (t == undefined) t = language['cn'][m];
    if (t == undefined) t = language['en'][m];

    if (t == undefined) return true; //如果还是没有就跳出

    //文字放置位置有（html,val等，可以自己添加）
    switch (p) {
        case 'html':
            me.html(t);
            break;
        case 'val':
        case 'value':
            me.val(t);
            break;
        default:
            me.html(t);
    }

});
