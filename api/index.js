'use strict';

const config = require('./config');

const express = require('express');
const bodyParser = require("body-parser");

const expressValidator = require('express-validator');

// Constants
const port = config.port;
const host = '0.0.0.0';

// App
const app = express();

app.use(function(req, res, next) {
  res.header("Access-Control-Allow-Origin", "*");
  res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
  res.header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE");
  next();
});

app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());
app.use(expressValidator());

//Routes
const routes = require('./routes');
app.use('/', routes);


app.listen(port, host);