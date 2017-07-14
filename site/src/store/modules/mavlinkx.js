import Vue from 'vue';
import * as types from '../mutation-types';
import mavlink from '../../mavlink';

const mav = new mavlink.MAVLink();

// initial state
const state = {
  mavlink: {},
};

// actions
const actions = {
};

// mutations
const mutations = {
  [types.MAVLINK_DECODE](st, bytes) {
    try {
      const msg = mav.decode(bytes);
      Vue.set(st.mavlink, msg.name, msg);
    } catch (e) {
      // do nothing
    }
  },
};

export default {
  state,
  actions,
  mutations,
};
