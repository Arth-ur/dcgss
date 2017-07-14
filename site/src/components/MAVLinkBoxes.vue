<template lang="pug">
draggable.pure-g.dcgss-mavlink-boxes(
    :class="{centered:updatedMAVmsg.length==0}"
    v-model="mavmsg",
    :options="{group:'mavlink',handle:'.title'}",
    filter=".ignore-elements",
    @start="drag=true",
    @end="drag=false")
  MAVLinkBox(:value="element.name",
    v-for="(element,i) in updatedMAVmsg",
    @destroy="mavmsg.splice(i, 1)")
  div.pure-u-1.pure-u-md-1-2.dcgss-empty.ignore-elements(v-if="updatedMAVmsg.length == 0")
    h1 Instructions
    p
      | On the left hand side of the screen, choose the UDP port to listen to
      | (default is 4550), then click on #[code Connect].
    p
      | Drag and drop elements from MAVLink to this zone to show the associated
      | MAVLink message.

</template>

<script>
import draggable from 'vuedraggable';
import MAVLinkBox from './MAVLinkBox';

export default {
  name: 'MAVLinkBoxes',
  components: {
    draggable,
    MAVLinkBox,
  },
  data() {
    return {
      mavmsg: [],
    };
  },
  computed: {
    updatedMAVmsg() {
      return this.mavmsg.map(x => this.$store.state.mavlinkx.mavlink[x.name]);
    },
  },
};
</script>

<style lang="stylus" scoped>
.dcgss-mavlink-boxes
  min-height 64px
  height 100%
.centered
  justify-content center

.dcgss-empty
  position absolute
  h1
    text-align center

</style>
