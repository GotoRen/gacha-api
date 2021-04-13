export const state = () => ({
    login: false,
    id:0,
    username: '',
    token:''
})

export const mutations = {
    changeLogin(state, bool) {
        state.login = bool
    },
    changeUserData(state, payload) {
        state.id = payload.id;
        state.username = payload.username;
        state.token = payload.token;
    }
}