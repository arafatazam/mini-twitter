<template>
  <div>
    <Tweet></Tweet>
  </div>
</template>
<script lang="ts">
import {defineComponent} from 'vue'
import Tweet from './../components/Tweet.vue'

export default defineComponent({
  components: {
    Tweet
  }
})
</script>