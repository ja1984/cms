import Vue from 'vue';
import GAuth from 'vue-google-oauth2';

import App from './App.vue';
import router from './router';
import store from './store';

Vue.config.productionTip = false;
Vue.use(GAuth, { clientId: '628934306033-9o4jjcisjt20ihltmkbck7jg6d3kl5mf.apps.googleusercontent.com', scope: 'profile email https://www.googleapis.com/auth/plus.login' });
new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app');
