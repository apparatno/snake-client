<template>
  <div class="dpad-container">
    <Screen />
    <Dpad :playerToken="this.$attrs.playerToken" />
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
  }
};
</script>

<style lang="scss" scoped>
</style>
