<template lang="pug">
form.pure-form.pure-form-stacked(v-on:submit.prevent="connect()")
  fieldset
    input.pure-input-1(type="number", step=1, v-bind:value="value", v-on:input="privateValue=$event.target.value", placeholder="Port", :disabled="this.connected")
    button(type="submit",class="pure-button pure-button-primary" v-if="!connected") Connect
    button(type="button",class="pure-button", v-if="connected", @click="disconnect()") Disconnect
</template>

<script>
import ReconnectingWebSocket from 'reconnecting-websocket';

export default {
  name: 'DCGSSConnection',
  components: {},
  data() {
    return {
      privateValue: null,
      connected: false,
      ws: null,
    };
  },
  props: ['value'],
  methods: {
    connect() {
      this.ws = new ReconnectingWebSocket(`ws://localhost:8505/ws?protocol=udp&port=${this.value}`);
      this.ws.addEventListener('open', () => {
        this.ws.binaryType = 'arraybuffer';
        // this.ws.binaryType = 'blob';
      });
      this.ws.addEventListener('message', (m) => {
        this.$store.commit('MAVLINK_DECODE', new Uint8Array(m.data));
      });
      this.connected = true;
    },
    disconnect() {
      this.ws.close(1000, '', { keepClosed: true });
      this.connected = false;
    },
  },
};
</script>

<style lang="stylus">
</style>
