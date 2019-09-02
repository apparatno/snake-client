<template>
  <div class="page-container">
    <div class="header-container">
      <img src="./assets/apparat.png" alt="Apparat logo" width="300" />
      <h1>snake</h1>
    </div>
    <div class="content-container">
      <div v-if="isPlaying && playerToken !== ''" class="dpad-container">
        <Playing :playerToken="playerToken" v-on:restartgame="restartGame" />
      </div>
      <div v-else-if="isPlaying && playerToken === ''" class="game-in-session">
        <h2>Another user is playing right now</h2>
      </div>
      <div v-else-if="!hasError">
        <Lobby v-on:playgame="showGame" />
      </div>
      <div v-else-if="hasError" class="error-container">
        <p>{{errorMessage}}</p>
      </div>
    </div>
  </div>
</template>

<script>
import config from "./config";
import Lobby from "./components/Lobby.vue";
import Playing from "./components/Playing.vue";

export default {
  name: "Home",
  components: {
    Lobby,
    Playing
  },
  data() {
    return {
      playerToken: "",
      isPlaying: false,
      hasError: false,
      errorMessage: ""
    };
  },
  mounted() {
    this.getGameState();
  },
  methods: {
    restartGame(playerToken) {
      console.log("restart game", playerToken);
      this.playerToken = playerToken;
      this.isPlaying = true;
    },
    showGame(playerToken) {
      console.log("show game", playerToken);
      this.playerToken = playerToken;
      this.isPlaying = true;
    },
    getGameState() {
      fetch(`${config.serverHost}${config.serverPort}/state`, {
        method: "GET"
      })
        .then(response => response.json())
        .then(json => {
          console.log("fetched initial game state", json);
          this.isPlaying = json.status === "playing";
        })
        .catch(err => {
          console.log("fetch initial game state error", err);
          this.hasError = true;

          if (err.toString().includes("NetworkError")) {
            this.errorMessage = "Error connecting to server";
          } else if (err.toString().includes("TypeError: Failed to fetch")) {
            this.errorMessage = "Server is not up";
          } else {
            this.errorMessage = err;
          }
        });
    }
  }
};
</script>

<style lang="scss" scoped>
.error-container {
  padding: 15px;
  background-color: #db5c5c;
  border-radius: 10px;
}
.game-in-session {
  padding: 15px;
  background-color: #5ca6db;
  border-radius: 10px;
}
</style>
