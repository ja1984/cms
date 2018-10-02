<template>
  <div class="page">
    <div class="container">
      <h1 class="page-title">Pages</h1>
      <div class="card ">
      <table class="table collections">
        <thead>
          <tr>
            <th>Name</th>
            <th>Status</th>
            <th>Created</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="page in pages" :key="page.id" :class="{'status-unpublished' : !page.published}">
            <td>
              <router-link :to=" {name: 'page', params: {key: page.id}} ">{{page.name}}</router-link>
              <p class="dimmed">{{page.slug}}</p>
            </td>
            <td>{{page.published}}</td>
            <td>{{page.created}}</td>
          </tr>
        </tbody>
      </table>
      <!-- <div class="column ">
        <div class="page-item " v-for="page in pages " :key="page.id ">
          {{ page.name }}
        </div>
      </div> -->
      
        <div class="card-header row ">
          <div class="column ">
            <span class="card-title ">Pages</span>
          </div>

          <div class="column column-wrap ">
            <router-link to="/pages/create " class="button button-primary ">New page</router-link>
          </div>
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

.status-unpublished {
  opacity: .5;
}
</style>
