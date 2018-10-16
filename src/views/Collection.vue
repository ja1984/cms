<template>
  <div class="page">

    <div class="container">
      <template v-if="collection">
        <h1 class="page-title">{{collection.key}}</h1>
        <div class="collection">
          <div class="collection-properties" v-for="property in collection.properties" :key="property.key">
            {{property.key}}
            <div class="property-values">
              <div class="property-value" v-for="(value, propertyName) in property.value" :key="propertyName">
                {{propertyName}} : {{value}}
              </div>
            </div>
          </div>
        </div>
        <div class="new-key">
          <div class="">
            <label>
              <span class="label-text">Key</span>
              <input type="text" v-model="name">
            </label>
          </div>

          <div class="row">
            <div class="column no-padding-bottom">
              <span class="label">Slug</span>
            </div>
            <div class="column no-padding-bottom">
              <span class="label">Value</span>
            </div>
          </div>
          <div class="row" v-for="field in fields" :key="field.id">
            <div class="column no-padding-top">
              <input type="text" v-model="field.slug" :placeholder="field.slug">
            </div>
            <div class="column no-padding-top">
              <input type="text" v-model="field.value">
            </div>
          </div>
          <div class="row">
            <div class="column">
              <button class="button button-primary" @click="addField">Add value</button>
            </div>
            <div class="column text-right">
              <button class="button button-success" :disabled="!canSave" @click="saveKey">Save key</button>
            </div>
          </div>

          <!-- <div class="form-row" v-for="language in languages" :key="language.slug">
            <label>
              <span class="label-text">{{language.language}}</span>
                <input type="text">
            </label>
          </div> -->
        </div>
      </template>
    </div>
  </div>

</template>

<script>
import { mapGetters } from 'vuex';
import { createId } from '@/scripts/utils';


export default {
  name: 'collection',
  data() {
    return {
      name: '',
      fields: [],
    };
  },
  mounted() {
    this.addFields();
  },
  methods: {
    addFields() {
      this.fields = [];
      this.fields.push(...this.languages.map((x) => {
        return {
          id: createId(),
          language: x.language,
          slug: x.slug,
          value: '',
        };
      }));
    },
    addField() {
      this.fields.push({
        language: '',
        slug: '',
        value: '',
      });
    },
    saveKey() {
      const data = {
        key: this.name,
        value: {},
      };

      this.fields.forEach((x) => {
        data.value[x.slug] = x.value;
      });

      this.$store.dispatch('collection/addKey', { collection: this.$route.params.key, data }).then(() => {
        this.name = '';
        this.addFields();
      });
    },
  },
  computed: {
    canSave() {
      if (this.name.length === 0) return false;
      if (this.fields.some(x => x.slug.length < 1 || x.value.length < 1)) return false;
      return true;
    },
    ...mapGetters({
      languages: 'language/all',
    }),
    collection() {
      return this.$store.getters['collection/bySlug'](this.$route.params.key);
    },
  },
};
</script>

<style lang="less" scoped>
.collection div {
  padding: 0.5rem;
  padding-left: 1rem;
}

.no-padding-bottom {
  padding-bottom: 0;
}
.no-padding-top {
  padding-top: 0;
}

.label {
  font-size: 1.4rem;
  font-weight: 600;
  margin-bottom: 0.75rem;
  display: block;
}
</style>
