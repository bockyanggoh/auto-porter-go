use porter_db;

drop table if exists tbl_tv_files;
drop table if exists tbl_tv;
drop table if exists tbl_movie_files;
drop table if exists tbl_movie;
drop table if exists tbl_batch_log;

create table tbl_tv (
    id varchar(36) primary key ,
    name varchar(200) unique not null ,
    year int not null ,
    language varchar(100) not null,
    folder_path varchar(255) unique not null ,
    episodic bool not null default false,
    created_ts timestamp default current_timestamp
);

create table tbl_tv_files (
    id varchar(36) primary key ,
    parent_id varchar(36) not null ,
    parent_folder varchar(255) not null ,
    parent_folder_path varchar(255) not null ,
    file_path varchar(255) unique not null ,
    file_name varchar(255) unique not null ,
    episode_number int ,
    season_number int ,
    created_ts timestamp default current_timestamp ,
    updated_ts timestamp default current_timestamp ,
    foreign key (parent_id) references tbl_tv(id)
);

create table tbl_movie (
    id varchar(36) primary key ,
    tmdb_id varchar(36) unique,
    name varchar(200) not null ,
    year int not null ,
    language varchar(100) not null,
    folder_path varchar(255) not null ,
    created_ts timestamp default current_timestamp ,
    updated_ts timestamp default current_timestamp
);

create table tbl_movie_files (
    id varchar(36) primary key ,
    parent_id varchar(36) not null ,
    parent_folder_name varchar(255) not null ,
    parent_folder_path varchar(255) not null ,
    file_path varchar(255) unique not null ,
    file_name varchar(255) unique not null ,
    created_ts timestamp default current_timestamp ,
    updated_ts timestamp default current_timestamp ,
    foreign key (parent_id) references tbl_movie(id)
);

create table tbl_batch_log(
    id int auto_increment primary key ,
    start_execution_time timestamp not null ,
    end_execution_time timestamp not null ,
    files_detected int not null ,
    folders_detected int not null ,
    job_type varchar(30) not null
);