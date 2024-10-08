const express = require('express');
const bodyParser = require('body-parser');
const axios = require('axios');
const sendJsonData = require('./sendServer/sendRequest');

const app = express();
app.use(bodyParser.json());

app.post('/events', async (req, res) => {
  const { type, data } = req.body;

  if (type === 'CommentCreated') {
    const status = data.content.includes('orange') ? 'rejected' : 'approved';

    await axios.post('http://event-bus-srv:4005/events', {
      type: 'CommentModerated',
      data: {
        id: data.id,
        postId: data.postId,
        status,
        content: data.content
      }
    });
  }
  sendJsonData({ log_level: 'NORMAL', date: new Date().toISOString(), current_service: 'moderation', source_service: 'event-bus', type_of_request: 'POST', content: 'Send results to evnet-bus' });

  res.send({});
});

app.listen(4003, () => {
  console.log('Listening on 4003');
});
