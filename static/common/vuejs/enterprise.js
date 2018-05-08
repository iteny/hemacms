// import jcc from 'jquery'
import Vue from 'vue'
import Vuex from 'vuex'
// import VueAjax from 'vue-resource/dist/vue-resource'
// import ElementUI from 'element-ui'
// import 'element-ui/lib/theme-default/index.css'
// import login from './admin/login/index.vue'
// import input from './admin/component/input.vue'
// import formInput from './admin/component/formInput.vue'
// import form from './admin/component/form.vue'
// import item from './admin/component/formItem.vue'
// import inputAddon from './admin/component/inputAddon.vue'
// import button from './admin/component/formButton.vue'

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
        // 'hm-login': login,
        // 'hm-input': input,
        // 'hm-input-addon': inputAddon,
        // 'hm-form-input': formInput,
        // 'hm-form': form,
        // 'hm-form-item': item,
        // 'hm-form-button': button,
    },
    // render: h => h(login)
})
