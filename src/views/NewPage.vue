<template>
  <div class="page page-with-sidebar page-with-sidebar-slide" :class="{'show': showAddNewField}">
    <header>
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

        <div class="colomn column-wrap">
          <div class="button-wrapper button-add-new-field">
            <button @click="showAddNewField = !showAddNewField" class="button button-primary button-block button-save">Add new field
              <i class="fas fa-caret-down"></i>
            </button>
            <div class="popup" :class="{'show': showAddNewField}">
              <div class="card">
                <div class="card-body">
                  <field-selection :disableRepeater="true" @addField="addNewField"></field-selection>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="column column-wrap">

          <div class="button-wrapper">
            <button @click="createPage(false)" :disabled="!canCreate" class="button button-success button-block button-save">Save page</button>
            <button @click="showPopup = !showPopup" :disabled="!canCreate" class="button button-success button-more">
              <i class="fas fa-caret-down"></i>
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
    </header>
    <div class="new-page container">
      <h1 class="page-title">New page</h1>
      <div class="row">
        <div class="column">
          <main>
            <ul class="tabs">
              <li class="tab" v-for="tab in tabs" v-bind:key="tab" :class="{'active': tab === selectedTab}" @click="selectedTab = tab">{{tab}}</li>
            </ul>

            <section v-show="selectedTab === 'Editor'">
              <div class="fields">
                <div class="field" v-for="field in groupedFields" :key="field.field.id">
                  <Field :data="field.field"></Field>
                  <div class="child-field" v-for="childField in field.children" :key="childField.id">
                    <div class="row row-no-gutters">
                      <div class="column">
                        <Field :data="childField"></Field>
                      </div>
                      <div class="column column-wrap">
                        <button class="button button-error button-small">
                          <i class="far fa-trash-alt"></i>
                        </button>
                      </div>
                    </div>
                  </div>
                  <button class="button button-primary button-small" @click="addField(field.field.id)" v-if="field.field.type === 'repeater'">Add field</button>
                </div>
              </div>
              <div class="addNewField" v-if="newTemplateField !== null">
                <field-type :canBeDeleted="false" :id="newTemplateField.id" :key="newTemplateField.id" v-model="newTemplateField.data" :type="newTemplateField.type"></field-type>
                <button class="button button-primary" @click="saveNewField">Add field</button>
                <button class="button button-link margin margin-left-1" @click="newTemplateField = null">Cancel</button>
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
        <!-- <div class="column column-wrap">
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

                </div>
              </div>
            </div>
          </aside>
        </div> -->
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import { createId } from '@/scripts/utils';

import Field from '@/components/Field.vue';
import { createPage } from '@/api/page';
import FieldSelection from '@/components/FieldSelection.vue';
import FieldType from '@/components/FieldType.vue';


