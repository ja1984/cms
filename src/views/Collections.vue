<template>
  <div class="container">
    <h1 class="page-title">Collections</h1>
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
  name: 'Collections',
};
</script>

<style lang="less" scoped>

</style>
