<template lang="pug">
table.pure-table
  thead
    tr
      th fieldname
      th(style="width:100%") value
  tbody
    tr(v-for="i in value.order_map")
      td {{value.fieldnames[i]}}
      td {{value[value.fieldnames[i]]}}
</template>

<script>
export default {
  name: 'MAVTable',
  props: ['value'],
};
</script>

<style lang="stylus" scoped>
table
  width 100%
</style>
