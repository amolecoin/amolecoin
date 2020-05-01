const PROXY_CONFIG = {
  "/api/*": {
    "target": "http://127.0.0.1:49981",
    "secure": false,
    "logLevel": "debug",
    "bypass": function (req) {
      req.headers["host"] = '127.0.0.1:49981';
      req.headers["referer"] = 'http://127.0.0.1:49981';
      req.headers["origin"] = 'http://127.0.0.1:49981';
    }
  }
};

module.exports = PROXY_CONFIG;
