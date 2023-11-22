const path = require('path');
const webpack = require('webpack');
const { merge } = require('webpack-merge');
// const { BundleAnalyzerPlugin } = require('webpack-bundle-analyzer');
const TerserPlugin = require('terser-webpack-plugin');
const base = require('./webpack.common');

// const withReport = process.env.npm_config_withReport;
const nodeEnv = process.env.NODE_ENV;

module.exports = merge(base, {
  mode: 'development',
  devtool: 'cheap-module-source-map',
  output: {
    clean: true,
    globalObject: 'this',
    filename: '[name].dev.js',
    chunkFilename: '[name].dev.js',
    path: `${__dirname}/build/${nodeEnv}`,
    publicPath: '',
  },
  // experiments: {
  //   asyncWebAssembly: true,
  //   syncWebAssembly: true,
  // },
  optimization: {
    nodeEnv: 'development',
    minimizer: [
      new TerserPlugin({
        parallel: 4,
        minify: TerserPlugin.esbuildMinify,
        // terserOptions: {
        //   sourceMap: 'external',
        // },
      }),
    ],
  },
  devServer: {
    //server: 'https',
    static: path.join(__dirname, './build/development'),
    compress: true,
    port: 3000,
    open: true,
    hot: true,
    allowedHosts: 'all',
    headers: {
      'Access-Control-Allow-Origin': '*',
      'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, PATCH, OPTIONS',
      'Access-Control-Allow-Headers':
        'X-Requested-With, content-type, Authorization',
    },
  },
  // module: {
  //   rules: [
  //     {
  //       test: /\.wasm$/,
  //       use: ['wasm-loader'],
  //     },
  //     {
  //       test: /\.wasm$/,
  //       type: 'webassembly/experimental', // Specify the type here
  //     },
  //   ],
  // },
  plugins: [
    new webpack.HotModuleReplacementPlugin(),
    new webpack.IgnorePlugin({ resourceRegExp: /[^/]+\/[\S]+.dev$/ }),
    // new BundleAnalyzerPlugin(),
  ],
});
