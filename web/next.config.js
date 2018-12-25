const webpack = require('webpack');
const childProcess = require('child_process');
const withTypescript = require('@zeit/next-typescript');

const version = childProcess
  .execSync('git describe --always --long')
  .toString();

module.exports = withTypescript({
  webpack(config) {
    config.plugins.push(
      new webpack.DefinePlugin({
        VERSION: JSON.stringify(version),
      }),
    );

    return config;
  },
});
