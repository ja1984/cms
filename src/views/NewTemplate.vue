<template>
  <div class="new-template page page-with-sidebar">
    <div class="container">
      <div class="row">
        <div class="column">
          <main>
            <div class="form-row">
              <label>
                Template name
                <input type="text" v-model="name">
              </label>
            </div>
            <div class="divider"></div>
            <strong>Fields</strong>
            <field-type v-for="field in fields" :key="field.id" v-model="field.data" :type="field.type"></field-type>
          </main>
        </div>
      </div>
    </div>

    <aside class="page-sidebar">
      <div class="page-sidebar-body">
        <ul>
          <li>
            <button class="toolbar-button" @click="addField('input')">
              <i class="fas fa-toggle-on fa-fw"></i> Text</button>
          </li>
          <li>
            <button class="toolbar-button" @click="addField('textarea')">
              <i class="fas fa-code fa-fw"></i> Textarea</button>
          </li>
          <li>
            <button class="toolbar-button" @click="addField('boolean')">
              <i class="fas fa-toggle-on fa-fw"></i> True / false</button>
          </li>
          <li>
            <button class="toolbar-button" @click="addField('number')">
              <i class="fas fa-hashtag"></i> Number</button>
          </li>
          <li>
            <button class="toolbar-button" @click="addField('date')">
              <i class="far fa-calendar-alt fa-fw"></i> Date</button>
          </li>
          <!-- <li>
                  <button class="toolbar-button" @click="addField('group')">
                    <i class="far fa-folder fa-fw"></i> Group</button>
                </li> -->
          <li>
            <button class="toolbar-button" @click="addField('media')">
              <i class="far fa-image fa-fw"></i> Media</button>
          </li>
          <li>
            <button class="toolbar-button" @click="addField('list')">
              <i class="far fa-image fa-fw"></i> List</button>
          </li>
          <li></li>
          <li></li>
        </ul>
      </div>
      <div class="page-sidebar-footer">
        <button @click="createTemplate" class="button button-success button-block" :disabled="!canCreate">Create template</button>

      </div>
    </aside>
  </div>
</template>

<script>
import FieldType from '@/components/FieldType.vue';

export default {
  name: 'NewTemplate',
  components: {
    FieldType,
  },
  data() {
    return {
      id: this.createId(),
      name: '',
      fields: [],
    };
  },
  computed: {
    canCreate() {
      if (this.fields.length < 1) return false;
      if (this.name.length === 0) return false;
      if (this.fields.some(x => x.data.name.length === 0)) return false;
      return true;
    },
  },
  methods: {
    addField(type) {
      this.fields.push({
        id: this.createId(),
        type,
        data: {
          name: '',
          required: false,
          tooltip: '',
          slug: '',
          options: [],
        },
      });
    },
    removeField(id) {
      const fieldIndex = this.fields.findIndex(x => x.id === id);
      console.log(fieldIndex, id);
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


li {

  i {
    font-size: 2rem;
    vertical-align: middle;
    margin-right: 0.5rem;
  }
}

.toolbar-button {
  border: none;
  padding: 0.75rem;
  cursor: pointer;
  transition: all ease 0.3s;
  // font-weight: 600;
  font-size: 1.5rem;
  display: block;
  width: 100%;
  text-align: left;
  background: transparent;

  &:active,
  &:focus {
    outline: none;
  }

  &:hover {
    background: #f9f9f9;
    color: #282d3a;
  }
}

.divider {
  width: 100%;
  height: 0.1rem;
  background: #ccc;
  margin-bottom: 1.5rem;
}
</style>
