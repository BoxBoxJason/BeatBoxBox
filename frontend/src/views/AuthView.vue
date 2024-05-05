<template>
  <div>
    <h2>Authentication</h2>
    <select v-model="authType">
      <option value="login">Login</option>
      <option value="register">Register</option>
    </select>
    <div v-if="authType === 'login'">
      <h3>Login</h3>
      <input type="text" v-model="loginData.username" placeholder="Username or Email">
      <input type="password" v-model="loginData.password" placeholder="Password">
      <button @click="login">Login</button>
    </div>
    <div v-else-if="authType === 'register'">
      <h3>Register</h3>
      <input type="text" v-model="registerData.username" placeholder="Username">
      <input type="email" v-model="registerData.email" placeholder="Email">
      <input type="password" v-model="registerData.password" placeholder="Password">
      <button @click="register">Register</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      authType: 'login',
      loginData: {
        username: '',
        password: ''
      },
      registerData: {
        username: '',
        email: '',
        password: ''
      }
    };
  },
  methods: {
    login() {
      console.log('Login:', this.loginData);
      // Add login logic here
    },
    register() {
      console.log('Register:', this.registerData);
      // Add register logic here
    }
  }
};
</script>
