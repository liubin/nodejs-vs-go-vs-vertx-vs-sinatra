var http = require('http');
http.createServer(function (req, res) {
  res.writeHead(200, {'Content-Type': 'text/plain'});
  var hash = require('crypto').createHash('md5');
  var time = new Date().getTime(); 
  var resp = hash.update('' + time).digest('hex');
  res.end(resp + '\n');
}).listen(5555, '127.0.0.1');
console.log('Server running at http://127.0.0.1:5555/');
