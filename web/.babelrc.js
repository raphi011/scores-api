module.exports = {
  presets: ['next/babel'],
  plugins: ['transform-flow-strip-types'],
  env: {
    test: {
      presets: [['next/babel', { 'preset-env': { modules: 'commonjs' } }]],
    },
  },
};
