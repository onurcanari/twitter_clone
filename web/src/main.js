import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import router from './router'
import { BootstrapVue, LayoutPlugin, IconsPlugin } from "bootstrap-vue"
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import VueTimeago from 'vue-timeago'

Vue.config.productionTip = false
Vue.use(BootstrapVue)
Vue.use(LayoutPlugin)
Vue.use(IconsPlugin)
Vue.use(VueTimeago, {
  name: 'Timeago',
  locale: 'en',
})
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
