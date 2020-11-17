<template>
  <div>
    <h1 class="title">Login / Signup</h1>
    <LoginForm />
    <SignupForm />
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import LoginForm from "./../components/LoginForm.vue";
import SignupForm from "./../components/SignupForm.vue";

export default defineComponent({
  components: {
    LoginForm,
    SignupForm
  },
});
</script>