<template>
  <div class="new-page container">
    <h1 class="page-title">New page</h1>
    <div class="row">
      <div class="column">
        <main>
          <div class="form-row">
            <div class="row row-center-vertically">
              <div class="column">
                <input type="text" placeholder="Page name" v-model="name">
              </div>
              <div class="column">
                <template v-if="slug.length > 2">
                  <span class="tag" v-if="!editSlug" @click="editSlug = true">{{slug}}</span>
                  <input type="text" v-model="slug" v-else>
                </template>
              </div>
            </div>
          </div>

          <ul class="tabs">
            <li class="tab" v-for="tab in tabs" v-bind:key="tab" :class="{'active': tab === selectedTab}" @click="selectedTab = tab">{{tab}}</li>
          </ul>

          <section v-show="selectedTab === 'Editor'">
            <div class="fields">
              <Field v-for="field in fields" :key="field.id" :data="field"></Field>
            </div>
          </section>
          <section v-show="selectedTab === 'Pretty'">
            <h1>{{name}}</h1>
            <p class="slug">
              <span class="dimmed">{{baseUrl}}</span>{{slug}}</p>
            <div class="fields">
              <div class="field-preview" v-for="field in fields" :key="field.id">
                <h4>{{field.name}}</h4>
                <div v-html="field.value"></div>
              </div>
            </div>
          </section>
          <section class="json-preview" v-show="selectedTab === 'JSON'">
            <pre v-html="pagePreview"></pre>
            <div class="preview-footer">
              Response size: ~{{pageSize}} bytes
            </div>
          </section>
        </main>
      </div>
      <div class="column column-wrap">
        <aside>

          <div class="">
            <div class="card-body">
              <div class="form-row">
                <label>
                  Template
                  <select v-model="selectedTemplate">
                    <option v-for="template in templates" :value="template" :key="template.id">{{template.name}}</option>
                  </select>
                </label>
              </div>
            </div>
            <div class="card-body">
            </div>
            <div class="card-footer">
              <div>
                <button @click="createPage" :disabled="!canCreate" class="button button-success button-block">Create page</button>
              </div>
            </div>
          </div>
        </aside>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import Field from '@/components/Field.vue';

export default {
  name: 'NewPage',
  components: {
    Field,
  },
  data() {
    return {
      selectedTemplate: null,
      id: this.createId(),
      name: '',
      slug: '',
      fields: [],
      tabs: ['Editor', 'Pretty', 'JSON'],
      selectedTab: 'Editor',
      editSlug: false,
    };
  },
  mounted() {
    const template = this.$store.getters['template/byId'](this.$route.params.id);
    this.selectedTemplate = template;
    // if (template) {
    //   this.setup(template);
    // }
  },
  computed: {
    ...mapGetters({
      templates: 'template/all',
    }),
    baseUrl() {
      return document.location.href;
    },
    pagePreview() {
      const fields = {};
      this.fields.forEach((field) => {
        fields[field.slug] = field.value;
      });
      const page = {
        name: this.name,
        slug: this.slug,
        fields,
      };

      return JSON.stringify(page, null, 2);
    },
    pageSize() {
      return this.pagePreview.length;
    },
    canCreate() {
      if (this.fields.length < 1) return false;
      if (this.name.length === 0) return false;
      if (this.fields.some(x => (x.required && x.value.length === 0))) return false;
      return true;
    },
  },
  methods: {
    setup(template) {
      this.fields = [];

      template.fields.forEach((field) => {
        console.log(field.type, field.type === 'boolean');
        const value = field.type === 'boolean' ? false : null;
        this.fields.push({
          id: field.id,
          name: field.data.name,
          type: field.type,
          slug: field.data.slug,
          value,
          required: field.data.required,
          tooltip: field.data.tooltip,
        });
      });
    },
    slugify(input) {
      return input.toString().toLowerCase()
        .replace(/\s+/g, '-')
        .replace(/[^\w\-]+/g, '')
        .replace(/\-\-+/g, '-')
        .replace(/^-+/, '')
        .replace(/-+$/, '');
    },
    createPage() {
      this.$store.dispatch('page/create', {
        id: this.id,
        name: this.name,
        slug: this.slug,
        fields: this.fields,
      });
    },
    createId() {
      return Math.floor((1 + Math.random()) * 0x10000)
        .toString(16)
        .substring(1);
    },
  },
  watch: {
    name(name) {
      this.slug = `/${this.slugify(name)}/`;
    },
    selectedTemplate(template) {
      this.setup(template);
      // this.fields = [];

      // template.fields.forEach((field) => {
      //   this.fields.push({
      //     id: field.id,
      //     name: field.data.name,
      //     type: field.type,
      //     slug: field.data.slug,
      //     value: '',
      //     required: field.data.required,
      //     tooltip: field.data.tooltip,
      //   });
      // });
    },
  },
};
</script>

<style lang="less" scoped>
main {
}
aside {
  width: 30rem;
}

.json-preview {
  display: flex;
  border-radius: 0.3rem;

  border: 0.1rem solid #dfdfdf;
  flex-direction: column;
  overflow: hidden;

  pre {
    flex: 1;
    max-height: 50rem;
    overflow-y: scroll;
    background: #fff;
    font-size: 1.5rem;
    padding: 0.8rem;
  }

  .preview-footer {
    background: transparent;
    border-top: 0.1rem solid #dfdfdf;
    padding: 0.24rem 0.72rem;
    font-size: 1.4rem;
    line-height: 1.7;
  }
}

.field-preview {
  padding: 1rem 0;
}
</style>
