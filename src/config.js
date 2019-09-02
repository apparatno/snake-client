const config = {
  mock: {
    serverHost: 'http://127.0.0.1',
    serverPort: '8888',
  },
  dev: {
    serverHost: 'http://127.0.0.1',
    serverPort: '8081',
  },
  prod: {
    serverHost: 'https://api.snake.apparat.no',
    serverPort: '',
  },
};

// Bytt ut gjeldende miljø her:
export default config.prod;
