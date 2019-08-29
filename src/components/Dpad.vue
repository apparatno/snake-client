<template>
  <div class="dpad">
    <div class="dpad__upper-row">
      <DpadButton label="UP" keyId="u" :playerToken="this.$attrs.playerToken" />
    </div>
    <div class="dpad__bottom-row">
      <DpadButton label="LEFT" keyId="l" :playerToken="this.$attrs.playerToken" />
      <DpadButton label="DOWN" keyId="d" :playerToken="this.$attrs.playerToken" />
      <DpadButton label="RIGHT" keyId="r" :playerToken="this.$attrs.playerToken" />
    </div>
  </div>
</template>

<script>
import DpadButton from "./DpadButton";

export default {
  name: "Dpad",
  components: {
    DpadButton
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
