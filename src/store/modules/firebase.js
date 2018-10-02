/* eslint no-shadow: ["error", { "allow": ["state"] }] */
/* eslint no-param-reassign: ["error", { "props": false }] */
import firebase from 'firebase/app';
import 'firebase/database';
import 'firebase/auth';

const FBConfig = {
  apiKey: 'AIzaSyB1wUoIVsnxN-JqHWPI0kHT88b-p3WoD8g',
  authDomain: 'cogcms-383a0.firebaseapp.com',
  databaseURL: 'https://cogcms-383a0.firebaseio.com',
  projectId: 'cogcms-383a0',
  storageBucket: 'cogcms-383a0.appspot.com',
  messagingSenderId: '544607828602',
};
firebase.initializeApp(FBConfig);

// initial state
const state = {
  user: null,
  error: '',
  initialized: false,
};

// getters
const getters = {
  user: state => state.user,
};

// actions
const actions = {
  init({ commit, state }) {
    if (state.initialized) {
      return;
    }
    state.initialized = true;

    firebase.auth().onAuthStateChanged(function(user) {
      if (user) {
        console.log(user);
        commit('SET_USER', user);
      } else {
        commit('SET_USER', null);
      }
    });
  },
  login({ state, dispatch }, credentials) {
    return dispatch('init')
      .then(() =>
        firebase.auth().signInWithEmailAndPassword(credentials.email, credentials.password),
      )
      .then(res => {
        console.log(res);
      })
      .catch(function(error) {
        console.log(error);
        state.error = error;
      });
  },
  signup({ state, dispatch }, credentials) {
    return dispatch('init')
      .then(() =>
        firebase.auth().createUserWithEmailAndPassword(credentials.email, credentials.password),
      )
      .then(res => {
        console.log(res);
      })
      .catch(function(error) {
        console.log(error);
        state.error = error;
      });
  },
};

// mutations
const mutations = {
  SET_USER(state, user) {
    state.user = user;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
