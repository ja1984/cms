/* eslint no-shadow: ["error", { "allow": ["state"] }] */
/* eslint no-param-reassign: ["error", { "props": false }] */
import { getPages } from '@/api/page';

// initial state
const state = {
  pages: [],
};

// getters
const getters = {
  pages: state => state.pages,
};

// actions
const actions = {
  create({ commit }, payload) {
    commit('create', payload);
  },
  list({ commit }) {
    getPages().then((response) => {
      console.log(response);
      if (response.status === 200) {
        commit('add_pages', response.data);
      }
    });
  },
};

// mutations
const mutations = {
  create(state, data) {
    state.pages.push(data);
  },
  add_pages(state, data) {
    state.pages = data;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
