<template>
  <label>
    {{name}}
    <input type="number" placeholder="Enter number" :value="value" @input="$emit('input', parseInt($event.target.value))">
  </label>
</template>
<script>
export default {
  name: 'InputField',
  props: {
    value: { type: Number },
    name: { type: String },
  },
};
</script>

<style lang="less" scoped>
</style>
