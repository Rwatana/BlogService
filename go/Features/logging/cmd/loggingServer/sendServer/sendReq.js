const amqp = require('amqplib');

const log_data = {
    log_level: 'EMERGE',
    date: new Date().toISOString(),
    current_service: 'event-bus',
    source_service: 'comments',
    type_of_request: 'POST',
    content: 'demo'
};

async function sendLog(logMessage) {
  try {
    const connection = await amqp.connect('amqp://rabbitmq');
    const channel = await connection.createChannel();
    const queue = 'logs';

    await channel.assertQueue(queue, { durable: true });
    const logMessageBuffer = Buffer.from(JSON.stringify(logMessage));
    channel.sendToQueue(queue, logMessageBuffer, { persistent: true });

    console.log('Log sent:', logMessage);
    await channel.close();
    await connection.close();
  } catch (error) {
    console.error('Error sending log:', error);
  }
}

sendLog(log_data);