export default {
  name: 'NewPage',
  components: {
    Field,
    FieldSelection,
    FieldType,
  },
  data() {
    return {
      selectedTemplate: null,
      id: null,
      name: '',
      slug: '',
      fields: [],
      childFields: [],
      tabs: ['Editor', 'Pretty', 'JSON'],
      selectedTab: 'Editor',
      editSlug: false,
      publish: false,
      showPopup: false,
      showAddNewField: false,
      newTemplateField: null,
    };
  },
  mounted() {
    if (this.$route.name === 'editpage') {
      const page = this.$store.getters['page/byId'](this.$route.params.id);
      if (!page) return;
      this.name = page.name;
      this.id = page.id;
      this.slug = page.slug;
      page.fields.forEach((field) => {
        this.fields.push(this.createField({
          type: field.type,
          data: {
            name: field.name,
            type: field.type,
            required: field.required,
            tooltip: field.tooltip,
            options: field.options,
            slug: field.slug,
          },
        }, field.value, [], field.id));
      });
      // this.selectedTemplate =
      // EDIT!!!;
    } else {
      const template = this.$store.getters['template/byId'](this.$route.params.id);
      this.selectedTemplate = template;
    }
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
    groupedFields() {
      const fields = [];
      this.fields.forEach((field) => {
        const children = [];
        const childGroups = this.childFields.filter(x => x.id === field.id);
        childGroups.forEach((group) => {
          children.push(...group.fields);
        });
        fields.push({
          field,
          children,
        });
      });
      return fields;
    },
    pagePreview() {
      const fields = {};
      this.fields.forEach((field) => {
        if (field.type === 'repeater') {
          const childGroups = this.childFields.filter(x => x.id === field.id);
          const data = [];
          childGroups.forEach((group) => {
            const values = {};
            group.fields.forEach((field) => {
              values[field.slug] = field.value;
            });
            data.push(values);
          });
          fields[field.slug] = data;
        } else {
          fields[field.slug] = field.value;
        }
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
    saveNewField() {
      const value = this.newTemplateField.type === 'boolean' ? false : null;
      console.log(this.newTemplateField, value);
      this.fields.push(this.createField(this.newTemplateField, value, []));
      this.newTemplateField = null;
    },
    addNewField(type) {
      this.showAddNewField = false;
      this.newTemplateField = {
        id: createId(),
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
    },
    addField(id) {
      console.log('addfield', id);
      const template = this.selectedTemplate.fields.find(x => x.id === id);
      const repeater = this.fields.find(x => x.id === id);

      const childFields = template.data.childFields.map((childField) => {
        const childValue = childField.type === 'boolean' ? false : null;
        return this.createField(childField, childValue, [], createId());
      });
      this.childFields.push({
        id: repeater.id,
        fields: childFields,
      });

      // repeater.childFields.push(...childFields);
    },
    setup(template) {
      this.fields = [];

      template.fields.forEach((field) => {
        const value = field.type === 'boolean' ? false : null;

        const childFields = field.data.childFields.map((childField) => {
          const childValue = childField.type === 'boolean' ? false : null;
          return this.createField(childField, childValue, [], createId());
        });

        this.childFields.push({
          id: field.id,
          fields: childFields,
        });


        this.fields.push(this.createField(field, value, childFields));
      });
    },
    createField(field, value, childFields, id) {
      return {
        id: id || field.id,
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
      const page = {
        id: this.id,
        name: this.name,
        slug: this.slug,
        fields: this.fields,
        updated: new Date(),
        template: this.$route.params.id,
        published,
      };

      if (this.id !== null) {
        this.$store.dispatch('page/update', page);
      } else {
        page.id = createId();
        page.created = new Date();
        this.$store.dispatch('page/create', page);
      }

      // const page = {
      //   id: this.id,
      //   name: this.name,
      //   slug: this.slug,
      //   fields: this.fields,
      // };
      // createPage(page).then((response) => {
      //   if (response.status === 201) {
      //     this.$store.dispatch('page/create', {
      //       id: this.id,
      //       name: this.name,
      //       slug: this.slug,
      //       fields: this.fields,
      //     });
      //   }
      // });
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
.page {
  padding-top: 7rem;
}
aside {
  width: 30rem;
  top: 7rem;
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

.field {
  border-bottom: 0.1rem solid #e8e8e8;
  margin-bottom: 1rem;
  padding-bottom: 1rem;

  &:last-child {
    border: 0;
  }
}

.child-field {
  margin-bottom: 0.1rem;
  padding-left: 1rem;
  border-left: 0.1rem dashed #e8e8e8;
}

.button-wrapper {
  width: 25rem;
  display: flex;
  position: relative;
  &.button-add-new-field {
    width: auto;

    i {
      margin-left: 1rem;
    }
  }

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
  bottom: -1.5rem;
  z-index: 15;
  transform: translateY(100%);

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
    &:before {
      content: '';
      position: absolute;
      border: 0.8rem solid #f3f3f3;
      border-top-color: transparent;
      border-left-color: transparent;
      border-right-color: transparent;
      left: 50%;
      transform: translateX(-50%) translateY(-100%);
      top: 0;
    }
    &:after {
      content: '';
      position: absolute;
      border: 0.7rem solid #fff;
      border-top-color: transparent;
      border-left-color: transparent;
      border-right-color: transparent;
      left: 50%;
      transform: translateX(-50%) translateY(-100%);
      top: 0;
    }
  }
}

header {
  position: fixed;
  top: 0;
  left: 24rem;
  z-index: 10;
  right: 0;
  background: #fff;
  border-bottom: 0.1rem solid #e8e8e8;
  padding: 0.5rem 1.5rem;
}
</style>
