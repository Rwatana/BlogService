const mysql = require('mysql');

const con = mysql.createConnection({
    host: '127.0.0.1',
    port: '3306',
    user: 'root',
    password: 'lOjit212',
    database: 'test_db',
    table: 'test_log'
});

// connect
con.connect((err) => {
    if (err) throw err;

    console.log('connected to mysql');
});

function sendLogToDB(API, error_message) {
    current_time = new Date();
    data = {
        date: current_time,
        api: API,
        error: error_message
    };
    // insert log
    con.query('INSERT INTO test_log SET ?', data, (err, res) => {
        if (err) throw err;
        console.log('Inserted:', res.insertId);
    });
}


sendLogToDB('example_api', 'example_error_message');

// disconnect
con.end((err) => {
    if (err) throw err;

    console.log('disconnected to mysql');
});
