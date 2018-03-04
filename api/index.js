'use strict';

const config = require('./config');

const express = require('express');
const bodyParser = require("body-parser");

// Constants
const port = config.port;
const host = '0.0.0.0';

// App
const app = express();

app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

//Routes
const routes = require('./routes');
app.use('/', routes);


app.listen(port, host);