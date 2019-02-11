const withTypescript = require('@zeit/next-typescript');

module.exports = withTypescript({
  env: {
    VERSION: JSON.stringify(process.env.VERSION),
  },
  onDemandEntries: {
    websocketPort: 3001,
    websocketProxyPort: 7000,
  },
});
