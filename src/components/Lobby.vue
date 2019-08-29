<template>
  <div class="lobby-container">
    <StartButton v-on:startedgame="startedGame" />
  </div>
</template>

<script>
import StartButton from "./StartButton.vue";

export default {
  name: "Lobby",
  components: {
    StartButton
  },
  methods: {
    startedGame(playerToken) {
      console.log("started game", playerToken);
      this.$emit("playgame", playerToken);
    }
  }
};
</script>

<style lang="scss" scoped>
</style>
