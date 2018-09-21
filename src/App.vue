<template>
  <div id="app">
    <header-bar></header-bar>
    <main class="content-wrapper">
      <router-view/>
    </main>
  </div>
</template>
<script>
import HeaderBar from '@/components/HeaderBar.vue';

export default {
  name: 'App',
  components: {
    HeaderBar,
  },
};
</script>
<style lang="less">
@import '~@/styles/site.less';

.content-wrapper {
  padding-top: 10rem;
}
</style>
