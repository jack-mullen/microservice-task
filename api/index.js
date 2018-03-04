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

app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());
app.use(expressValidator());

//Routes
const routes = require('./routes');
app.use('/', routes);


app.listen(port, host);