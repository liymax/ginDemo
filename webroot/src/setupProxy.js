const proxy = require('http-proxy-middleware');

module.exports = function (app) {
  app.use(
    proxy('/*', {
      target: 'http://localhost:8080/',
      changeOrigin: true
    })
  );
  app.use(
    proxy('/sys', {
      target: 'http://localhost:8080/',
      changeOrigin: true
    })
  );
};