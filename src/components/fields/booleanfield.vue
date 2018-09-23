<template>
  <div class="form-row">
    <check-box v-model="data.value">{{this.props.name}}</check-box>
  </div>
</template>

<script>
import CheckBox from '@/components/CheckBox.vue';

export default {
  name: 'booleanfield',
  components: {
    CheckBox,
  },
  props: {
    data: { type: Object },
  },
};
</script>

<style lang="less" scoped>
</style>
