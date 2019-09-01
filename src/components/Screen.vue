<template>
  <div class="screen">
    <v-stage v-if="!hasError && !isGameOver" :config="configKonva">
      <v-layer>
        <v-rect v-bind:key="cell.id" v-for="cell in grid" :config="cell" />
      </v-layer>
    </v-stage>
    <div v-if="isGameOver" class="game-over">
      <h2>Game Over</h2>
      <StartButton v-on:startedgame="restartedGame" />
    </div>
    <div v-if="hasError && errorMessage !== ''" class="error-message">
      <h2>{{errorMessage}}</h2>
    </div>
  </div>
</template>

<script>
import uuidv4 from "uuid/v4";
import config from "../config";
import StartButton from "./StartButton";

const updateInterval = 500;
const numHorizontalCells = 20;
const numVerticalCells = 15;
const cellWidth = 12;
const cellHeight = 12;
const colors = {
  empty: "#c5ccd1",
  snake: "#5199db",
  fruit: "#8acf7e",
  missing: "#cbcf7e"
};

export default {
  name: "Screen",
  components: {
    StartButton
  },
  data() {
    return {
      configKonva: {
        width: numHorizontalCells * cellWidth,
        height: numVerticalCells * cellHeight
      },
      grid: [],
      isGameOver: false,
      hasError: false,
      errorMessage: "",
      interval: null
    };
  },
  mounted() {
    this.interval = setInterval(() => this.updateGrid(), updateInterval);
  },
  methods: {
    restartedGame(playerToken) {
      this.$emit("restartgame", playerToken);
      this.isGameOver = false;
      this.interval = setInterval(() => this.updateGrid(), updateInterval);
    },
    updateGrid() {
      // console.log("update grid");

      fetch(`http://${config.serverHost}:${config.serverPort}/screen`)
        .then(response => {
          switch (response.status) {
            case 200:
              return response.text();
            case 400:
              clearInterval(this.interval);
              this.hasError = true;
              this.errorMessage = "Error calling server";
              return Promise.resolve();
            case 404:
              clearInterval(this.interval);
              this.isGameOver = true;
              this.$emit("gameover");
              return Promise.resolve();
            default:
              clearInterval(this.interval);
              this.hasError = true;
              this.errorMessage = "Error fetching screen state";
              return Promise.resolve();
          }
        })
        .then(gameState => {
          if (gameState) {
            this.parseGrid(gameState);
          }
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
      rows.push(row);

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

          // Only set the minimum required fields on every update loop
          const rect = {
            y: i * cellHeight,
            x: j * cellWidth,
            fill: color
          };

          // Check if we already have a rect on this X,Y position
          const existingRect = this.grid.find(
            r => r.x === rect.x && r.y === rect.y
          );

          if (existingRect) {
            // If we have a rect, only update the color
            existingRect.fill = rect.fill;
          } else {
            // If this is a new rect, set one-time/static fields only once
            rect.id = uuidv4();
            rect.width = cellWidth;
            rect.height = cellHeight;
            this.grid.push(rect);
          }
        }
      }
    }
  }
};
</script>

<style lang="scss" scoped>
$boxWidth: 550px;
$boxHeight: 350px;

$boxWidthMob: 100%;
$boxHeightMob: 250px;

.screen {
  display: flex;
  justify-content: center;
}
.game-over {
  display: flex;
  flex-direction: column;
  justify-content: center;
  width: $boxWidth;
  height: $boxHeight;
  background-color: #c5ccd1;

  @media (max-width: 800px) {
    width: $boxWidthMob;
    height: $boxHeightMob;
  }
}
.error-message {
  display: flex;
  flex-direction: column;
  justify-content: center;
  width: $boxWidth;
  height: $boxHeight;
  background-color: #c5ccd1;

  @media (max-width: 800px) {
    width: $boxWidthMob;
    height: $boxHeightMob;
  }
}
</style>
