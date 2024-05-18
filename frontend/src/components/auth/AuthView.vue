<template>
  <div class="login">
    <h2>Login</h2>
    <form @submit.prevent="login">
      <div>
        <label for="username_or_email">Username or Email:</label>
        <input type="text" id="username_or_email" v-model="username_or_email" required>
      </div>
      <div>
        <label for="password">Password:</label>
        <input type="password" id="password" v-model="password" required>
      </div>
      <button type="submit">Login</button>
    </form>
    <p v-if="error" class="error">{{ error }}</p>
  </div>
  <div class="register">
    <h2>Register</h2>
    <form @submit.prevent="register">
      <div>
        <label for="username">Username:</label>
        <input type="text" id="username" v-model="username" required>
      </div>
      <div>
        <label for="email">Email:</label>
        <input type="email" id="email" v-model="email" required>
      </div>
      <div>
        <label for="password">Password:</label>
        <input type="password" id="password" v-model="password" required>
      </div>
      <div>
        <label for="password_confirm">Password Confirmation:</label>
        <input type="password" id="password_confirm" v-model="password_confirm" required>
      </div>
      <button type="submit">Register</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username_or_email: '',
      password: '',
      username: '',
      email: '',
      password_confirm: '',
      error: ''
    };
  },
  methods: {
    async login() {
      try {
        const response = await fetch('/api/users/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
          },
          body: new URLSearchParams({
            username_or_email: this.username_or_email,
            password: this.password
          })
        });
        if (!response.ok) {
          throw new Error('Invalid Username / Password');
        }
        const data = await response.json();
        // Handle successful login, e.g., store user data, redirect, etc.
        console.log(data);
      } catch (err) {
        this.error = err.message;
      }
    },
    async register() {
      try {
        const response = await fetch('/api/users/register', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
          },
          body: new URLSearchParams({
            username: this.username,
            email: this.email,
            password: this.password,
            password_confirm: this.password_confirm
          })
        });
        if (!response.ok) {
          throw new Error('Invalid Registration');
        }
        const data = await response.json();
        // Handle successful registration, e.g., store user data, redirect, etc.
        console.log(data);
      } catch (err) {
        this.error = err.message;
      }
    }
  }
};
</script>

<style scoped>
.login {
  max-width: 400px;
  margin: 0 auto;
  padding: 1rem;
  border: 1px solid #ccc;
  border-radius: 5px;
}
.error {
  color: red;
}
</style>
