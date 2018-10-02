<template>
  <div class="field-wrapper">
    <div class="row row-center-vertically field">
      <div class="column column-wrap">
        <i class="icon fa-fw" :class="icon"></i>
      </div>

      <div class="column"><input type="text" ref="name" placeholder="Name" :value="value.name" @input="update(true)" /></div>
      <div class="column" style="display: none;">
        <div v-show="value.slug.length > 0">
          <span class="tag" v-show="!editSlug" @click="editSlug = true">{{value.slug}}</span>
          <input type="text" v-show="editSlug" ref="slug" :value="value.slug" @input="update(false)">
        </div>
      </div>
      <div class="column">
        <input type="text" ref="tooltip" placeholder="Tooltip" :value="value.tooltip" @input="update(false)">
      </div>
      <div class="column column-wrap">
        <checkbox v-model="value.required" :checked="value.required" ref="required" @input="update(false)">Required</checkbox>
      </div>

      <div class="column column-wrap">
        <button class="button button-primary button-small" @click="newOption" v-if="type === 'select'">Add option</button>
        <button class="button button-primary button-small" @click="$emit('setAddTo')" v-if="type === 'repeater'">Add field</button>
      </div>
      <div class="column column-wrap">
        <div class="field-tools">
          <!-- <button class="field-tools-button" @click="removeField(field.id)">
            <i class="far fa-trash-alt"></i>
          </button> -->
        </div>
      </div>
    </div>

    <div class="child-input" v-for="(option, index) in options" :key="option.id">
      <div class="row row-no-gutter">
        <div class="column">
          <input type="text"  :placeholder="`Option ${index+1}`" @input="update(false)" v-model="option.value">
        </div>
        <div class="column column-wrap">
          <button :disabled="index === 0" class="button button-small button-error toolbar-button">
            <i class="far fa-trash-alt"></i>
          </button>
        </div>
      </div>
    </div>
    <div class="child-field" v-if="type === 'repeater'" v-for="field in value.childFields" :key="field.id">
      <field-type v-model="field.data" :type="field.type"></field-type>
    </div>
  </div>
</template>


<script>
import Checkbox from '@/components/CheckBox.vue';
import FieldType from '@/components/FieldType.vue';

export default {
  name: 'FieldType',
  components: {
    Checkbox,
    FieldType,
  },
  props: {
    value: { type: Object },
    type: { type: String },
    children: { type: Array, default: () => [] },
  },
  data() {
    return {
      editSlug: false,
      options: [],
      fields: [],
    };
  },
  mounted() {
    if (this.type === 'select') {
      this.newOption();
    }
  },
  computed: {
    icon() {
      switch (this.type) {
        case 'input':
          return 'fas fa-toggle-on';
        case 'textarea':
          return 'fas fa-code';
        case 'boolean':
          return 'fas fa-toggle-on';
        case 'number':
          return 'fas fa-hashtag';
        case 'date':
          return 'far fa-calendar-alt';
        case 'media':
          return 'far fa-image';
        case 'select':
          return 'far fa-list-alt';
        case 'repeater':
          return 'fas fa-redo';
        default:
          return '';
      }
    },
  },
  methods: {
    newOption() {
      this.options.push({
        id: this.createId(),
        value: '',
      });
      this.update(false);
    },
    update(setSlug) {
      const slug = setSlug ? this.slugify(this.$refs.name.value) : this.$refs.slug.value;
      this.$emit('input', {
        name: this.$refs.name.value,
        required: this.value.required,
        tooltip: this.$refs.tooltip.value,
        slug,
        options: this.options,
        childFields: this.value.childFields,
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

.card {
  margin-bottom: 1.5rem;
}

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

.child-field {
  margin-bottom: 0.1rem;
  padding-left: 1rem;
  border-left: 0.1rem dashed #e8e8e8;
  margin-left: 1rem;
}

.field-wrapper {
  border-bottom: 0.1rem solid #e8e8e8;
  padding-bottom: 1rem;
  margin-bottom: 1rem;

  .field-wrapper {
    border-bottom: none;
  }
}

.icon {
  font-size: 2rem;
}

.child-input {
  margin-top: 1rem;

  &:first-child {
    margin: 0;
  }
}

.toolbar-button {
  margin-left: 1rem;
}
</style>
