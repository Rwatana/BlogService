const express = require("express");
const bodyParser = require("body-parser");
const axios = require("axios");
const sendJsonData = require('./sendServer/sendRequest');
const app = express();
app.use(bodyParser.json());

const events = [];

app.post("/events", (req, res) => {
  const event = req.body;

  events.push(event);

  axios.post("http://posts-clusterip-srv:4000/events", event).catch((err) => {
    console.log(err.message);
  });
  sendJsonData({ log_level: 'NORMAL', date: new Date().toISOString(), current_service: 'event-bus', source_service: 'posts', type_of_request: 'POST', content: 'demo' });

  axios.post("http://comments-srv:4001/events", event).catch((err) => {
    console.log(err.message);
  });
  sendJsonData({ log_level: 'NORMAL', date: new Date().toISOString(), current_service: 'event-bus', source_service: 'comments', type_of_request: 'POST', content: 'demo' });
  axios.post("http://query-srv:4002/events", event).catch((err) => {
    console.log(err.message);
  });
  sendJsonData({ log_level: 'NORMAL', date: new Date().toISOString(), current_service: 'event-bus', source_service: 'query', type_of_request: 'POST', content: 'demo' });
  axios.post("http://moderation-srv:4003/events", event).catch((err) => {
    console.log(err.message);
  });
  sendJsonData({ log_level: 'NORMAL', date: new Date().toISOString(), current_service: 'event-bus', source_service: 'moderation', type_of_request: 'POST', content: 'demo' });
  res.send({ status: "OK" });
});

app.get("/events", (req, res) => {
  res.send(events);
});

app.listen(4005, () => {
  console.log("Listening on 4005");
});

