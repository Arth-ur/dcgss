<template lang="pug">
div.pure-u-1.dcgss-drg(
  :class="{'pure-u-lg-1-2':size!=='large','pure-u-xl-1-3':size==='small'}")
  div.dcgss-mavlink-box
    div.title.pure-g
      span.pure-u-1.pure-u-sm-1
        button.button-xsmall.button-logging(title="Start logging",
          :class="{'pure-button':true, 'pure-button-active':logging}",
          @click="logging=!logging") #[i.fa.fa-circle]
        span.msgname {{cachedMSG.name}}
      span.toolbox.pure-u-1
        select(v-model="view")
          option(value="raw") Raw
          option(value="table") Table
          option(value="plot") Plot
        span.pure-button-group
          button.pure-button.pure-visible-xl.size(:class="{'pure-button-active':size==='large'}",@click="resize('large')") #[i.square.large-square]
          button.pure-button.pure-visible-xl.size(:class="{'pure-button-active':size==='medium'}",@click="resize('medium')") #[i.square.medium-square]#[i.square.medium-square]
          button.pure-button.pure-visible-xl.size(:class="{'pure-button-active':size==='small'}",@click="resize('small')") #[i.square.small-square]#[i.square.small-square]#[i.square.small-square]
          button.pure-button.pure-visible-lg.size(:class="{'pure-button-active':size==='large'}",@click="resize('large')") #[i.square.large-square]
          button.pure-button.pure-visible-lg.size(:class="{'pure-button-active':size!=='large'}",@click="resize('medium')") #[i.square.medium-square]#[i.square.medium-square]
        button.pure-button(@click="$emit('destroy')") #[i.fa.fa-times]
    div.content
      Raw(v-if="view=='raw'",:value="cachedMSG")
      MAVTable(v-if="view=='table'",:value="cachedMSG")
      MAVPlot(v-if="view=='plot'",:value="cachedMSG")
</template>

<script>
import download from 'downloadjs';
import _ from 'lodash';
import Raw from './views/Raw';
import MAVTable from './views/MAVTable';
import MAVPlot from './views/MAVPlot';

export default {
  name: 'MAVLinkBox',
  components: {
    Raw,
    MAVTable,
    MAVPlot,
  },
  props: ['value'],
  data() {
    return {
      view: 'table',
      size: 'large',
      lastUpdate: Date.now(),
      cachedMSG: {},
      unwatch: () => {},
      logging: false,
      logdata: '',
    };
  },
  computed: {
    msg() {
      // return this.$store.state.mavlinkx.mavlink[this.value];
      return this.cachedMSG;
    },
    throttledUpdate() {
      return _.throttle(this.update, this.$store.state.throttling);
    },
  },
  watch: {
    logging(l) {
      if (l !== true) {
        download(this.logdata, `${this.value}.txt`, 'text/plain');
      } else {
        this.logdata = 'time,';
        this.logdata += this.cachedMSG.order_map.map(i =>
          this.cachedMSG.fieldnames[i]).join(',');
        this.logdata += '\n';
      }
    },
  },
  methods: {
    resize(newSize) {
      if (newSize === this.size) {
        return;
      }
      this.size = newSize;
      this.$emit('resize', newSize);
    },
    update() {
      this.cachedMSG = this.$store.state.mavlinkx.mavlink[this.value];
    },
  },
  mounted() {
    this.unwatch = this.$store.watch(
      state => state.mavlinkx.mavlink[this.value],
      () => {
        this.throttledUpdate();
        if (this.logging) {
          this.logdata += Date.now();
          this.logdata += ',';
          this.logdata += this.cachedMSG.order_map.map(i =>
            this.cachedMSG[this.cachedMSG.fieldnames[i]]).join(',');
          this.logdata += '\n';
        }
      },
      { deep: false },
    );
    this.throttledUpdate();
  },
  beforeDestroy() {
    this.unwatch();
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="stylus" scoped>
div.dcgss-mavlink-box
  background white
  border 1px solid whitesmoke
  border-radius 1em
  margin 0.5em
  padding 0.1em

  div.title
    padding-left 1em
    border-bottom 1px solid whitesmoke;
    cursor move

.toolbox
  text-align right
button.size
  padding 0.5em
  width 2.5em
i.square
  border 1px solid grey
  background grey
  display inline-block
  vertical-align center
  &:last-child
    margin-right 0
i.large-square
  width 1.5em
  height 0.8em
i.medium-square
  width 0.7em
  height 0.6em
  margin-right 0.05em
i.small-square
  width 0.4em
  height 0.4em
  margin-right 0.05em
</style>
