<template>
  <div class="container">
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
};
</script>

<style lang="less" scoped>

</style>
