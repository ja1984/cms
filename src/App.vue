<template>
  <div id="app">
    <side-bar></side-bar>
    <main class="content-wrapper">
      <router-view/>
    </main>
  </div>
</template>
<script>
import SideBar from '@/components/SideBar.vue';

export default {
  name: 'App',
  components: {
    SideBar,
  },
};
</script>
<style lang="less">
@import '~@/styles/site.less';

</style>
