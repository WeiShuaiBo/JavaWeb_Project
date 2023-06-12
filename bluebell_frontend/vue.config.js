module.exports = {
    assetsDir: "static",
    devServer: {
        publicPath: "/",
        host:"127.0.0.1",
        port:"3333",
        proxy: {
            '/api/v1': {
              target: 'http://127.0.0.1:8081',
              changeOrigin: true,
            }
        }
    }
  }


