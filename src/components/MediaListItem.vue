<template>
  <div class="media-list-item">
    <div class="card" :class="{'selected': selectedFile && file.id === selectedFile.id}">
      <div class="card-body">
        <div class="preview-wrapper">
          <div class="image" :class="{'pending': !file.complete}" v-if="isImage" :style="{'backgroundImage': `url(${file.object.image.src})`}">
          </div>
          <div class="file" v-if="!isImage">
            <i class="far fa-fw" :class="fileClass"></i>
          </div>
          <i class="loader" v-if="!file.complete"></i>
        </div>
      </div>
      <!-- <img :src="file.object.image.src" class="image-preview" :class="{'pending': !file.complete && !file.failed}"> -->
      <div class="card-footer">
        {{file.object.fileName}}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MediaListItem',
  props: {
    file: {
      type: Object,
    },
  },
  methods: {
    bytesToSize(bytes) {
      const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
      if (bytes === 0) return '0 Byte';
      const i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)), 10);
      return `${Math.round(bytes / (1024 ** i), 2)} ${sizes[i]}`;
    },
    deleteFile() {
      // TODO: Maybe add so that we ceck if the folder is empty first.
      if (this.file.deleted) {
        this.$alertify.confirmWithTitle(this.$t('confirm.title.delete-file', 'Delete file?'), `${this.$t('confirm.subtitle.delete-file', 'Are you sure you want to delete "')} ${this.file.name}", this can not be undone?`, () => {
          this.$store.dispatch('media/DELETE_FILE', { id: this.file.id, permanently: true });
        }, () => { });
      } else {
        this.$alertify.confirmWithTitle(this.$t('confirm.title.move-to-trash', 'Move to trash?'), `${this.$t('confirm.subtitle.move-to-trash', 'Are you sure you want to move "')} ${this.file.name}" to the trash?`, () => {
          this.$store.dispatch('media/DELETE_FILE', { id: this.file.id, permanently: false });
        }, () => { });
      }
    },
    restoreFile() {
      // TODO: Maybe add so that we ceck if the folder is empty first.
      this.$store.dispatch('media/RESTORE_FILE', this.file.id);
    },
    insertUrl() {
      window.parent.postMessage({ type: 'media_select_link', imageUrl: this.file.permaLink }, '*');
    },
  },
  computed: {
    fileClass() {
      if (this.file.object.fileName.indexOf('.pdf') > -1) return 'fa-file-pdf';
      if (this.file.object.fileName.indexOf('.rar') > -1 || this.file.object.fileName.indexOf('.zip') > -1 || this.file.object.fileName.indexOf('.7zip') > -1) return 'fa-file-archive';
      if (this.file.object.type.indexOf('video/') > -1) return 'fa-file-video';
      if (this.file.object.type.indexOf('text/') > -1) return 'fa-file-alt';
      if (this.file.object.type.indexOf('audio/') > -1) return 'fa-file-audio';
      return 'fa-file';
    },
    fileSize() {
      return this.bytesToSize(this.file.object.size);
    },
    isImage() {
      return this.file.object.type.indexOf('image') > -1;
    },
  },
};
</script>

<style lang="less" scoped>
.preview-wrapper {
  position: relative;
  padding-top: 56.25%; /* 16:9 Aspect Ratio */
  background: url('~@/assets/background-transparent.gif');
}

.image,
.file {
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  background-size: contain;
  right: 0;
  background-position: center;
  background-repeat: no-repeat;

  &.pending {
    filter: grayscale(100%);
  }
}

.file {
  background: #fff;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 5rem;
}

.loader {
  position: absolute;
  top: .25rem;
  right: .25rem;
  background-color: #fff;
  border-radius: 50%;
  width: 3rem;
  height: 3rem;
}
</style>
