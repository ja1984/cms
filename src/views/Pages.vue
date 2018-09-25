<template>
  <div class="container">
    <h1 class="page-title">Pages</h1>
            <div class="column">
          <div class="page-item" v-for="page in pages" :key="page.id">
            {{ page.name }}
          </div>
        </div>
    <div class="card">
      <div class="card-header row">
        <div class="column">
          <span class="card-title">Pages</span>
        </div>

        <div class="column column-wrap">
          <router-link to="/pages/create" class="button button-primary">New page</router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Templates',
  mounted() {
    this.$store.dispatch('page/list');
  },
  computed: {
    pages() {
      return this.$store.getters['page/pages'];
    },
  },
};
</script>

<style lang="less" scoped>

</style>
