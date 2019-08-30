<template>
  <div class="screen-container">
    <v-stage :config="configKonva">
      <v-layer>
        <v-rect v-bind:key="cell.id" v-for="cell in grid" :config="cell" />
      </v-layer>
    </v-stage>
  </div>
</template>

<script>
import uuidv4 from "uuid/v4";
import config from "../config";

export default {
  name: "Screen",
  data() {
    return {
      configKonva: {
        width: 600,
        height: 450
      },
      grid: []
    };
  },
  mounted() {
    this.interval = setInterval(() => this.updateGrid(), 500);
  },
  methods: {
    updateGrid() {
      console.log("update grid");

      fetch(`http://${config.serverHost}:${config.serverPort}/screen`, {
        method: "GET"
      })
        .then(response => {
          return response.text();
        })
        .then(state => {
          this.parseGrid(state);
        });
    },
    parseGrid(state) {
      const list = [];
      var sublist = [];

      for (let i = 0; i < state.length; i++) {
        const c = state[i];

        if (i > 0 && i % 20 === 0) {
          list.push(sublist);
          sublist = [];
          sublist.push(c);
        } else {
          sublist.push(c);
        }
      }

      const width = 30;
      const height = 30;

      for (let i = 0; i < list.length; i++) {
        const sublist = list[i];

        for (let j = 0; j < sublist.length; j++) {
          const sym = sublist[j];

          var color = "";
          switch (sym) {
            case "0":
              color = "#c5ccd1";
              break;
            case "1":
              color = "#5199db";
              break;
            case "2":
              color = "#8acf7e";
              break;
            default:
              color = "#cbcf7e";
              break;
          }

          const rect = {
            id: uuidv4(),
            y: i * height,
            x: j * width,
            width,
            height,
            fill: color
          };

          this.grid.push(rect);
        }
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.screen-container {
  display: flex;
  justify-content: center;
}
</style>
