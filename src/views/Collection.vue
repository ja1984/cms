<template>
  <div class="collection">
    <div class="container">
      <template v-if="collection">
        <h1 class="page-title">{{collection.key}}</h1>
        <div class="collection-properties" v-for="property in collection.properties" :key="property.key">
          {{property.key}}
          <div class="property-values">
            <div class="property-value" v-for="(value, propertyName) in property.value" :key="propertyName">
              {{propertyName}} : {{value}}
            </div>
          </div>
        </div>
        <div class="new-key">
          <div class="form-row">
            <label>
              <span class="label-text">Key</span>
              <input type="text" v-model="name">
            </label>
          </div>

          <div class="row">
            <div class="column" v-for="field in fields" :key="field.slug">
              <label>
                <span class="label-text">{{field.language}}</span>
                <input type="text" v-model="field.value">
              </label>
            </div>
          </div>
          <div>
            <button class="button button-success" @click="saveKey">Save key</button>
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

export default {
  name: 'collection',
  data() {
    return {
      name: '',
      fields: [],
    };
  },
  mounted() {
    this.fields.push(...this.languages.map((x) => {
      return {
        language: x.language,
        slug: x.slug,
        value: '',
      };
    }));
  },
  methods: {
    saveKey() {
      const data = {
        key: this.name,
        value: {},
      };

      this.fields.forEach((x) => {
        data.value[x.slug] = x.value;
      });

      this.$store.dispatch('collection/addKey', { collection: this.$route.params.key, data });
      console.log(data);
    },
  },
  computed: {
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
</style>
