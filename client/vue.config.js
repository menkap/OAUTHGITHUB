module.exports = {
    devServer: {
      proxy: {
        "/server": {
          target: "http://localhost:4887",
          ws: true,
          changeOrigin: true,
          pathRewrite: { "^/server": "" },
        },
      },
    },
  };
  