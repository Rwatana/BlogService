function getServerOptions(dataLength) {
    return {
        hostname: 'localhost',
        port: 8080,
        path: '/data',
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Content-Length': dataLength
        }
    };
}

module.exports = getServerOptions;
