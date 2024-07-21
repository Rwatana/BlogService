const http = require('http');
const getServerOptions = require('./serverOptions');

function sendJsonData(data, callback) {
    const jsonData = JSON.stringify(data);
    const options = getServerOptions(Buffer.byteLength(jsonData));

    const req = http.request(options, (res) => {
        let responseBody = '';

        res.on('data', (chunk) => {
            responseBody += chunk;
        });

        res.on('end', () => {
            callback(null, responseBody);
        });
    });

    req.on('error', (error) => {
        callback(error);
    });

    req.write(jsonData);
    req.end();
}

module.exports = sendJsonData;
