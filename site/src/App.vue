<template lang="pug">
div#app
  div.pure-g
    div.pure-u-1-5.pure-hidden-sm.pure-hidden-xs
      MAVLinkList
    div.pure-u-1.pure-u-md-4-5
        MAVLinkBoxes

</template>

<script>
import MAVLinkList from './components/MAVLinkList';
import MAVLinkBoxes from './components/MAVLinkBoxes';

export default {
  name: 'app',
  components: {
    MAVLinkList,
    MAVLinkBoxes,
  },
};
</script>

<style lang="stylus">
#app
  font-family: 'Avenir', Helvetica, Arial, sans-serif
  -webkit-font-smoothing: antialiased
  -moz-osx-font-smoothing: grayscale
  color: #2c3e50
</style>
