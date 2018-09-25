<template>
  <input type="number" placeholder="Enter number" :value="value" @input="$emit('input', parseInt($event.target.value))">
</template>
<script>
export default {
  name: 'InputField',
  props: {
    value: { type: Number },
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
