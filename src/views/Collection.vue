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
      </template>
    </div>
  </div>
</template>

<script>
export default {
  name: 'collection',
  computed: {
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
