import axi from "axios";
const axios = axi.create({ baseURL: "http://localhost:8888"});

export default {
  sendPost(data) {
    return axios
      .post("/addPost", {
        content: data
      })
  },
  signin: function (username, password) {
    return axios
      .post("/signin", {
        Username: username,
        Password: password
      })
  },
  getPosts(offset) {
    return axios
      .post("/getPosts", {
        offset: offset
      })
  },
  isSigned() {
    return axios
      .get("/isSigned")
  },
  getUserDetails(username) {
    return axios
      .post("/getUserDetails", {
        username: username
      })
  },
  createSocket() {
    console.log("socket created!");
    return new WebSocket('ws://localhost:8888/livePosts');
  }

}