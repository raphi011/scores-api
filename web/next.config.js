const webpack = require('webpack');
const child_process = require('child_process');
const withTypescript = require('@zeit/next-typescript');

const version = child_process
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
