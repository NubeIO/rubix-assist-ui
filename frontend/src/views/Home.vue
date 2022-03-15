<template>
  <v-app id="inspire">
    <div class="home">
      <img alt="Vue logo" src="../assets/logo.png" />
      <HelloWorld msg="Welcome to Your Vue.js + TypeScript App" />
      aa is: {{aa}}
      <button v-on:click="getLocale2">Greet</button>
    </div>
  </v-app>
</template>

<script lang="ts">
import Vue from "vue";
import HelloWorld from "../components/hello/HelloWorld.vue";

export default Vue.extend({
  name: "Home",
  data: () => ({
    aa: 2,
  }),
  components: {
    HelloWorld,
  },
  methods: {
    getLocale2() {
      this.aa++;
    },
  },
});
</script>
