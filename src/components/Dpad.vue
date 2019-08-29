<template>
  <div class="dpad">
    <div class="dpad__upper-row">
      <Button label="UP" keyId="u" />
    </div>
    <div class="dpad__bottom-row">
      <Button label="LEFT" keyId="l" />
      <Button label="DOWN" keyId="d" />
      <Button label="RIGHT" keyId="r" />
    </div>
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
