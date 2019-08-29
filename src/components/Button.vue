<template>
  <div class="button-container">
    <input class="button" type="button" :value="label" v-on:click="onClick" />
    {{ dump }}
  </div>
</template>

<script>
export default {
  name: "Button",
  props: ["label", "keyId"],
  data() {
    return {
      dump: null
    };
  },
  methods: {
    onClick() {
      console.log(`${this.label} button clicked`);

      fetch("https://rickandmortyapi.com/api/character", { method: "GET" })
        .then(response => {
          return response.json();
        })
        .then(json => {
          this.dump = json.results[0];
        });
    }
  }
};
</script>

<style lang="scss" scoped>
</style>
