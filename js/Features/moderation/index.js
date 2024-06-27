const express = require('express');
const bodyParser = require('body-parser');
const axios = require('axios');

const app = express();
app.use(bodyParser.json());

const dbcn = require("../log/dbConnect");
const { insertLog } = require("../log/dbSendLog");
const service = 'moderation';

app.post('/events', async (req, res) => {
  current_date = new Date();
  insertLog(dbcn, current_date, service, 'demo');
  const { type, data } = req.body;

  if (type === 'CommentCreated') {
    const status = data.content.includes('orange') ? 'rejected' : 'approved';

    await axios.post('http://localhost:4005/events', {
      type: 'CommentModerated',
      data: {
        id: data.id,
        postId: data.postId,
        status,
        content: data.content
      }
    });
    // if error exists send log to db with error message
    if (status === 'rejected') {
      insertLog(dbcn, current_date, service, 'rejected');
    }
    
  }

  res.send({});
});

app.listen(4003, () => {
  current_date = new Date();
  insertLog(dbcn, current_date, service, 'moderation service is listerning on 4003');
  console.log('Listening on 4003');
});
