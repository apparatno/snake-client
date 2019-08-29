<template>
  <div class="dpad">
    <Button label="UP" />
    <Button label="DOWN" />
    <Button label="LEFT" />
    <Button label="RIGHT" />
  </div>
</template>

<script>
import Button from "./Button";

export default {
  name: "Dpad",
  components: {
    Button
  }
};
</script>

<style lang="scss" scoped>
</style>
