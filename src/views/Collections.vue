<template>
  <div class="page">
  <div class="container">
    <h1 class="page-title">Collections</h1>

    <div class="collections">
      <div class="collection" v-for="collection in collections" :key="collection.key">
        <router-link :to="{name: 'collection', params: {key: collection.key}}"> {{collection.key}}</router-link>
      </div>
    </div>
  </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
  name: 'Collections',
  computed: {
    ...mapGetters({
      collections: 'collection/get',
    }),
  },
};
</script>

<style lang="less" scoped>

</style>
