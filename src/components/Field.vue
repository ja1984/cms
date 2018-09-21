<template>
  <div class="form-row">
  <label>
    {{data.name}}
    <input type="text" v-model="data.value" v-if="data.type === 'input'">
    <textarea type="text" v-model="data.value" v-if="data.type === 'textarea'"></textarea>
  </label>
  </div>
</template>

<script>
export default {
  name: 'Field',
  props: {
    data: { type: Object },
  },
};
</script>

<style lang="less" scoped>
</style>
