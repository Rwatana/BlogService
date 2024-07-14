-- demo data
INSERT INTO logging (log_level, date, current_service, source_service, type_of_request, content)
VALUES 
('INFO', NOW(), 'ServiceA', 'ServiceB', 'GET', 'This is an info log message'),
('ERROR', NOW(), 'ServiceB', 'ServiceC', 'POST', 'This is an error log message'),
('DEBUG', NOW(), 'ServiceC', 'ServiceA', 'PUT', 'This is a debug log message'),
('WARN', NOW(), 'ServiceA', 'ServiceC', 'DELETE', 'This is a warning log message'),
('INFO', NOW(), 'ServiceB', 'ServiceA', 'PATCH', 'This is another info log message');
