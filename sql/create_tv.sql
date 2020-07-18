use porter_db;

drop table if exists tbl_tv_series_files;
drop table if exists tbl_tv_series;

create table tbl_tv_series (
    id varchar(36) primary key ,
    name varchar(200) not null ,
    year int not null ,
    language varchar(100) not null,
    folder_path varchar(255) not null ,
    episodic bool not null default false,
    created_ts timestamp default current_timestamp
);

create table tbl_tv_series_files (
    id varchar(36) primary key ,
    series_id varchar(36) not null ,
    file_path varchar(255) not null ,
    file_name varchar(255) not null ,
    episode_number int ,
    season_number int ,
    foreign key (series_id) references tbl_tv_series(id)
);
