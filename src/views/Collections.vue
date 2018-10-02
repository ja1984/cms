<template>
  <div class="page">
    <div class="container">

      <h1 class="page-title">Collections</h1>
      <table class="table collections">
        <thead>
          <tr>
            <th>Name</th>
            <th>Properties</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="collection in collections" :key="collection.key">
            <td>
              <router-link :to="{name: 'collection', params: {key: collection.key}}"> {{collection.key}}</router-link>
            </td>
            <td>{{collection.properties.length}}</td>
          </tr>
        </tbody>
      </table>
      <h3 class="sub-title">Add new collection</h3>
      <form @submit.prevent="createCollection">
        <div class="row row-bottom">
          <div class="column">
            <label>
              <span class="label-text">Name</span>
              <input type="text" v-model="key" placeholder="Translations">
            </label>
          </div>
          <div class="column column-wrap">
            <button type="submit" class="button button-primary">New collection </button>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
  name: 'Collections',
  data() {
    return {
      key: '',
    };
  },
  computed: {
    ...mapGetters({
      collections: 'collection/get',
    }),
  },
  methods: {
    createCollection() {
      this.$store.dispatch('collection/create', { key: this.key, properties: [] }).then(() => {
        this.key = '';
      });
    },
  },
};
</script>

<style lang="less" scoped>
</style>
