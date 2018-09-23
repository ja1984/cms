/* eslint no-shadow: ["error", { "allow": ["state"] }] */
/* eslint no-param-reassign: ["error", { "props": false }] */

// initial state
const state = {
  templates: [],
};

// getters
const getters = {
  all: state => state.templates,
  byId: state => id => state.templates.find(t => t.id === id),
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
    state.templates.push(data);
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
