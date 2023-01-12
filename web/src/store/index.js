import user from './module/user'
import app from './module/app'
import {createStore} from 'vuex'

export default createStore({
    state: {},
    getters: {},
    mutations: {
        updateUser(state, user) {
            state.user.username = user.username;
        }
    },

    actions: {
    },
    modules: {
        user: user,
        app: app,
    }
})
