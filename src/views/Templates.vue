<template>
  <div class="container">
    <h1 class="page-title">Templates</h1>
    <div v-for="template in templates" :key="template.id">
      <router-link :to="{name: 'createpagefromtemplate', params: {id: template.id}}">
        <div>
          {{template.name}} {{template.fields.length}}

        </div>

      </router-link>
    </div>
    <div class="card">
      <div class="card-header row">
        <div class="column">
          <span class="card-title">Templates</span>
        </div>
        <div class="column column-wrap">
          <router-link to="/templates/create" class="button button-primary">New template</router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

import { mapGetters } from 'vuex';

export default {
  name: 'Templates',
  computed: {
    ...mapGetters({
      templates: 'template/all',
    }),
  },
};
</script>

<style lang="less" scoped>
</style>
