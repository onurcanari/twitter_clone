<template lang="pug">
  div
    PostCard(v-for="post in posts" v-bind:key="post.ID" v-bind:post="post")
    
</template>
<script>
import backendFunctions from "@/backendFunctions.js";

import PostCard from "@/components/PostCard.vue";
export default {
  name: "Posts",
  data: function() {
    return {
      posts: [],
      selectedPost: Object,
      offset: 0,
      toastCount: 0
    };
  },
  components: {
    PostCard
  },
  created: function() {
    backendFunctions.getPosts(this.offset).then(response => {
      this.posts = response.data;
    });
    backendFunctions.createSocket().onmessage = this.addToPosts;
  },
  mounted: function() {
    this.scroll();
  },
  methods: {
    scroll() {
      window.onscroll = () => {
        let bottomOfWindow =
          Math.max(
            window.pageYOffset,
            document.documentElement.scrollTop,
            document.body.scrollTop
          ) +
            window.innerHeight ===
          document.documentElement.offsetHeight;

        if (bottomOfWindow) {
          this.offset += 10;
          backendFunctions.getPosts(this.offset).then(response => {
            this.posts.push(...response.data);
          });
        }
      };
    },
    addToPosts(event) {
      var parsedJson = JSON.parse(event.data);
      this.popToast(parsedJson.Username);
      this.posts.unshift(parsedJson);
    },
    popToast(username) {
      const h = this.$createElement;
      this.toastCount++;
      const vNodesMsg = h("p", { class: ["text-center", "mb-0"] }, [
        h("b-spinner", { props: { type: "grow", small: true } }),
        " A new post from ",
        h("strong", ` ${username}`),
        h("b-spinner", { props: { type: "grow", small: true } })
      ]);
      const vNodesTitle = h(
        "div",
        {
          class: [
            "d-flex",
            "flex-grow-1",
            "align-items-baseline",
            "text-dark",
            "mr-2"
          ]
        },
        [
          h("strong", { class: ["mr-2", "text-dark"] }, "1 New Post"),
          h("small", { class: "ml-auto text-italics text-dark" }, "now")
        ]
      );
      this.$bvToast.toast([vNodesMsg], {
        title: [vNodesTitle],
        solid: true,
        variant: "dark"
      });
    },
    select: function(event) {
      console.log(event); // returns 'foo'
    }
  }
};
</script>
<style>
</style>