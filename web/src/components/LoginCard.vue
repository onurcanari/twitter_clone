<template lang="pug">
  b-card(class="shadow-lg mx-auto border-0 login-container align-middle bg-card")
    b-form(@submit="checkForm")
      p(v-if="errors.length")
          b-alert(show variant="danger" class="black-text" v-for="error in errors" :key="error") {{ error }}
      b-form-group(id="username-group" label="Username or email" label-for="username")
        b-form-input(id="username" v-model="username" type="text" name="username")
      b-form-group(id="password-group" label="Password" label-for="password")
        b-input(id="password" v-model="password" type="password" name="password" aria-describedby="password-help-block")
      b-button(block :disabled="!password || !username" type="submit" class="float-right" v-on:click="checkForm($event)") Sign in
</template>

<script>
import backendFunctions from "@/backendFunctions.js";

export default {
  name: "LoginForm",
  data: () => {
    return { errors: [], username: null, password: null };
  },
  methods: {
    checkForm: function(e) {
      this.errors = [];
      e.preventDefault();
      if (this.username && this.password) {
        backendFunctions
          .signin(this.username, this.password)
          .then(() => {
            this.$router.push({ name: "Home" });
          })
          .catch(() => {
            this.errors.push("Username or password is wrong.");
          });
        return true;
      } else {
        this.errors.push("Boş alan olamaz.");
      }
    }
  }
};
</script>

<style lang="sass">
.black-text
  color: black !important
.login-container
  position: relative
  max-width: 310px
</style>