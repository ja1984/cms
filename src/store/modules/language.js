/* eslint no-shadow: ["error", { "allow": ["state"] }] */
/* eslint no-param-reassign: ["error", { "props": false }] */

// initial state
const state = {
  languages: [],
};

// getters
const getters = {
  get: state => state.languages,
};

// actions
const actions = {
  create({ commit }, payload) {
    commit('create', payload);
  },
};

// mutations
const mutations = {
  create(state, data) {
    state.languages.push(data);
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
