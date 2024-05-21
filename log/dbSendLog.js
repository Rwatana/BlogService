const insertLog = (dbcn, log_level,current_time, current_service, source_service, API, error_message,  res) => {
    const db_data = {
      log_level: log_level,
        date: current_time,
        current_service: current_service,
        source_service: source_service,
        type_of_request: API,
        content: error_message
      };
    dbcn.query('INSERT INTO test_log3 SET ?', db_data, (err, result) => {
      if (err) {
        console.error('Error inserting data into DB:', err);
        return res.status(500).send('Error inserting data into DB');
      }
      console.log('Inserted:', result.insertId);
    });
  };

module.exports = {
    insertLog
};