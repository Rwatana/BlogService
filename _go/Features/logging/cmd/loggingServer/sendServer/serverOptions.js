function getServerOptions(dataLength) {
    return {
        hostname: 'loggingserver-srv',
        port: 4007,
        path: '/data',
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Content-Length': dataLength
        }
    };
}

module.exports = getServerOptions;
