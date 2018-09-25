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
  delete({ commit }, payload) {
    commit('delete', payload);
  },
};

// mutations
const mutations = {
  create(state, data) {
    data.isDefault = state.languages.length === 0;
    state.languages.push(data);
  },
  delete(state, data) {
    const languageIndex = state.languages.findIndex(language => language.slug === data.slug);

    if (languageIndex > -1) {
      state.languages.splice(languageIndex, 1);
    }

    const defaultLanguageIndex = state.languages.findIndex(language => language.isDefault);

    if (defaultLanguageIndex > -1) return;
    state.languages[0].isDefault = true;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
