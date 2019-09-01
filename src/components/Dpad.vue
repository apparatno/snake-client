<template>
  <div class="dpad">
    <div class="dpad__upper-row">
      <DpadButton
        keyId="u"
        :playerToken="this.$attrs.playerToken"
        :disabled="this.$attrs.disabled"
        ref="moveUpButton"
      />
    </div>
    <div class="dpad__bottom-row">
      <DpadButton
        keyId="l"
        :playerToken="this.$attrs.playerToken"
        :disabled="this.$attrs.disabled"
        ref="moveLeftButton"
      />
      <DpadButton
        keyId="d"
        :playerToken="this.$attrs.playerToken"
        :disabled="this.$attrs.disabled"
        ref="moveDownButton"
      />
      <DpadButton
        keyId="r"
        :playerToken="this.$attrs.playerToken"
        :disabled="this.$attrs.disabled"
        ref="moveRightButton"
      />
    </div>
  </div>
</template>

<script>
import DpadButton from "./DpadButton";

export default {
  name: "Dpad",
  components: {
    DpadButton
  },
  created() {
    window.addEventListener("keyup", this.onKey);
  },
  beforeDestroy() {
    window.removeEventListener("keyup", this.onKey);
  },
  methods: {
    onKey(key) {
      switch (key.code) {
        case "KeyW":
        case "ArrowUp":
          this.$refs.moveUpButton.onClick();
          break;
        case "KeyS":
        case "ArrowDown":
          this.$refs.moveDownButton.onClick();
          break;
        case "KeyA":
        case "ArrowLeft":
          this.$refs.moveLeftButton.onClick();
          break;
        case "KeyD":
        case "ArrowRight":
          this.$refs.moveRightButton.onClick();
          break;
        default:
          break;
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.dpad {
  display: flex;
  flex-direction: column;

  &__upper-row {
    display: flex;
    justify-content: center;
  }

  &__bottom-row {
    display: flex;
    justify-content: center;
  }
}
</style>
