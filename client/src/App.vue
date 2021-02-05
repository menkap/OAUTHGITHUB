<template>
  <div id="app">
    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
    <Repository msg="Welcome to Your Vue.js App"/>
     <router-view></router-view>
  </div>
</template>

<script>
import Repository from './components/Repository.vue'

export default {
  name: 'App',
  components: {
    Repository
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
