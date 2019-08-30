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
    serverHost: '10.0.0.155',
    serverPort: '8080',
  },
};

// Bytt ut gjeldende miljø her:
export default config.dev;
