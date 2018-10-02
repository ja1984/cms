<template>
  <select :value="value" :required="required" @change.stop="$emit('input', $event.target.value)">
    <option disabled value="">Please select one</option>
    <option v-for="option in options" :key="option.id" :value="option.value">{{option.value}}</option>
  </select>
</template>
<script>
export default {
  name: 'SelectField',
  props: {
    value: { type: String },
    options: { type: Array, default: () => [] },
    name: { type: String },
    required: { type: Boolean, default: false },
    tooltip: { type: String },
  },
};
</script>

<style lang="less" scoped>
.hastooltip {
  display: inline-block;
}
</style>
