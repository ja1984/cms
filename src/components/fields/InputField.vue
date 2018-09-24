<template>
  <label>
    {{name}}
    <input type="text" :value="value" @input="$emit('input', $event.target.value)">
  </label>
</template>
<script>
export default {
  name: 'InputField',
  props: {
    value: { type: String },
    name: { type: String },
  },
};
</script>

<style lang="less" scoped>
</style>
