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

const updateInterval = 500;
const numHorizontalCells = 20;
const numVerticalCells = 15;
const cellWidth = 30;
const cellHeight = 30;
const colors = {
  empty: "#c5ccd1",
  snake: "#5199db",
  fruit: "#8acf7e",
  missing: "#cbcf7e"
};

export default {
  name: "Screen",
  data() {
    return {
      configKonva: {
        width: numHorizontalCells * cellWidth,
        height: numVerticalCells * cellHeight
      },
      grid: []
    };
  },
  mounted() {
    this.interval = setInterval(() => this.updateGrid(), updateInterval);
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
        .then(gameState => {
          this.parseGrid(gameState);
        });
    },
    parseGrid(gameState) {
      const rows = [];
      var row = [];

      for (let i = 0; i < gameState.length; i++) {
        const pixelState = gameState[i];

        if (i > 0 && i % numHorizontalCells === 0) {
          rows.push(row);
          row = [];
          row.push(pixelState);
        } else {
          row.push(pixelState);
        }
      }

      for (let i = 0; i < rows.length; i++) {
        const row = rows[i];

        for (let j = 0; j < row.length; j++) {
          const pixelState = row[j];

          var color = "";
          switch (pixelState) {
            case "0":
              color = colors.empty;
              break;
            case "1":
              color = colors.snake;
              break;
            case "2":
              color = colors.fruit;
              break;
            default:
              color = colors.missing;
              break;
          }

          const rect = {
            id: uuidv4(),
            y: i * cellHeight,
            x: j * cellWidth,
            width: cellWidth,
            height: cellHeight,
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
