<template>
  <div class="fields">
    <button class="button button-primary button-small" @click="addField">Add field</button>
    <!-- <Field v-for="childField in childFields" :key="childField.id" :data="childField"></Field> -->
  </div>
</template>
<script>
import Field from '@/components/Field.vue';

export default {
  name: 'RepeaterField',
  components: {
    Field,
  },
  methods: {
    addField() {
      console.log(this, this.value);
      this.$emit('addField', this.id);
    },
  },
  props: {
    childFields: { type: Array, default: () => [] },
    value: { type: Object },
    id: { type: String },
  },
};
</script>

<style lang="less" scoped>
.hastooltip {
  display: inline-block;
}
</style>
