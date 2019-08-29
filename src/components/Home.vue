<template>
  <div class="page-container">
    <div class="header-container">
      <img src="../assets/apparat.png" alt="Apparat logo" width="300" />
      <h2>snake</h2>
    </div>
    <div class="content-container">
      <div v-if="isPlaying && playerToken !== ''" class="dpad-container">
        <Playing :playerToken="playerToken" />
      </div>
      <div v-else-if="isPlaying && playerToken === ''">
        <h2>another user is playing :(</h2>
      </div>
      <div v-else>
        <Lobby v-on:playgame="showGame" />
      </div>
    </div>
  </div>
</template>

<script>
import config from "../config";
import Lobby from "./Lobby.vue";
import Playing from "./Playing.vue";

export default {
  name: "Home",
  components: {
    Lobby,
    Playing
  },
  data() {
    return {
      playerToken: "",
      isPlaying: false
    };
  },
  mounted() {
    this.getGameState();
  },
  methods: {
    showGame(playerToken) {
      console.log("show game", playerToken);
      this.playerToken = playerToken;
      this.isPlaying = true;
    },
    getGameState() {
      fetch(`http://${config.serverHost}:${config.serverPort}/state`, {
        method: "GET"
      })
        .then(response => response.json())
        .then(json => {
          console.log("fetched initial game state", json);
          this.isPlaying = json.status === "playing";
        });
    }
  }
};
</script>

<style lang="scss" scoped>
</style>
