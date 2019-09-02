<template>
  <div class="start-button-container">
    <div class="start-button" type="button" role="button" v-on:click="onClick">
      <v-icon class="start-button__icon" name="play" />
    </div>
  </div>
</template>

<script>
import config from "../config";

export default {
  name: "StartButton",
  data() {
    return {};
  },
  methods: {
    onClick() {
      console.log("start button clicked");

      fetch(`${config.serverHost}${config.serverPort}/play`, {
        method: "POST"
      })
        .then(response => {
          return response.json();
        })
        .then(json => {
          console.log("started game");
          this.$emit("startedgame", json.playerToken);
        })
        .catch(err => {
          console.log("start game error", err);
        });
    }
  }
};
</script>

<style lang="scss" scoped>
.start-button-container {
  display: flex;
  justify-content: center;
  margin: 5px;
}
.start-button {
  display: flex;
  justify-content: center;
  background-color: #6a91ba;
  border-color: #809096;
  border-style: none;
  width: 125px;
  height: 75px;
  transition: all 0.1s ease-in-out;
  font-size: 12pt;
  cursor: pointer;
  box-shadow: 10px 10px 5px -3px rgba(0, 0, 0, 0.2);

  &:hover {
    transform: scale(1.05);
    background-color: #76a0cc;
    box-shadow: 10px 10px 5px 0px rgba(0, 0, 0, 0.1);
  }

  &__icon {
    width: 30px;
  }
}
</style>
