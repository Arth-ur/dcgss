import Vue from 'vue';
import Vuex from 'vuex';
import jQuery from 'jquery';
import mavlinkx from './modules/mavlinkx';

// import * as actions from './actions'
// import * as getters from './getters'
// import cart from './modules/cart'
// import products from './modules/products'
// import createLogger from '../../../src/plugins/logger'

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== 'production';

export default new Vuex.Store({
//   actions,
//   getters,
  state: {
    throttling: 1000,
    version: '0',
  },
  modules: {
    mavlinkx,
  },
  mutations: {
    version(state, version) {
      state.version = version;
    },
  },
  actions: {
    checkVersion(context) {
      jQuery.get('http://localhost:8505/version')
        .done(v => context.commit('version', v))
        .fail(console.error);
    },
  },
  strict: debug,
//  plugins: debug ? [createLogger()] : [],
});
