<template>
  <div>
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
      this.$emit("playgame", playerToken);
    }
  }
};
</script>
