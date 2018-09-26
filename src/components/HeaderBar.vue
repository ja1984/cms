<template>
  <aside class="nav-sidebar">

    <router-link to="/" class="logo">c<i></i>gCMS</router-link>
    <div class="site-selection">
      <div class="selected-site">cogCMS<i class="fas fa-caret-down"></i></div>
    </div>
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
      <li>
        <router-link to="/settings">Settings</router-link>
      </li>
    </nav>
<div class="account-wrapper">
  <i class="fas fa-sliders-h"></i>
</div>
  </aside>
</template>

<script>
export default {
  name: 'HeaderBar',
};
</script>

<style lang="less" scoped>
@import (reference) '~@/styles/site.less';
@headerbarHeight: 24rem;

.nav-sidebar {
  background: #282d3a;
  width: @headerbarHeight;
  position: fixed;
  top: 0;
  left: 0;
  bottom: 0;
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
  display: block;
  padding: 1rem;
}

a {
  font-weight: 600;
  font-size: 1.4rem;
  opacity: 0.5;
  color: #fff;
  display: inline-block;

  &:hover {
    opacity: 1;
    color: #fff;
  }
}

.logo {
  font-family: 'Josefin Sans', sans-serif;
  font-size: 4rem;
  color: #fff;
  opacity: 1;
  padding: 2rem;
  display: block;
  width: 100%;
  text-align: center;
  background: rgba(0, 0, 0, 0.1);

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
    color: #fff;
    i {
      transform: rotate(360deg);
    }
  }
}

.account-wrapper {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 1rem;
  color: #fff;
    cursor: pointer;
  transition: all ease .3s;
  &:hover {
    background: rgba(255,255,255,.1);
  }
}

.site-selection {
  border-bottom: .1rem solid rgba(0, 0, 0, 0.1);
}

.selected-site {
  color: #fff;
  font-weight: 600;
  padding: 1rem;
  position: relative;
  cursor: pointer;
  transition: all ease .3s;
  &:hover {
    background: rgba(255,255,255,.1);
  }

  i {
    position: absolute;
    right: 1rem;
    top: 1rem;
  }
}
</style>
