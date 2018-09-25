import Vue from 'vue';
import Vuex from 'vuex';
import template from '@/store/modules/template';
import page from '@/store/modules/page';
import language from '@/store/modules/language';

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== 'production';

export default new Vuex.Store({
  modules: {
    template,
    page,
    language,
  },
  strict: debug,
});
