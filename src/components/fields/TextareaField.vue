<template>
  <label>
    {{name}}
    <textarea type="text" :value="value" @input="$emit('input', $event.target.value)"></textarea>
  </label>
</template>
<script>
export default {
  name: 'TextareaField',
  props: {
    value: { type: String },
    name: { type: String },
  },
};
</script>

<style lang="less" scoped>
</style>
