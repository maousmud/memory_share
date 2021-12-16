drop table m_movie;

create table m_movie
(id INT AUTO_INCREMENT NOT NULL,
 title VARCHAR(100) NOT NULL,
 movie_file BLOB NOT NULL,
 register_username VARCHAR(100) NOT NULL,
 remarks VARCHAR(3000),
 memory_shubetsu VARCHAR(100) NOT NULL,
 created_date DATE NOT NULL,
 update_date DATE NOT NULL,
 del_flg BOOLEAN NOT NULL,
 PRIMARY KEY(id,title)
);