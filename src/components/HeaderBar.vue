<template>
  <header class="headerbar">
    <div class="container">

      <div class="row row-center-vertically">
        <div class="column"></div>
        <div class="column">
          <nav>
            <li>
              <router-link to="/templates">Templates</router-link>
            </li>
            <li>
              <router-link to="/pages">Pages</router-link>
            </li>
            <li>
              <router-link to="/media">Media</router-link>
            </li>
            <li>
              <router-link to="/collections">Collections</router-link>
            </li>
          </nav>
        </div>
        <div class="column"></div>
      </div>

      <!-- <div class="page-logo">
      <img src="/static/cloud-logo.svg">
    </div>
    <nav class="menu">
      <ul>
        <li>
          <router-link to="/">
            <clr-icon shape="network-globe" size="22"></clr-icon>
            <span>Translations</span>
          </router-link>
        </li>
        <li>
          <router-link to="/">
            <clr-icon shape="flag" size="22"></clr-icon>
            <span>Feature flags</span>
          </router-link>
        </li>
      </ul>
    </nav> -->
    </div>
  </header>
</template>

<script>
export default {
  name: 'HeaderBar',
};
</script>

<style lang="less" scoped>
@headerbarHeight: 7rem;

.headerbar {
  border-bottom: 0.1rem solid #e8e8e8;
  background: #fff;
  height: @headerbarHeight;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;

  &.mini {
    span {
      visibility: hidden;
      opacity: 0;
    }
  }
}

.container, .row {
  height: 100%;
}

li {
  display: inline-block;
  padding: 0 0.75rem;
}

a {
  font-weight: 600;
  font-size: 1.4rem;
  opacity: 0.5;
  color: #333;
  display: inline-block;

  &:hover {
    opacity: 1;
    color: #333;
  }
}
</style>
