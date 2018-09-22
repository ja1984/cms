<template>
  <div class="form-row">
    <label>
      <span>{{data.name}}
        <div class="hastooltip" v-if="data.tooltip" :data-balloon="data.tooltip" data-balloon-pos="up">
          <i class="far fa-question-circle">
          </i>
        </div>
      </span>
      <input type="text" v-model="data.value" :required="data.required" v-if="data.type === 'input'">
      <textarea type="text" v-model="data.value" :required="data.required" v-if="data.type === 'textarea'"></textarea>
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
.hastooltip {
  display: inline-block;
}
</style>
