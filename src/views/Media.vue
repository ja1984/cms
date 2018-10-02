<template>
  <div class="page page-with-sidebar page-with-sidebar-slide" :class="{'show': selectedFile !== null}">
    <div class="container">
      <h1 class="page-title">Media library</h1>
      <div class="files">
        <div class="file" v-for="file in media" :key="file.id">
          <div class="card" @click="selectFile(file)" :class="{'selected': selectedFile && file.id === selectedFile.id}">
            <img :src="file.url">
            <div class="card-footer">
              {{file.filename}}
            </div>
          </div>
        </div>
      </div>
    </div>
    <aside class="page-sidebar page-sidebar-wide">
      <div class="page-sidebar-body page-sidebar-body-with-padding" v-if="selectedFile">
        <img :src="selectedFile.url">

        <h4>{{selectedFile.filename}}</h4>
      </div>
    </aside>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
  name: 'Media',
  data() {
    return {
      selectedFile: null,
    };
  },
  mounted() {
    window.addEventListener('keydown', this.detectEsc);
  },
  destroyed() {
    window.removeEventListener('keydown', this.detectEsc);
  },
  methods: {
    detectEsc(event) {
      console.log('event', event);
      const evt = event || window.event;
      if (evt.keyCode === 27) {
        this.selectedFile = null;
      }
    },
    selectFile(file) {
      if (this.selectedFile === file) {
        this.selectedFile = null;
        return;
      }
      this.selectedFile = file;
    },
  },
  computed: {
    ...mapGetters({
      media: 'media/all',
    }),
  },
};
</script>

<style lang="less" scoped>
@import (reference) '~@/styles/site.less';

.files {
  display: flex;
  flex-wrap: wrap;
  margin-left: -1rem;
  margin-right: -1rem;
}

.file {
  padding: 1rem;
  flex: 0 0 50%;

  @media (min-width: 768px) {
    flex: 0 0 (100%/3);
  }
  @media (min-width: 1024px) {
    flex: 0 0 25%;
  }
  @media (min-width: 1200px) {
    flex: 0 0 20%;
  }

  img {
    display: block;
    max-width: 100%;
  }
}

.card {
  cursor: pointer;
  transition: all ease 0.3s;
  &.selected {
    box-shadow: 0 0 0.4rem 0 @primary;
  }
}

img {
  width: 100%;
}
</style>
