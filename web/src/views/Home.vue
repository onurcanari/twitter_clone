<template lang="pug">
  b-row
    b-col(cols="3" class="pt-3")
      div(class="sticky-top")
        ProfileCard(v-if="isUserSigned")
        b-card(v-else class="bg-card")
          b-link(to="/login")
            span SIGN IN NOW
    b-col(cols="9" class="pt-3")
      AddPostCard(v-if="isUserSigned")
      Posts

</template>

<script>
import Posts from "@/components/Posts.vue";
import ProfileCard from "@/components/ProfileCard.vue";
import AddPostCard from "@/components/AddPostCard.vue";
import backendFunctions from "@/backendFunctions.js";

export default {
  data() {
    return {
      body: "",
      isUserSigned: false
    };
  },
  components: {
    ProfileCard,
    Posts,
    AddPostCard
  },
  created: function() {
    backendFunctions
      .isSigned()
      .then(() => {
        this.isUserSigned = true;
      })
      .catch(() => {
        this.isUserSigned = false;
      });
  }
};
</script>

<style lang="sass">
.bg-card
  background-color: #232331 !important

.profile-avatar
  width: 3em
  height: 3em
body, button
  background-color: #353552 !important
</style>