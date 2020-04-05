<template lang="pug">
  div(class="mb-2")
    b-card(class="border-0 shadow-sm mb-6 bg-card")
      b-form-textarea(v-model="content" placeholder="Enter something..." id="textarea" rows=4 no-resize class="border-0 bg-card")
      div(class="ml-auto")
        b-button(:disabled="!content" v-on:click='sendPost' class="float-right")  Share!

</template>
<script>
import backendFunctions from "@/backendFunctions.js";

export default {
  name: "AddPostCard",
  data: function() {
    return {
      content: ""
    };
  },
  methods: {
    sendPost: function() {
      backendFunctions.sendPost(this.content);
      this.content = "";
    }
  }
};
</script>

<style lang="sass">
.profile-avatar
  width: 3em
  height: 3em
</style>