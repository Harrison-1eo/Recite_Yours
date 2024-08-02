import { createStore } from 'vuex'

export default createStore({
    state: {
        user: null,
    },
    mutations: {
        setUser(state, user) {
            state.user = user
        },
        clearUser(state) {
            state.user = null
        }
    },
    actions: {
        login({ commit }, user) {
            // 模拟登录
            commit('setUser', user)
        },
        logout({ commit }) {
            commit('clearUser')
        }
    },
    getters: {
        isAuthenticated: state => !!state.user,
        getUser: state => state.user,
    }
})
