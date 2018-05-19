const env = require('./env-config');

module.exports = {
  presets: ['next/babel'],
  plugins: ['transform-flow-strip-types', ['transform-define', env]],
};
