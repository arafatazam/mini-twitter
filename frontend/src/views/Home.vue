<template>
  <div>
    <h1 class="title">Home</h1>
    <NewTweet />
    <Tweet></Tweet>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import Tweet from "./../components/Tweet.vue";
import NewTweet from "./../components/NewTweet.vue";

export default defineComponent({
  components: {
    Tweet,
    NewTweet,
  },
});
</script>