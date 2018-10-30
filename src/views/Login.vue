<template>
  <div class="login-page">
    <div class="login-form">
      <div class="card">
        <div class="card-body">
          <h1 class="logo">c
            <i></i>gCMS</h1>
          <div class="form-row">
            <label>
              <span class="label-text">Email</span>
              <input type="text" v-model="userName">
            </label>
          </div>
          <div class="form-row">
            <label>
              <span class="label-text">Password</span>
              <input type="password" v-model="password">
            </label>
          </div>
          <div class="text-center">
            <button class="button button-primary" @click="login">Login</button>
            <button class="button button-primary" @click="loginWithGoogle">Googe</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>

import { login, validate } from '@/api/auth';

export default {
  name: 'Login',
  data() {
    return {
      userName: '',
      password: '',
    };
  },
  methods: {
    login() {
      login({ userName: this.userName, password: this.password }).then((response) => {
        console.log(response);
      }).catch((error) => {
        console.log(error);
      });
    },
    loginWithGoogle() {
      this.$gAuth.signIn((user) => {
        validate(user.getAuthResponse().id_token).then((res) => {
          console.log(res);
        });
        console.log(user.getAuthResponse().id_token);
      }, (error) => {
        console.log(error);
      });
    },
  },
};
</script>
<style lang="less" scoped>
.login-page {
  width: 100%;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

.login-form {
  width: 90%;
  max-width: 40rem;
}

.form-row {
  margin-bottom: 1rem;
}

.logo {
  font-family: 'Josefin Sans', sans-serif;
  font-size: 7rem;
  color: #282d3a;
  opacity: 1;
  padding: 2rem;
  display: block;
  width: 100%;
  text-align: center;

  i {
    font-size: 2.1rem;
    width: 3.3rem;
    height: 3.3rem;
    background: url('~@/assets/cog-logo-dark.svg');
    display: inline-block;
    background-size: contain;
    margin-bottom: -0.2rem;
    background-repeat: no-repeat;
    margin-left: 0.1rem;
    transition: all ease 2.5s;
  }

  &:hover {
    i {
      transform: rotate(360deg);
    }
  }
}
</style>
