module.exports = {
  transpileDependencies: ['vuetify'],
  assetsDir: 'static',
  chainWebpack: config => {
    config.plugin('html').tap(args => {
      args[0].title = 'Welcome To MyBlog'
      return args
    })
  },
  productionSourceMap: false
}