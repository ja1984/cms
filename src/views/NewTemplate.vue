<template>
  <div class="new-template">
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
            <div class="form-row" v-for="field in fields" :key="field.id">
              <div class="row row-center-vertically">
                <div class="column"><input type="text" v-model="field.name" /></div>
                <div class="column">{{field.type}}</div>
              </div>
            </div>
          </main>
        </div>
        <div class="column column-wrap">
          <aside>
            <div class="new-field-toolbar">
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
                    <i class="fas fa-sort-numeric-up"></i> Number</button>
                </li>
                <li>
                  <button class="toolbar-button" @click="addField('date')">
                    <i class="fas fa-calendar-alt fa-fw"></i> Date</button>
                </li>
                <li>
                  <button class="toolbar-button" @click="addField('group')">
                    <i class="far fa-folder fa-fw"></i> Group</button>
                </li>
                <li>
                  <button class="toolbar-button" @click="addField('media')">
                    <i class="far fa-image fa-fw"></i> Media</button>
                </li>
                <li></li>
                <li></li>
              </ul>
            </div>
            <div class="toolbar">
              <button @click="createTemplate" class="button button-success button-block">Save template</button>
            </div>
          </aside>
        </div>

      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'NewTemplate',
  data() {
    return {
      id: this.createId(),
      name: '',
      fields: [],
    };
  },
  methods: {
    addField(type) {
      this.fields.push({
        id: this.createId(),
        name: '',
        type,
      });
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
.editor {
  display: flex;
}

aside {
  width: 25rem;
}

li {
  padding: 0.25rem 0;

  i {
    font-size: 2rem;
    vertical-align: middle;
    margin-right: 0.5rem;
  }
}

.toolbar-button {
  border: none;
  padding: 1rem;
  cursor: pointer;
  transition: all ease 0.3s;
  font-weight: 600;
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
    background: #fff;
  }
}

.divider {
  width: 100%;
  height: 0.1rem;
  background: #ccc;
  margin-bottom: 1.5rem;
}
</style>
