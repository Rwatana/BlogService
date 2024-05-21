const { serviceAvailable } = require("./serviceAvailable");

const express = require("express");
const bodyParser = require("body-parser");

const app = express();
app.use(bodyParser.json());


serviceAvailable('query', app);