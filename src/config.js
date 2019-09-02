const config = {
  mock: {
    serverHost: '127.0.0.1',
    serverPort: '8888',
  },
  dev: {
    serverHost: '127.0.0.1',
    serverPort: '8081',
  },
  prod: {
    serverHost: 'api.snake.apparat.no',
    serverPort: '80',
  },
};

// Bytt ut gjeldende miljø her:
export default config.prod;
