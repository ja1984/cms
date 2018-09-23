<template>
  <div>
    <div class="row row-center-vertically field">
      <div class="column"><input type="text" ref="name" placeholder="Name" :value="value.name" @input="update(true)" /></div>
      <div class="column">
        <div v-show="value.slug.length > 0">
          <span class="tag" v-show="!editSlug" @click="editSlug = true">{{value.slug}}</span>
          <input type="text" v-show="editSlug" ref="slug" :value="value.slug" @input="update(false)">
        </div>
      </div>
      <div class="column">{{type}}</div>
      <div class="column column-wrap">
        <div class="field-tools">
          <!-- <button class="field-tools-button" @click="removeField(field.id)">
            <i class="far fa-trash-alt"></i>
          </button> -->
        </div>
      </div>
    </div>
    <div class="row row-center-vertically">
      <div class="column column-wrap">
        <checkbox v-model="value.required" :checked="value.required" ref="required" @input="update(false)">Required</checkbox>
      </div>
      <div class="column column-wrap">
        <input type="text" ref="tooltip" placeholder="Tooltip" :value="value.tooltip" @input="update(false)">
      </div>
    </div>
  </div>
</template>


<script>
import Checkbox from '@/components/CheckBox.vue';

export default {
  name: 'FieldType',
  components: {
    Checkbox,
  },
  props: {
    value: { type: Object },
    type: { type: String },
  },
  data() {
    return {
      editSlug: false,
    };
  },
  methods: {
    update(setSlug) {
      const slug = setSlug ? this.slugify(this.$refs.name.value) : this.$refs.slug.value;
      this.$emit('input', {
        name: this.$refs.name.value,
        required: this.value.required,
        tooltip: this.$refs.tooltip.value,
        slug,
      });
    },
    createId() {
      return Math.floor((1 + Math.random()) * 0x10000)
        .toString(16)
        .substring(1);
    },
    slugify(input) {
      return input.toString().toLowerCase()
        .replace(/\s+/g, '_')
        .replace(/\s+/g, '_')
        .replace(/[^\w\-]+/g, '')
        .replace(/\-\-+/g, '_')
        .replace(/^-+/, '')
        .replace(/-+$/, '');
    },
  },
};
</script>

<style lang="less" scoped>
@import (reference) '~@/styles/site.less';

.field-tools {
  opacity: 0;
  transition: all ease 0.3s;

  .field-tools-button {
    padding: 0.5rem;
    cursor: pointer;
  }
}

.field:hover .field-tools {
  opacity: 1;
}
</style>
