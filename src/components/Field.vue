<template>
  <div class="field-wrapper">
    <div class="form-row">
      <label>
        <span class="label-text">{{data.name}}
          <div class="hastooltip" v-if="data.tooltip" :data-balloon="data.tooltip" data-balloon-pos="up">
            <i class="far fa-question-circle">
            </i>
          </div>
        </span>
        <select-field v-model="data.value" :name="data.name" :options="data.options" :required="data.required" :tooltip="data.tooltip" v-if="data.type === 'select'"></select-field>
        <input-field v-model="data.value" :name="data.name" :required="data.required" :tooltip="data.tooltip" v-if="data.type === 'input'"></input-field>
        <number-field v-model="data.value" :name="data.name" :required="data.required" :tooltip="data.tooltip" v-if="data.type === 'number'"></number-field>
        <textarea-field v-model="data.value" :name="data.name" :required="data.required" :tooltip="data.tooltip" v-if="data.type === 'textarea'"></textarea-field>
        <boolean-field v-model="data.value" class="checkbox-block" :name="data.name" :required="data.required" :tooltip="data.tooltip" v-if="data.type === 'boolean'">{{data.name}}</boolean-field>
        <date-picker-field v-model="data.value" :name="data.name" :required="data.required" :tooltip="data.tooltip" v-if="data.type === 'date'"></date-picker-field>
        <repeater-field v-model="data.value" :id="data.id" :name="data.name" @addField="addField" :childFields="data.childFields" :required="data.required" :tooltip="data.tooltip" v-if="data.type === 'repeater'"></repeater-field>
      </label>
    </div>
  </div>
</template>

<script>

import InputField from '@/components/fields/InputField.vue';
import TextareaField from '@/components/fields/TextareaField.vue';
import BooleanField from '@/components/fields/BooleanField.vue';
import NumberField from '@/components/fields/NumberField.vue';
import DatePickerField from '@/components/fields/DatePickerField.vue';
import SelectField from '@/components/fields/SelectField.vue';
import RepeaterField from '@/components/fields/RepeaterField.vue';

export default {
  name: 'Field',
  components: {
    InputField,
    TextareaField,
    BooleanField,
    NumberField,
    DatePickerField,
    SelectField,
    RepeaterField,
  },
  props: {
    data: { type: Object },
  },
  methods: {
    addField(id) {
      this.$emit('addField', id);
    },
  },
  mounted() {
    console.log(this.data);
  },
};
</script>

<style lang="less" scoped>
.hastooltip {
  display: inline-block;
}

.checkbox-block {
  display: block;
}

.label-text {
  margin-bottom: 0.5rem;
  display: block;
}

.form-row {
  margin-bottom: 1rem;
}

</style>
