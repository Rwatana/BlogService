const insertLog = (dbcn, current_time, API,  res) => {
    const db_data = {
      date: current_time,
      api: API,
      error: 'demo'
    };
    dbcn.query('INSERT INTO test_log SET ?', db_data, (err, result) => {
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
