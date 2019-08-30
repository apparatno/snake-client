<template>
  <div class="button-container">
    <div class="button" type="button" role="button" v-on:click="onClick">
      <v-icon class="button__icon" :name="iconName" />
    </div>
  </div>
</template>

<script>
import config from "../config";

export default {
  name: "DpadButton",
  props: ["label", "keyId"],
  data() {
    return {
      iconName: ""
    };
  },
  mounted() {
    this.setIconByKeyId();
  },
  methods: {
    setIconByKeyId() {
      switch (this.keyId) {
        case "u":
          this.iconName = "arrow-up";
          break;
        case "d":
          this.iconName = "arrow-down";
          break;
        case "l":
          this.iconName = "arrow-left";
          break;
        case "r":
          this.iconName = "arrow-right";
          break;
        default:
          this.iconName = "x";
          break;
      }
    },
    onClick() {
      console.log(`move ${this.getDirection(this.keyId)} clicked`);

      const formBody = `keyPressed=${this.keyId}&playerToken=${this.$attrs.playerToken}`;

      fetch(`http://${config.serverHost}:${config.serverPort}/action`, {
        method: "POST",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
          Accept: "application/json"
        },
        body: formBody
      })
        .then(response => {
          return response.text();
        })
        .then(() => {
          console.log(`moved ${this.getDirection(this.keyId)}`);
        })
        .catch(err => {
          console.log("move error", err);
        });
    },
    getDirection(keyId) {
      switch (keyId) {
        case "u":
          return "up";
        case "d":
          return "down";
        case "l":
          return "left";
        case "r":
          return "right";
        default:
          return "UNKNOWN";
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.button-container {
  margin: 5px;
}
.button {
  display: flex;
  justify-content: center;
  background-color: #87a662;
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
    background-color: #a5cc76;
    box-shadow: 10px 10px 5px 0px rgba(0, 0, 0, 0.1);
  }

  &__icon {
    width: 30px;
  }
}
</style>
