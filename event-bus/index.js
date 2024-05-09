const express = require("express");
const bodyParser = require("body-parser");
const axios = require("axios");

const app = express();
app.use(bodyParser.json());

const events = [];

const dbcn = require("../log/dbConnect");
const { insertLog } = require("../log/dbSendLog");
const e = require("express");
const current_service = 'event-bus';


app.post("/events", (req, res) => {
  current_date = new Date();
  const event = req.body;

  events.push(event);

  axios.post("http://localhost:4000/events", event).catch((err) => {
    insertLog(dbcn, current_date, current_service, source_service, API, err.message);
    console.log(err.message);
  });
  axios.post("http://localhost:4001/events", event).catch((err) => {
    insertLog(dbcn, current_date, current_service, source_service, API, err.message);

    console.log(err.message);
  });
  axios.post("http://localhost:4002/events", event).catch((err) => {
    insertLog(dbcn, current_date, current_service, source_service, API, err.message);

    console.log(err.message);
  });
  axios.post("http://localhost:4003/events", event).catch((err) => {
    insertLog(dbcn, current_date, current_service, source_service, API, err.message);

    console.log(err.message);
  });
  res.send({ status: "OK" });
});

app.get("/events", (req, res) => {
  current_date = new Date();
  insertLog(dbcn, current_date, current_service, source_service, API, 'demo');
  res.send(events);
});

app.listen(4005, () => {
  current_date = new Date();
  insertLog(dbcn, current_date, current_service, source_service, API, 'event-bus server is running on port 4005');
  console.log("Listening on 4005");

});
