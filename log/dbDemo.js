const mysql = require('mysql');

const con = mysql.createConnection({
    host: '127.0.0.1',
    port: '3306',
    user: 'root',
    password: 'lOjit212',
    database: 'test_db',
    table: 'test_log3'
});

// connect
con.connect((err) => {
    if (err) throw err;

    console.log('connected to mysql');
});




function sendLogToDB(current_service,source_service,API, error_message) {
    current_time = new Date();
    const data = {
        log_level: 'EMERGE',
        date: current_time,
        current_service: current_service,
        source_service: source_service,
        type_of_request: API,
        content: error_message
      };
    // insert log
    con.query('INSERT INTO test_log3 SET ?', data, (err, res) => {
        if (err) throw err;
        console.log('Inserted:', res.insertId);
    });
}


sendLogToDB('event-bus', 'comments' ,'POST' ,'demo');

// disconnect
con.end((err) => {
    if (err) throw err;

    console.log('disconnected to mysql');
});
