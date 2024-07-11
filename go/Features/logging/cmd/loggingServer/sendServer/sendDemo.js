const sendJsonData = require('./sendRequest');

const current_time = new Date().toISOString();

const log_data = {
    log_level: 'EMERGE',
    date: current_time,
    current_service: 'event-bus',
    source_service: 'comments',
    type_of_request: 'POST',
    content: 'demo'
};

sendJsonData(log_data, (error, response) => {
    if (error) {
        console.error(`Error: ${error.message}`);
    } else {
        console.log(`Response: ${response}`);
    }
});
