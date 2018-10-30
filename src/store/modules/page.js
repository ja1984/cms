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
  byId: state => id => state.pages.find(x => x.id === id),
};

// actions
const actions = {
  create({ commit }, payload) {
    commit('create', payload);
  },
  update({ commit }, payload) {
    commit('update', payload);
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
  update(state, data) {
    const pageIndex = state.pages.findIndex(x => x.id === data.id);
    if (pageIndex > -1) {
      state.pages[pageIndex] = data;
    }
    // state.pages.push(data);
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
