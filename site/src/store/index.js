import Vue from 'vue';
import Vuex from 'vuex';
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
  },
  modules: {
    mavlinkx,
  },
  strict: debug,
//  plugins: debug ? [createLogger()] : [],
});
