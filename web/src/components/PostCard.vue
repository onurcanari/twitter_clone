<template lang="pug">
div
  b-modal(v-bind:id="post.ID.toString()" class="modal-container" centered hide-footer hide-header)
    PostCard(v-bind:post="post")
  div(class="mb-2" @click="popupPost")
    b-card(class="border-0 shadow bg-card")
      span
        b-img(class="profile-avatar shadow-sm mb-1" rounded="circle" src="/assets/profile.jpg")
        span  @{{post.Username}}
      timeago(class="pl-3" :datetime="post.CreatedAt")
      b-card-text(class="font-weight-light") {{ post.Content }}
      div(class="d-flex justify-content-end mb-0 pb-0")
        div
          p(class="h5")
            b-icon(icon='reply-fill')
        div
          p(class="h5")
            b-icon(icon='heart-fill')
            span {{ post.LikeCount }}
</template>
<script>
export default {
  name: "PostCard",
  props: ["post"],
  methods: {
    popupPost: function() {
      this.$bvModal.show(this.post.ID.toString());
    }
  }
};
</script>

<style lang="sass">
svg
  color: white
.modal-body
  padding: 0 !important
  margin: 0 !important
</style>