const { optimize, NoEmitOnErrorsPlugin } = require('webpack');

module.exports = {
  entry : [
    'webpack-dev-server/client?http://localhost:6789', // WebpackDevServer host and port
    'webpack/hot/dev-server', // "only" prevents reload on syntax errors
    './client/index' // Your appʼs entry point
  ],

  output : {
    filename   : 'bundle.js',
    publicPath : 'http://localhost:6789/assets/'
  },

  resolve : {
    extensions : ['.jsx', '.js', '.json', '.less']
  },

  module : {
    rules : [{
      test    : /\.(js|jsx)$/,
      exclude : [/node_modules/],
      loader  : 'babel-loader',
      options : {
        presets : [
          'env',
          'react'
        ],
        plugins : [
          'syntax-decorators',
          'transform-class-properties',
          'transform-decorators-legacy',
          'transform-export-extensions',
          'transform-object-rest-spread'
        ]
      }
    }, {
      test    : /\.woff(2)?(\?v=[0-9]\.[0-9]\.[0-9])?$/,
      loader  : 'url-loader',
      options : {
        limit    : 10000,
        minetype : 'application/font-woff'
      }
    }, {
      test   : /\.(ttf|eot|svg)(\?v=[0-9]\.[0-9]\.[0-9])?$/,
      loader : 'file-loader'
    }, {
      test    : /\.(jpe?g|png|gif|svg)$/i,
      loader  : 'url-loader',
      options : { limit: 10000 }
    }, {
      test : /\.css$/,
      use  : [
        'style-loader',
        {
          loader  : 'css-loader',
          options : {
            importLoaders : 1,
            minimize      : true
          }
        }
      ]
    }]
  },

  plugins : [
    new optimize.ModuleConcatenationPlugin(),
    new NoEmitOnErrorsPlugin()
  ]
};
