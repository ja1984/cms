<template>
  <header class="headerbar">
    <div class="container">

      <div class="row row-center-vertically">
        <div class="column column-wrap">
          <a href="/" class="logo">c<i></i>gCMS</a>
        </div>
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
@import (reference) '~@/styles/site.less';
@headerbarHeight: 7rem;

.headerbar {
  border-bottom: 0.1rem solid #e8e8e8;
  background: #fff;
  height: @headerbarHeight;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 500;

  &.mini {
    span {
      visibility: hidden;
      opacity: 0;
    }
  }
}

.container,
.row {
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

.logo {
  font-family: 'Josefin Sans', sans-serif;
  font-size: 4rem;
  color: #555;
  opacity: 1;

  i {
    font-size: 2.1rem;
    width: 2.1rem;
    height: 2.1rem;
    background: url('~@/assets/cog-logo.svg');
    display: inline-block;
    background-size: contain;
    margin-bottom: -0.2rem;
    background-repeat: no-repeat;
    margin-left: 0.1rem;
    transition: all ease 2.5s;
  }

  &:hover {
    color: #555;
    i {
      transform: rotate(360deg);
    }
  }
}
</style>
