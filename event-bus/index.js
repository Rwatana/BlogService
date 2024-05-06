const express = require("express");
const bodyParser = require("body-parser");
const axios = require("axios");

const app = express();
app.use(bodyParser.json());

const events = [];

const dbcn = require("../log/dbConnect");
const { insertLog } = require("../log/dbSendLog");
const service = 'event-bus';

app.post("/events", (req, res) => {
  current_date = new Date();
  const event = req.body;

  events.push(event);

  axios.post("http://localhost:4000/events", event).catch((err) => {
    insertLog(dbcn, current_date, service, 'demo');
    console.log(err.message);
  });
  axios.post("http://localhost:4001/events", event).catch((err) => {
    insertLog(dbcn, current_date, service, 'demo');

    console.log(err.message);
  });
  axios.post("http://localhost:4002/events", event).catch((err) => {
    insertLog(dbcn, current_date, service, 'demo');

    console.log(err.message);
  });
  axios.post("http://localhost:4003/events", event).catch((err) => {
    insertLog(dbcn, current_date, service, 'demo');

    console.log(err.message);
  });
  res.send({ status: "OK" });
});

app.get("/events", (req, res) => {
  current_date = new Date();
  insertLog(dbcn, current_date, service, 'demo');
  res.send(events);
});

app.listen(4005, () => {
  current_date = new Date();
  insertLog(dbcn, current_date, service, 'event-bus service is listerning on 4005');
  console.log("Listening on 4005");

});
