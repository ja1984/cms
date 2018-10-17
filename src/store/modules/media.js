/* eslint no-shadow: ["error", { "allow": ["state"] }] */
/* eslint no-param-reassign: ["error", { "props": false }] */
// initial state
const state = {
  media: [
    // {
    //   id: 0,
    //   type: 'img',
    //   url: 'https://via.placeholder.com/250x150',
    //   filename: 'file0',
    // },
    // {
    //   id: 1,
    //   type: 'img',
    //   url: 'https://via.placeholder.com/250x150',
    //   filename: 'file0',
    // },
    // {
    //   id: 2,
    //   type: 'img',
    //   url: 'https://via.placeholder.com/250x150',
    //   filename: 'file0',
    // },
    // {
    //   id: 3,
    //   type: 'img',
    //   url: 'https://via.placeholder.com/250x150',
    //   filename: 'file0',
    // },
    // {
    //   id: 4,
    //   type: 'img',
    //   url: 'https://via.placeholder.com/250x150',
    //   filename: 'file0',
    // },
    // {
    //   id: 5,
    //   type: 'img',
    //   url: 'https://via.placeholder.com/250x150',
    //   filename: 'file0',
    // },
    // {
    //   id: 6,
    //   type: 'img',
    //   url: 'https://via.placeholder.com/250x150',
    //   filename: 'file0',
    // },
    // {
    //   id: 7,
    //   type: 'img',
    //   url: 'https://via.placeholder.com/250x150',
    //   filename: 'file0',
    // },
    // {
    //   id: 8,
    //   type: 'img',
    //   url: 'https://via.placeholder.com/250x150',
    //   filename: 'file0',
    // },
    // {
    //   id: 9,
    //   type: 'img',
    //   url: 'https://via.placeholder.com/250x150',
    //   filename: 'file0',
    // },
    // {
    //   id: 10,
    //   type: 'img',
    //   url: 'https://via.placeholder.com/250x150',
    //   filename: 'file0',
    // },
    // {
    //   id: 11,
    //   type: 'img',
    //   url: 'https://via.placeholder.com/250x150',
    //   filename: 'file0',
    // },
    // {
    //   id: 12,
    //   type: 'img',
    //   url: 'https://via.placeholder.com/250x150',
    //   filename: 'file0',
    // },
    // {
    //   id: 13,
    //   type: 'img',
    //   url: 'https://via.placeholder.com/250x150',
    //   filename: 'file0',
    // },
    // {
    //   id: 14,
    //   type: 'img',
    //   url: 'https://via.placeholder.com/250x150',
    //   filename: 'file0',
    // },
  ],
};

// getters
const getters = {
  all: state => state.media,
};

// actions
const actions = {
  create({ commit }, payload) {
    console.log(payload);
    commit('create', payload);
  },
};

// mutations
const mutations = {
  create(state, data) {
    state.media.push(data);
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
