const mysql = require("mysql");

const dbcn = mysql.createConnection({
    host: '127.0.0.1',
    port: '3306',
    user: 'root',
    password: 'lOjit212',
    database: 'test_db',
    table: 'test_log3'
});

module.exports= dbcn;

