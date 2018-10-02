<template>
  <aside class="nav-sidebar">

    <router-link to="/" class="logo">c<i></i>gCMS</router-link>
    <div class="site-selection">

      <div class="selected-site">
        <p class="small">Selected site</p>
        cogCMS
        <i class="fas fa-caret-down"></i>
      </div>
    </div>
    <nav>
      <li>
        <router-link to="/templates">
          <i class="far fa-file-alt fa-fw"></i> Templates</router-link>
      </li>
      <li>
        <router-link to="/pages">
          <i class="far fa-copy fa-fw"></i> Pages</router-link>
      </li>
      <li>
        <router-link to="/media">
          <i class="far fa-file-image fa-fw"></i> Media</router-link>
      </li>
      <li>
        <router-link to="/collections">
          <i class="fas fa-box-open fa-fw"></i> Collections</router-link>
      </li>
      <li>
        <router-link to="/settings">
          <i class="fas fa-cogs fa-fw"></i> Settings</router-link>
      </li>
    </nav>
    <div class="account-wrapper">
      <div class="row row-no-gutter row-center-vertically">
        <div class="column column-wrap">
          <img src="https://lh4.googleusercontent.com/-0R5Sbwj-Vug/AAAAAAAAAAI/AAAAAAAAD7w/e-ELKyQda1c/photo.jpg" class="profile-image">
        </div>
        <div class="column">
          <strong class="profile-name">Jonathan Andersson</strong>
        </div>
      </div>
    </div>
  </aside>
</template>

<script>
export default {
  name: 'SideBar',
};
</script>

<style lang="less" scoped>
@import (reference) '~@/styles/site.less';
@sidebarWidth: 24rem;

.nav-sidebar {
  background: #282d3a;
  width: @sidebarWidth;
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

  i {
    margin-right: 0.2rem;
    font-size: 1.8rem;
    vertical-align: text-top;
  }

  &:hover {
    background: rgba(255, 255, 255, 0.1);
  }
}

a {
  font-weight: 600;
  font-size: 1.4rem;
  padding: 1rem;

  opacity: 0.5;
  color: #fff;
  display: block;

  &.router-link-active {
    opacity: .8;
  }

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
  transition: all ease 0.3s;
  &:hover {
    background: rgba(255, 255, 255, 0.1);
  }
}

.site-selection {
  // border-bottom: 0.1rem solid rgba(0, 0, 0, 0.1);
}

.selected-site {
  color: #fff;
  font-weight: 600;
  padding: 1rem;
  position: relative;
  cursor: pointer;
  transition: all ease 0.3s;
  &:hover {
    background: rgba(255, 255, 255, 0.1);
  }

  i {
    position: absolute;
    right: 1rem;
    top: 50%;
    transform: translateY(-50%);
  }
}

.small {
  margin-bottom: 0.2rem;
  text-transform: uppercase;
  font-size: 1rem;
  font-weight: normal;
  color: rgba(255, 255, 255, 0.3);
}

.profile-image {
  border-radius: 50%;
  border: 0.2rem solid #fff;
  width: 3rem;
  margin-right: 1rem;
  display: block;
}

.profile-name {
  font-size: 1.3rem;
}
</style>
