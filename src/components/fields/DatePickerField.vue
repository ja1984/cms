<template>
  <label>
    {{name}}
    <datepicker :value="value" @selected="selected"></datepicker>
  </label>
</template>
<script>
import Datepicker from 'vuejs-datepicker';

export default {
  name: 'DatePickerField',
  components: {
    Datepicker,
  },
  props: {
    value: { type: Date },
    name: { type: String },
  },
  methods: {
    selected(date) {
      this.$emit('input', date);
    },
  },
};
</script>

<style lang="less" scoped>
</style>
