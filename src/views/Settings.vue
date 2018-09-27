<template>
  <div class="page">
  <div class="container">
    <h1 class="page-title">Settings</h1>

    <ul class="tabs">
      <li class="tab" v-for="tab in tabs" v-bind:key="tab" :class="{'active': tab === selectedTab}" @click="selectedTab = tab">{{tab}}</li>
    </ul>

    <section v-if="selectedTab === 'Localization'">
      <h3 class="sub-title">Add new locale</h3>
      <form @submit.prevent="saveLanguage">
        <div class="form-row">
          <label>
            <span>Language</span>
            <input type="text" required placeholder="English" v-model="language">
          </label>
        </div>
        <div class="form-row">
          <label>
            <span>API Slug</span>
            <input type="text" placeholder="en" required v-model="slug">
          </label>
        </div>
        <div>
          <button class="button button-success" type="submit">Add language</button>
        </div>
      </form>

      <div class="site-languages">
        <h3  class="sub-title">Your locales</h3>
        <div class="site-language row row-center-vertically" v-for="language in languages" :key="language.slug">
          <div class="column">{{language.language}}</div>
          <div class="column">{{language.slug}}</div>
          <div class="column">{{language.isDefault}}</div>
          <div class="column column-wrap">
            <button class="button button-error" @click="deleteLanguage(language.slug)">
              <i class="far fa-trash-alt"></i>
            </button>
          </div>
        </div>
      </div>
    </section>
  </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
  name: 'Templates',
  computed: {
    ...mapGetters({
      languages: 'language/all',
    }),
  },
  data() {
    return {
      tabs: ['Localization'],
      selectedTab: 'Localization',
      language: '',
      slug: '',
    };
  },
  methods: {
    saveLanguage() {
      this.$store.dispatch('language/create', { language: this.language, slug: this.slug });
    },
    deleteLanguage(slug) {
      this.$store.dispatch('language/delete', { slug });
    },
  },
};
</script>

<style lang="less" scoped>
section {
  padding: 1rem 0;
}
</style>
