import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from "@/store";
import axios from "axios";
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import Multiselect from "vue-multiselect";
import "vue-multiselect/dist/vue-multiselect.min.css";
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)
Vue.component("multiselect", Multiselect);

Vue.config.productionTip = false
axios.defaults.baseURL = "/server";

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
