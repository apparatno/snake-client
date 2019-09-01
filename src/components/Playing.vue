<template>
  <div class="playing">
    <div class="screen-container">
      <Screen v-on:gameover="gameOver" v-on:restartgame="restartGame" />
    </div>
    <div class="dpad-container">
      <Dpad :playerToken="this.$attrs.playerToken" :disabled="this.isGameOver" />
    </div>
  </div>
</template>

<script>
import Dpad from "./Dpad.vue";
import Screen from "./Screen.vue";

export default {
  name: "Playing",
  components: {
    Dpad,
    Screen
  },
  data() {
    return {
      isGameOver: false
    };
  },
  methods: {
    gameOver() {
      this.isGameOver = true;
    },
    restartGame(playerToken) {
      this.isGameOver = false;
      this.$emit("restartgame", playerToken);
    }
  }
};
</script>

<style lang="scss" scoped>
.screen-container {
  margin-bottom: 20px;
}
.dpad-container {
  margin-bottom: 40px;
}
</style>
