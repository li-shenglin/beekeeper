const {defineConfig} = require('@vue/cli-service')
const path = require('path')
const resolve = dir => {
    return path.join(__dirname, dir)
}
const BASE_URL = process.env.NODE_ENV === 'production'
    ? '/'
    : '/'
module.exports = defineConfig({
    transpileDependencies: true,
    publicPath: BASE_URL,
    lintOnSave: true,
    chainWebpack: config => {
        config.resolve.alias
            .set('@', resolve('src'))
            .set('_c', resolve('src/components'))
            .set("vue-i18n", "vue-i18n/dist/vue-i18n.cjs.js")
    },
    productionSourceMap: false,
    devServer: {
        port: 8090,
        proxy: {
            '/apis': {
                target: 'http://127.0.0.1:8080',
                changeOrigin: true,
                pathRewrite: {
                    '^/apis': '/'
                }
            }
        }
    }
})