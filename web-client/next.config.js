const webpack = require('webpack');
const withTypescript = require('@zeit/next-typescript');

module.exports = withTypescript({
  webpack(config) {
    config.plugins.push(
      new webpack.DefinePlugin({
        VERSION: JSON.stringify(process.env.VERSION),
      }),
    );

    return config;
  },
});
