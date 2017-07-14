<template lang="pug">
div
  LineChart#line(
    v-if="plotData.length>0"
    :data="plotData",
    resize=true,
    :xkey="xkey",
    :ykeys="ykeys",
    :labels="labels",
    :line-colors="linecolors",
    grid="true",
    grid-text-weight="bold")
  div
    label(v-for="(checked,label) in active")
      input(type="checkbox",v-model="active[label]")
      span {{label}}
    button(title="clear",@click="plotData=[]") #[i.fa.fa-trash]
</template>

<script>
import Raphael from 'raphael/raphael';
import jquery from 'jquery';
import { LineChart } from 'vue-morris';

global.Raphael = Raphael;
global.jQuery = jquery;

export default {
  name: 'MAVPlot',
  props: ['value'],
  components: {
    LineChart,
  },
  data() {
    return {
      plotData: [],
      unwatch: () => {},
      start: Date.now(),
      xkey: '$time',
      active: {},
    };
  },
  computed: {
    linecolors() {
      return null;
    },
    ykeys() {
      return this.value.fieldnames.filter(f => this.active[f] === true);
    },
    labels() {
      return this.value.fieldnames.filter(f => this.active[f] === true);
    },
  },
  watch: {
    value() {
      const newdata = {};
      newdata.$time = (Date.now());
      this.value.fieldnames
      .filter(f => this.active[f] === true)
      .forEach((f) => {
        newdata[f] = this.value[f];
      });
      this.plotData.push(newdata);
      if (this.plotData.length > 10) {
        this.plotData.splice(0, 1);
      }
    },
  },
  mounted() {
    this.ykeys = this.value.fieldnames;
    this.labels = this.value.fieldnames;
    this.xkey = '$time';
    this.value.fieldnames.forEach((f) => { this.active[f] = true; });
  },
};
</script>

<style lang="stylus" scoped>
</style>
