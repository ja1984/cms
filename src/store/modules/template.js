/* eslint no-shadow: ["error", { "allow": ["state"] }] */
/* eslint no-param-reassign: ["error", { "props": false }] */

import { createTemplate, getTemplates } from '@/api/template';

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
    return new Promise((resolve, reject) => {
      createTemplate(payload)
        .then((response) => {
          console.log('create page', response);
          commit('create', payload);
        })
        .catch(error => reject(error));
    });
  },
  list({ commit }) {
    getTemplates().then((response) => {
      if (response.status === 200) {
        commit('add_templates', response.data);
      }
    });
  },
};

// mutations
const mutations = {
  create(state, data) {
    state.templates.push(data);
  },
  add_templates(state, data) {
    state.templates = data;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
