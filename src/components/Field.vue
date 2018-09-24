<template>
  <div class="form-row">
    <!-- <span>{{data.name}}
        <div class="hastooltip" v-if="data.tooltip" :data-balloon="data.tooltip" data-balloon-pos="up">
          <i class="far fa-question-circle">
          </i>
        </div>
      </span> -->
    <input-field v-model="data.value" :name="data.name" :required="data.required" v-if="data.type === 'input'"></input-field>
    <textarea-field v-model="data.value" :name="data.name" :required="data.required" v-if="data.type === 'textarea'"></textarea-field>
    <boolean-field v-model="data.value" :required="data.required" v-if="data.type === 'boolean'">{{data.name}}</boolean-field>
  </div>
</template>

<script>

import InputField from '@/components/fields/InputField.vue';
import TextareaField from '@/components/fields/TextareaField.vue';
import BooleanField from '@/components/fields/BooleanField.vue';

export default {
  name: 'Field',
  components: {
    InputField,
    TextareaField,
    BooleanField,
  },
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
