<template>
  <div id="app" @dragenter="toggleDropZone" @drop="toggleDropZone" @dragleave="toggleDropZone">
    <side-bar></side-bar>
    <main class="content-wrapper">
      <router-view/>
    </main>
    <drop-zone :show="showDropZone"></drop-zone>
  </div>

</template>
<script>
import SideBar from '@/components/SideBar.vue';
import DropZone from '@/components/DropZone.vue';

export default {
  name: 'App',
  data() {
    return {
      showDropZone: false,
    };
  },
  methods: {
    toggleDropZone() {
      this.showDropZone = !this.showDropZone;
    },
  },
  components: {
    SideBar,
    DropZone,
  },
};
</script>
<style lang="less">
@import '~@/styles/site.less';

.content-wrapper {
  padding-top: 2rem;
}
</style>
