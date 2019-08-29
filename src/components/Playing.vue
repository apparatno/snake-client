<template>
  <div class="dpad-container">
    <Dpad :playerToken="this.$attrs.playerToken" />
  </div>
</template>

<script>
import Dpad from "./Dpad.vue";

export default {
  name: "Playing",
  components: {
    Dpad
  }
};
</script>

<style lang="scss" scoped>
</style>
