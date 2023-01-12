import {createRouter, createWebHashHistory} from 'vue-router'
import routes from './routers'
import store from '@/store'
import ViewUIPlus from 'view-ui-plus'
import {setToken, getToken, canTurnTo, setTitle} from '@/libs/util'
import config from '@/config'
const {homeName} = config

const router = createRouter({
    history: createWebHashHistory(),
    routes: routes
})
const LOGIN_PAGE_NAME = 'login'
const REGISTER_PAGE_NAME = 'register'

const turnTo = (to, access, next) => {
    if (canTurnTo(to.name, access, routes)) next()
    else next({replace: true, name: 'error_401'})
}

router.beforeEach((to, from, next) => {
    ViewUIPlus.LoadingBar.start();
    const token = getToken()
    if (!token && (to.name !== LOGIN_PAGE_NAME && to.name !== REGISTER_PAGE_NAME)) {
        next({
            name: LOGIN_PAGE_NAME
        })
    } else if (!token && (to.name === LOGIN_PAGE_NAME || to.name === REGISTER_PAGE_NAME)) {
        next()
    } else if (token && to.name === LOGIN_PAGE_NAME) {
        next({
            name: homeName
        })
    } else {
        if (store.state.user.hasGetInfo) {
            turnTo(to, store.state.user.access, next)
        } else {
            store.dispatch('getUserInfo').then(user => {
                turnTo(to, user.access, next)
            }).catch(() => {
                setToken('')
                next({
                    name: 'login'
                })
            })
        }
    }
})

router.afterEach(to => {
    setTitle(to, window.i18n)
    ViewUIPlus.LoadingBar.finish()
    window.scrollTo(0, 0)
})

export default router
