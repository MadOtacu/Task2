const MiniCssExtractPlugin = require("mini-css-extract-plugin")
const TsconfigPathsPlugin = require('tsconfig-paths-webpack-plugin');

const path = require("path")

module.exports = {
    entry: "./static/scripts/index.ts",
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                use: 'ts-loader',
            },
            {
                test: /\.css$/i,
                use: [
                    MiniCssExtractPlugin.loader,
                    'css-loader'
                ]
            },
        ],
    },
    resolve: {
        extensions: ['.css', '.ts', '.js'],
        plugins: [new TsconfigPathsPlugin({
            configFile: 'tsconfig.json'
        })]
    },
    output: {
        path: path.resolve(__dirname, './static/dist'),
        filename: 'index_bundle.js'
    },
    plugins: [
        new MiniCssExtractPlugin(),
      ]
};