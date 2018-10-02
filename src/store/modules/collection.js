/* eslint no-shadow: ["error", { "allow": ["state"] }] */
/* eslint no-param-reassign: ["error", { "props": false }] */

// initial state
const state = {
  collections: [
    {
      key: 'test-save',
      properties: [
        {
          key: 'prop1',
          value: {
            en: 'cool',
          },
        },
        {
          key: 'prop2withoutLanguage',
          value: {
            default: 'cool',
          },
        },
        {
          key: 'prop3Languages',
          value: {
            en: 'horse',
            sv: 'häst',
          },
        },
      ],
    },
  ],
};

// getters
const getters = {
  get: state => state.collections,
  bySlug: state => slug => state.collections.find(collection => collection.key === slug),
};

// actions
const actions = {
  create({ commit }, payload) {
    return new Promise((resolve) => {
      commit('create', payload);
      resolve();
    });
  },
  delete({ commit }, payload) {
    commit('delete', payload);
  },
  addKey({ commit }, payload) {
    return new Promise((resolve) => {
      commit('addKey', payload);
      resolve();
    });
  },
};

// mutations
const mutations = {
  create(state, data) {
    state.collections.push(data);
  },
  addKey(state, data) {
    const collection = state.collections.find(x => x.key === data.collection);
    collection.properties.push(data.data);
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
