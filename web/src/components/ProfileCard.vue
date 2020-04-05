<template lang="pug">
b-card(class="border-0 shadow-sm bg-card")
    div(class="d-flex justify-content-center")
      div
        b-img(class="my-profile-avatar shadow p-1 mb-4" rounded="circle" center src="/assets/profile.jpg")
        div(class="mb-0" align="center") {{ fullname }}
        div(class="mt-0 grey-text" align="center") @{{ username }}
    div(class="d-flex justify-content-around")
        .detail
            p(class="mb-0") Followers
            div(align="center")
                P {{followers}}
        .v-line
        .detail
            p(class="mb-0") Follows
            div(align="center")
                P {{ follows }}
</template>

<script>
import backendFunctions from "@/backendFunctions.js";

export default {
  name: "ProfileCard",
  props: {
    Username: {}
  },
  data: function() {
    return {
      errors: [],
      fullname: "",
      posts: 0,
      followers: 0,
      follows: 0,
      username: this.Username
    };
  },
  created: function() {
    backendFunctions.getUserDetails(this.username).then(response => {
      this.fullname = response.data.Fullname;
      this.posts = response.data.Posts;
      this.followers = response.data.Followers;
      this.follows = response.data.Follows;
      this.username = response.data.Username;
    });
  }
};
</script>

<style lang="sass">
.my-profile-avatar
  width: 7em
  height: 7em
.v-line
  border-left: 1px solid #75757E
  height: 2em
  margin: auto 0
p, span, label, div, textarea
  color: white !important
time, .grey-text
  color: #75757E !important


</style>