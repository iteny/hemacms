// import jcc from 'jquery'
import Vue from 'vue'
import Vuex from 'vuex'
// import VueAjax from 'vue-resource/dist/vue-resource'
// import ElementUI from 'element-ui'
// import 'element-ui/lib/theme-default/index.css'
import login from './admin/login/index.vue'
import input from './admin/component/input.vue'
import formInput from './admin/component/formInput.vue'
import form from './admin/component/form.vue'
import item from './admin/component/formItem.vue'
import inputAddon from './admin/component/inputAddon.vue'
import button from './admin/component/formButton.vue'

// Vue.use(ElementUI)
Vue.use(Vuex)
// Vue.component("hm-login",{
//         template: login,
// });
// Vue.component("hm-input",{
//         template: input,
// });

const store = new Vuex.Store({
    state: {
        language: {
            refresh: hm.language.refresh ? hm.language.refresh : 'Refresh',
            reloadDatagrid: hm.language.reloadDatagrid ? hm.language.reloadDatagrid : 'Reload data',
            refreshCurrent: hm.language.refreshCurrent ? hm.language.refreshCurrent : "Refresh Current Page",
            ajaxAdd: hm.language.ajaxAdd ? hm.language.ajaxAdd : "Add",
            ajaxEdit: hm.language.ajaxEdit ? hm.language.ajaxEdit : "Edit",
            ajaxDel: hm.language.ajaxDel ? hm.language.ajaxDel : "Delete",
            ajaxAuth: hm.language.ajaxAuth ? hm.language.ajaxAuth : "Authorization",
            search: hm.language.search ? hm.language.search : "Search",
            clearSearch: hm.language.clearSearch ? hm.language.clearSearch : "Clear search",
            goback: hm.language.goback ? hm.language.goback : "Go back",
            selectIcons: hm.language.selectIcons ? hm.language.selectIcons : "Select icons",
            addTopMenu: hm.language.addTopMenu ? hm.language.addTopMenu : "Add topmenu",
            userName: hm.language.userName ? hm.language.userName : "Account",
            userPwd: hm.language.userPwd ? hm.language.userPwd : "Password",
            userPwdEd: hm.language.userPwdEd ? hm.language.userPwdEd : "Confirm password",
            userNickname: hm.language.userNickname ? hm.language.userNickname : "Nickname",
            userEmail: hm.language.userEmail ? hm.language.userEmail : "Email",
            userRoleGroup: hm.language.userRoleGroup ? hm.language.userRoleGroup : "Endow role",
            userStatus: hm.language.userStatus ? hm.language.userStatus : "Enabled",
            remark: hm.language.remark ? hm.language.remark : "Remark",
            userAdd: hm.language.userAdd ? hm.language.userAdd : "Add account",
            userBatchDel: hm.language.userBatchDel ? hm.language.userBatchDel : "Batch delete accounts",
            userSearch: hm.language.userSearch ? hm.language.userSearch : "Search accounts",
            userCreateTime: hm.language.userCreateTime ? hm.language.userCreateTime : "Create time",
            userTo: hm.language.userTo ? hm.language.userTo : "to",
            roleAdd: hm.language.roleAdd ? hm.language.roleAdd : "Add role",
            roleBatchDel: hm.language.roleBatchDel ? hm.language.roleBatchDel : "Batch delete roles",
            roleSearch: hm.language.roleSearch ? hm.language.roleSearch : "Search role name",
            roleChinese: hm.language.roleChinese ? hm.language.roleChinese : "Role chinese name",
            roleEnglish: hm.language.roleEnglish ? hm.language.roleEnglish : "Role english name",
            roleSort: hm.language.roleSort ? hm.language.roleSort : "Role sort",
            roleRemark: hm.language.roleRemark ? hm.language.roleRemark : "Role remark",
            superiorMenu: hm.language.superiorMenu ? hm.language.superiorMenu : 'Superior menu',
            menuIcon: hm.language.menuIcon ? hm.language.menuIcon : 'Menu icon',
            menuSort: hm.language.menuSort ? hm.language.menuSort : 'Menu sort',
            menuName: hm.language.menuName ? hm.language.menuName : 'Menu name',
            menuEnglish: hm.language.menuEnglish ? hm.language.menuEnglish : 'English name',
            menuChinese: hm.language.menuChinese ? hm.language.menuChinese : 'Chinese name',
            menuUrl: hm.language.menuUrl ? hm.language.menuUrl : 'Menu url',
            isshow: hm.language.isshow ? hm.language.isshow : 'isshow',
            menuEn: hm.language.menuEn ? hm.language.menuEn : 'Menu english',
            sort: hm.language.sort ? hm.language.sort : 'Sort',
            sop: hm.language.sop ? hm.language.sop : 'Sop',
            closeCurrentTab: hm.language.closeCurrentTab ? hm.language.closeCurrentTab : 'closed current tab',
            closeElseTab: hm.language.closeElseTab ? hm.language.closeElseTab : 'closed else tabs',
            closeAllTab: hm.language.closeAllTab ? hm.language.closeAllTab : 'closed all tabs',
            noticeLast: hm.language.noticeLast ? hm.language.noticeLast : 'last one notice',
            noticeAll: hm.language.noticeAll ? hm.language.noticeAll : 'last all notice',
            username: hm.language.username ? hm.language.username : 'Account',
            password: hm.language.password ? hm.language.password : 'Password',
            login: hm.language.login ? hm.language.login : 'Log in',
            userLogout: hm.language.userLogout ? hm.language.userLogout : 'Log out',
            noAuth: hm.language.noAuth ? hm.language.noAuth : 'You have no authority.',
            noFound: hm.language.noFound ? hm.language.noFound : 'Sorry! The page you visited cannot be found...',
            jsCss: hm.language.jsCss ? hm.language.jsCss : 'Copy css',
            jsRgb: hm.language.jsRgb ? hm.language.jsRgb : 'Copy rgb',
            status: hm.language.status ? hm.language.status : 'Status',
            loginSearchUser: hm.language.loginSearchUser ? hm.language.loginSearchUser : 'Search account',
            loginDelMonthLog: hm.language.loginDelMonthLog ? hm.language.loginDelMonthLog : 'Delete the logon log a month ago',
            oprateExcuteTime: hm.language.oprateExcuteTime ? hm.language.oprateExcuteTime : 'Excute time',
            oprateMillisecond: hm.language.oprateMillisecond ? hm.language.oprateMillisecond : 'Millisecond',
            oprateDelMonthLog: hm.language.oprateDelMonthLog ? hm.language.oprateDelMonthLog : 'Delete the oprate log a month ago',
            oprateDelAllLog: hm.language.oprateDelAllLog ? hm.language.oprateDelAllLog : 'Delete all operat logs',
            oprateTime: hm.language.oprateTime ? hm.language.oprateTime : 'Oprate time',

        },
        style: $.cookie('back-style') ? $.cookie('back-style') : 'hm_style_default',


    }
    // language: {
    //     refresh: hm.language.refresh ? hm.language.refresh : 'Refresh',
    // },
    // mutations: {
    //     increment (state) {
    //         state.count++
    //     }
    // }
})
new Vue({
    store,
    el: '#hm-app',
    delimiters: ['[[', ']]'],
    // data:{
    //     textRefresh:
    // },
    components: {
        'hm-login': login,
        'hm-input': input,
        'hm-input-addon': inputAddon,
        'hm-form-input': formInput,
        'hm-form': form,
        'hm-form-item': item,

        'hm-form-button': button,
    },
    // render: h => h(login)
})
