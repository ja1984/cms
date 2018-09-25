<template>
  <textarea type="text" :value="value" :required="required" @input="$emit('input', $event.target.value)"></textarea>
</template>
<script>
export default {
  name: 'TextareaField',
  props: {
    value: { type: String },
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
