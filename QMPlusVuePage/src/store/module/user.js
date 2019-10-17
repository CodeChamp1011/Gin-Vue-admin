import { login } from '@/api/user'
import router from '@/router/index'
export const user = {
    namespaced: true,
    state: {
        userInfo: {
            uuid: "",
            nickName: "",
            headerImg: "",
            authority: "",
        },
        token: "",
        expiresAt: ""
    },
    mutations: {
        setUserInfo(state, userInfo) {
            // 这里的 `state` 对象是模块的局部状态
            state.userInfo = userInfo
        },
        setToken(state, token) {
            // 这里的 `state` 对象是模块的局部状态
            state.token = token
        },
        setExpiresAt(state, expiresAt) {
            // 这里的 `state` 对象是模块的局部状态
            state.expiresAt = expiresAt
        },
        LoginOut(state) {
            state.userInfo = {}
            state.token = ""
            state.expiresAt = ""
            router.push({ name: 'login', replace: true })
        },
        ResetUserInfo(state, userInfo = {}) {
            state.userInfo = {...state.userInfo,
                ...userInfo
            }
        }
    },
    actions: {
        async LoginIn({ commit }, loginInfo) {
            const res = await login(loginInfo)
            commit('setUserInfo', res.data.user)
            commit('setToken', res.data.token)
            commit('setExpiresAt', res.data.expiresAt)
            if (res.success) {
                const redirect = router.history.current.query.redirect
                if (redirect) {
                    router.push({ path: redirect, replace: true })
                } else {
                    router.push({ path: '/layout/dashbord', replace: true })
                }
            }
        }
    },
    getters: {
        userInfo(state) {
            return state.userInfo
        },
        token(state) {
            return state.token
        },
        expiresAt(state) {
            return state.expiresAt
        }
    }
}