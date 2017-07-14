<template lang="pug">
div#app
  div.pure-g
    div.pure-u-1.pure-u-md-1-5.left-bar
      div.left-bar-title MAVLink Messages
      MAVLinkList
      div.left-bar-title
        div Tools
        div.tool(title="Onboard parameters") #[i.fa.fa-wrench]
      div.left-bar-title Connection
        Connection(:value=14550,@input="connect")
      div.left-bar-title Infos
      div
        small Throttling: {{throttling}}ms
    div.pure-u-1.pure-u-md-4-5.boxes
        MAVLinkBoxes

</template>

<script>
import MAVLinkList from './components/MAVLinkList';
import MAVLinkBoxes from './components/MAVLinkBoxes';
import Connection from './components/Connection';

export default {
  name: 'app',
  components: {
    MAVLinkList,
    MAVLinkBoxes,
    Connection,
  },
  methods: {
    connect(port) {
      console.log(port);
    },
  },
  computed: {
    throttling() {
      return this.$store.state.throttling;
    },
  },
};
</script>

<style lang="stylus">
html
  height 100%
body
  height 100%
#app
  height 100%
  font-family: 'Avenir', Helvetica, Arial, sans-serif
  -webkit-font-smoothing: antialiased
  -moz-osx-font-smoothing: grayscale
  color: #2c3e50
  &>div.pure-g
    height 100%
    &>div.boxes
      height 100%
.left-bar
  .left-bar-title
    color gray
    text-transform uppercase
    padding 0.5em 0
  .tool
    border-radius 1em
    border 1px solid #ccc
    padding 1em
    display inline-block
    &:hover
      cursor move
      background whitesmoke
</style>
