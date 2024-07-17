const MiniCssExtractPlugin = require("mini-css-extract-plugin")
const TsconfigPathsPlugin = require('tsconfig-paths-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');

const path = require("path");

module.exports = {
    entry: "./static/scripts/index.ts",
    module: {
        rules: [
            {
                test: /\.ts$/,
                use: 'ts-loader',
                exclude: /node_modules/,
            },
            {
                test: /\.css$/i,
                use: [
                    MiniCssExtractPlugin.loader,
                    'css-loader'
                ]
            },
            {
                test: /\.png$/i,
                type: 'asset',
              },

        ],
    },
    resolve: {
        alias: {
            Images: path.resolve(__dirname, './static/images')
        },
        extensions: [ '.ts', '.js'],
        plugins: [new TsconfigPathsPlugin({
            configFile: 'tsconfig.json'
        })],
    },
    output: {
        path: path.resolve(__dirname, './static/dist'),
        filename: `index_bundle.[contenthash].js`
    },
    plugins: [
        new MiniCssExtractPlugin(),
        new HtmlWebpackPlugin({
            template: './pages/index.html',
            inject: "head"
        })
      ]
};