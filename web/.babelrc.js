module.exports = {
  presets: ['next/babel', '@zeit/next-typescript/babel'],
  env: {
    test: {
      presets: [['next/babel', { 'preset-env': { modules: 'commonjs' } }]],
    },
  },
};
