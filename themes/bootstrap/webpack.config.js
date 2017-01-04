var webpack = require('webpack');
var path = require('path');
var ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
  entry: {
    'main': path.join(__dirname, 'src', 'main.js'),
    'vendor': ['jquery', 'bootstrap', 'marked'],
  },
  output: {
    filename: '[name].[chunkhash].js',
    path: './assets'
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
      // { test: /\.css$/, loader: "style-loader!css-loader" },
      { test: /\.png$/, loader: "url-loader?limit=100000" }
    ]
  },
  devtool: 'source-map',
  plugins: [
    new ExtractTextPlugin({ filename: 'bundle.css', disable: false, allChunks: true }),
    new webpack.optimize.CommonsChunkPlugin({name: ['vendor', 'manifest']})
  ]
}
