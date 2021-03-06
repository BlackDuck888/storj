// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

const path = require('path');
const CompressionWebpackPlugin = require('compression-webpack-plugin');
const productionGzipExtensions = ['js', 'css', 'ttf'];

module.exports = {
    publicPath: "/static/dist",
    productionSourceMap: false,
    parallel: true,
    configureWebpack: {
        plugins: [
            new CompressionWebpackPlugin({
                algorithm: 'gzip',
                test: new RegExp('\\.(' + productionGzipExtensions.join('|') + ')$'),
                threshold: 10240,
                minRatio: 0.8
            })
        ],
    },
    chainWebpack: config => {
        config.output.chunkFilename(`js/vendors.js`);
        config.output.filename(`js/app.js`);

        config.resolve.alias
            .set('@', path.resolve('src'));

        config
            .plugin('html')
            .tap(args => {
                args[0].template = './index.html';
                return args
            });
    }
};
