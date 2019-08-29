<template>
  <div class="page-container">
    <div class="header-container">
      <img src="../assets/apparat.png" alt="Apparat logo" width="300" />
      <h2>Snakes!</h2>
    </div>
    <div class="content-container">
      <div class="dpad-container">
        <Dpad />
      </div>
    </div>
  </div>
</template>

<script>
import Dpad from "./Dpad.vue";

export default {
  name: "Home",
  components: {
    Dpad
  }
};
</script>

<style lang="scss" scoped>
</style>
