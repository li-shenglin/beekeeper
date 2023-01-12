// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import {createApp} from 'vue'
import ViewUIPlus from 'view-ui-plus'
import App from './App.vue'
import router from './router'
import store from './store'
import 'view-ui-plus/dist/styles/viewuiplus.css'
import i18n from '@/locale'
import config from '@/config'
import './index.less'
import '@/assets/icons/iconfont.css'


let vue = createApp(App);
vue.use(i18n);
vue.use(router);
vue.use(store);
vue.use(ViewUIPlus, {
    i18n
});
window.i18n = i18n.global
window.$t = i18n.global.t
vue.config.productionTip = false
vue.config.globalProperties.$config = config
vue.mount("#app")

