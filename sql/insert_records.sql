use porter_db;

INSERT INTO tbl_batch_log(id, start_execution_time, end_execution_time, files_detected, folders_detected, job_type) VALUES (?, ?, ?, ?, ?, ?);

SELECT * FROM tbl_batch_log WHERE job_type = ? order by end_execution_time desc limit 1;

select count(*) from tbl_batch_log where id = ?