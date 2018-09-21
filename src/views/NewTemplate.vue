<template>
  <div class="new-template">
    <div class="container">
      <div class="card">
        <div class="card-header">
          <span class="card-title">Create new template</span>
        </div>
        <div class="card-body">
          <div class="row">
            <div class="column">
              <main>
                <div>
                  <label>
                    Template name
                    <input type="text" v-model="name">
                  </label>
                </div>
                <div class="form-row" v-for="field in fields" :key="field.id">
                  <div class="row">
                    <div class="column"><input type="text" v-model="field.name" /></div>
                    <div class="column">{{field.type}}</div>
                  </div>
                </div>
                <div class="toolbar">
                  <button @click="createTemplate" class="button button-primary">Save template</button>
                </div>
              </main>
            </div>
            <div class="column column-wrap">
              <aside>
                <div class="toolbar">
                  <button @click="addNewTextbox">New textbox</button>
                  <button @click="addNewTextarea">New textarea</button>
                </div>
              </aside>
            </div>

          </div>
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
    addNewTextbox() {
      this.fields.push({
        id: this.createId(),
        name: '',
        type: 'input',
      });
    },
    addNewTextarea() {
      this.fields.push({
        id: this.createId(),
        name: '',
        type: 'textarea',
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
</style>
