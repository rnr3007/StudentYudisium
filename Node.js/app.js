const http = require('http');
const path = require('path');
const bodyParser = require('body-parser');
const express = require('express');
const route = require('./src/router/routes');

const app = express();
app.use(bodyParser.urlencoded({extended: false}));
app.use(express.static(path.join(__dirname, 'assets')));
app.use(express.static(path.join(__dirname, 'src')));
app.use("", route);

const port = 8080;
const server = http.createServer(app);
server.listen(port);

console.log(`> Server started on http://localhost:${port}/`);