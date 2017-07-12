<template lang="pug">
draggable.pure-g.dcgss-mavlink-boxes(
    v-model="mavmsg",
    :options="{group:'mavlink'}"
    @start="drag=true",
    @end="drag=false")
  div.pure-u-1.pure-u-lg-1-2.pure-u-xl-1-3(v-for="element in mavmsg")
    div.dcgss-mavlink-box {{element.name}}

</template>

<script>
import draggable from 'vuedraggable';

export default {
  name: 'MAVLinkBoxes',
  components: {
    draggable,
  },
  data() {
    return {
      mavmsg: [
        { name: 'HEARTBEAT' },
        { name: 'SYS_STATUS' },
        { name: 'SYSTEM_TIME' },
      ],
    };
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="stylus" scoped>
div.dcgss-mavlink-box
  background whitesmoke
  margin 0.5em

.dcgss-mavlink-boxes
  min-height 64px
</style>
