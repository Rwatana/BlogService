const http = require('http');

const serviceAvailable = (current_service,  res) => {
    if (current_service === 'query') {
        http.get('http://localhost:4002/', (resp) => {
        }).on("error", (err) => {
            return err;
        });
        return res.status(200).send('Service is available');
    } else if (current_service === 'event-bus') {
        http.get('http://localhost:4005/', (resp) => {
            console.log('Service is available');
        }).on("error", (err) => {
            console.log('Service is not available');
        });
        return res.status(200).send('Service is available');
    } else if (current_service === 'comments') {
        http.get('http://localhost:4001/', (resp) => {
            console.log('Service is available');
        }).on("error", (err) => {
            console.log('Service is not available');
        });
    } else if (current_service === 'moderation') {
        http.get('http://localhost:4003/', (resp) => {
            console.log('Service is available');
        }).on("error", (err) => {
            console.log('Service is not available');
        });
        return res.status(200).send('Service is available');
    } else if (current_service === 'posts') {
        http.get('http://localhost:4000/', (resp) => {
            console.log('Service is available');
        }).on("error", (err) => {
            console.log('Service is not available');
        });
        return res.status(200).send('Service is available');
    } else {
        return res.status(500).send('Service is not available');
    }
};

module.exports = {
    serviceAvailable
};
