const webpack = require('webpack');
const path = require('path');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
const StatsPlugin = require('stats-webpack-plugin');
const CleanWebpackPlugin = require('clean-webpack-plugin');

const dist = 'assets'

module.exports = {
  entry: {
    'main': path.join(__dirname, 'src', 'main.js'),
    'vendor': ['jquery', 'tether', 'bootstrap', 'marked'],
  },
  output: {
    filename: '[name].[chunkhash].js',
    path: path.join(__dirname, dist),
  },
  module: {
    rules: [
      {
        test: /\.css$/,
        exclude: /node_modules/,
        loader: ExtractTextPlugin.extract({
          loader: 'css-loader?sourceMap'
        })
      },
      {
        test: /\.scss$/,
        loaders: ["style-loader", "css-loader", "sass-loader"]
      },
      { test: /\.(png|woff|woff2|eot|ttf|svg)$/, loader: "file-loader?name=[name].[hash].[ext]" }
    ]
  },
  devtool: 'source-map',
  plugins: [
    new CleanWebpackPlugin([dist]),
    // fix " Bootstrap's JavaScript requires jQuery"
    new webpack.ProvidePlugin({
      $: "jquery",
      jQuery: "jquery",
      "window.jQuery": "jquery",
      "window.Tether": 'tether'
    }),
    new StatsPlugin('stats.json', {
      chunkModules: true,
      exclude: [/node_modules[\\\/]react/]
    }),
    new ExtractTextPlugin({ filename: '[name].[chunkhash].css', disable: false, allChunks: true }),
    new webpack.LoaderOptionsPlugin({
      minimize: true,
      debug: false
    }),
    new webpack.DefinePlugin({
      'process.env': {
          'NODE_ENV': JSON.stringify('prod')
      }
    }),
    new webpack.optimize.CommonsChunkPlugin({name: ['vendor', 'manifest']}),
    new webpack.optimize.UglifyJsPlugin({
      beautify: false,
      mangle: {
          screw_ie8: true,
          keep_fnames: true
      },
      compress: {
          screw_ie8: true
      },
      comments: false,
      sourceMap: true,
    })
  ]
}
