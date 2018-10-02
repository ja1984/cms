<template>
  <div class="page">
    <header>
      
    </header>
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
                <div v-for="field in fields" :key="field.id">
                  <Field :data="field"></Field>
                  <div class="child-field" v-for="childField in field.childFields" :key="childField.id">
                    <Field :data="childField"></Field>
                  </div>
                </div>
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
              <div class="json-wrapper">
                <pre v-html="htmlEncode(pagePreview)"></pre>
              </div>
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
                  <div class="button-wrapper">
                    <button @click="createPage(false)" :disabled="!canCreate" class="button button-success button-block button-save">Save page</button>
                    <button @click="showPopup = !showPopup" :disabled="!canCreate" class="button button-success button-more">
                      <i class="fas fa-caret-up"></i>
                    </button>

                    <div class="popup" :class="{'show': showPopup}">
                      <div class="card">
                        <div class="card-body">
                          <button @click="createPage(true)" :disabled="!canCreate" class="button button-primary button-block button-save">Save and publish</button>
                          <button @click="createPage(true)" :disabled="!canCreate" class="button button-warning button-block button-save">Save as draft</button>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </aside>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import Field from '@/components/Field.vue';
import { createPage } from '@/api/page';

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
      publish: false,
      showPopup: false,
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
        const value = field.type === 'boolean' ? false : null;

        const childFields = field.data.childFields.map((childField) => {
          const childValue = childField.type === 'boolean' ? false : null;
          return this.createField(childField, childValue, []);
        });

        this.fields.push(this.createField(field, value, childFields));
      });
    },
    createField(field, value, childFields) {
      console.log(field, value, childFields);
      return {
        id: field.id,
        name: field.data.name,
        type: field.type,
        slug: field.data.slug,
        value,
        required: field.data.required,
        tooltip: field.data.tooltip,
        options: field.data.options || [],
        childFields,
      };
    },
    htmlEncode(input) {
      const text = input || '';
      return text.toString()
        .replace(/&/g, '&amp;')
        .replace(/"/g, '&quot;')
        .replace(/</g, '&lt;')
        .replace(/>/g, '&gt;');
    },
    slugify(input) {
      return input.toString().toLowerCase()
        .replace(/\s+/g, '-')
        .replace(/[^\w\-]+/g, '')
        .replace(/\-\-+/g, '-')
        .replace(/^-+/, '')
        .replace(/-+$/, '');
    },
    createPage(published) {
      this.$store.dispatch('page/create', {
        id: this.id,
        name: this.name,
        slug: this.slug,
        fields: this.fields,
        created: new Date(),
        published,
      });

      const page = {
        id: this.id,
        name: this.name,
        slug: this.slug,
        fields: this.fields,
      };
      createPage(page).then((response) => {
        if (response.status === 201) {
          this.$store.dispatch('page/create', {
            id: this.id,
            name: this.name,
            slug: this.slug,
            fields: this.fields,
          });
        }
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
@import (reference) '~@/styles/site.less';
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

  .json-wrapper {
    overflow: hidden;
  }

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

.child-field {
  margin-bottom: 0.1rem;
  padding-left: 1rem;
  border-left: 0.3rem solid #aaa;
}

.button-wrapper {
  display: flex;
  position: relative;
  .button-save {
    border-top-right-radius: 0;
    border-bottom-right-radius: 0;
  }
  .button-more {
    padding: 1.28rem 1.4rem;
    border-top-left-radius: 0;
    border-bottom-left-radius: 0;
    border-left: 0.1rem solid darken(@success, 5%);
  }
}

.popup {
  width: 100%;
  position: absolute;
  top: -1.5rem;
  transform: translateY(-100%);

  opacity: 0;
  visibility: hidden;

  &.show {
    visibility: visible;
    opacity: 1;
  }

  .button {
    margin-top: 1rem;

    &:first-child {
      margin: 0;
    }
  }

  .card {
    &:after {
      content: '';
      position: absolute;
      border: .7rem solid #fff;
      border-bottom-color: transparent;
      border-left-color: transparent;
      border-right-color: transparent;
      left: 50%;
      transform: translateX(-50%);
    }
  }
}
</style>
