<template lang="pug">
draggable.dcgss-mavlink-list(
    v-model="sortedMavmsg",
    :options="{sort:false,group:{name:'mavlink',pull:'clone',put:false}}"
    @update=""
    @start="drag=true",
    @end="drag=false")
  div(v-for="element in sortedMavmsg") {{element.name}}
</template>

<script>
import draggable from 'vuedraggable';

export default {
  name: 'MAVLinkList',
  components: {
    draggable,
  },
  data() {
    return {};
  },
  computed: {
    sortedMavmsg() {
      return Object.values(this.$store.state.mavlinkx.mavlink).sort(
        (a, b) => a.name.localeCompare(b.name));
    },
  },

};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="stylus" scoped>
.dcgss-mavlink-list
  list-style none
  padding 0
  margin 0
  text-align left
  min-height 64px
  div
    padding 5px
    border 1px solid whitesmoke
    font-size 0.7em
    user-select: none
    &:hover
      cursor move
</style>
