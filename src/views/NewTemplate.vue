<template>
  <div class="new-template page page-with-sidebar">
    <div class="container">
      <div class="row">
        <div class="column">
          <main>
            <div class="form-row">
              <label>
                <span class="label-text">Template name</span>
                <input type="text" v-model="name">
              </label>
            </div>
            <div class="divider"></div>
            <strong>Fields</strong>
            <field-type v-for="field in fields" @setAddTo="setAddTo(field)" :disableSetAddTo="addTo !== null" @removeField="removeField" :id="field.id" :key="field.id" v-model="field.data" :type="field.type"></field-type>
          </main>
        </div>
      </div>
    </div>
    <aside class="page-sidebar">
      <div class="page-sidebar-body">
        <div class="panel panel-warning" v-if="addTo !== null">
          Add to
          <strong>{{addTo.data.name}}</strong>
          <a href="javscript:void(0);" class="cancel-addto" @click="addTo = null">Cancel</a>
        </div>
        <field-selection :disableRepeater="addTo !== null" @addField="addField"></field-selection>
      </div>
      <div class="page-sidebar-footer">
        <button @click="createTemplate" class="button button-success button-block" :disabled="!canCreate">Create template</button>
      </div>
    </aside>
  </div>
</template>

<script>
import FieldType from '@/components/FieldType.vue';
import FieldSelection from '@/components/FieldSelection.vue';

export default {
  name: 'NewTemplate',
  components: {
    FieldType,
    FieldSelection,
  },
  data() {
    return {
      id: this.createId(),
      name: '',
      fields: [],
      addTo: null,
    };
  },
  computed: {
    canCreate() {
      if (this.fields.length < 1) return false;
      if (this.name.length === 0) return false;
      if (this.fields.some(x => x.data.name.length === 0)) return false;
      return true;
    },
    groupedFields() {
      const fields = [];
      this.fields.forEach((field) => {
        if (field.parentId === null) {
          fields.push({
            field,
            children: []          });
          return;
        }
        const oldField = fields.find(x => x.field.id === field.parentId);

        if (oldField) {
          oldField.children.push(field);
        }
      });
      return fields;
    },
  },
  methods: {
    setAddTo(field) {
      this.addTo = field;
    },
    addField(type) {
      const newField = {
        id: this.createId(),
        type,
        data: {
          name: '',
          required: false,
          tooltip: '',
          slug: '',
          options: [],
          childFields: [],
        },
      };

      if (this.addTo === null) {
        this.fields.push(newField);
      } else {
        console.log(this.addTo);
        this.addTo.data.childFields.push(newField);
      }
    },
    removeField(id) {
      const fieldIndex = this.fields.findIndex(x => x.id === id);
      if (fieldIndex !== null) {
        this.fields.splice(fieldIndex, 1);
      }
    },
    createTemplate() {
      this.$store.dispatch('template/create', { id: this.id, name: this.name, fields: this.fields });
    },
    createId() {
      return Math.floor((1 + Math.random()) * 0x10000)
        .toString(16)
        .substring(1);
    },
  },
};
</script>

<style lang="less" scoped>
@import (reference) '~@/styles/site.less';

.editor {
  display: flex;
}



.divider {
  width: 100%;
  height: 0.1rem;
  background: #ccc;
  margin-bottom: 1.5rem;
}

.cancel-addto {
  float: right;
}
</style>
