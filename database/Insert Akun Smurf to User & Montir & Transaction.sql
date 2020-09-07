USE `tublessin_transaction`;
INSERT INTO master_status_transaction(status) VALUE ('On Process'),('Success'),('Canceled');

USE `tublessin_user` ;

insert into user_account(username, password) value
('reizdendy', '$2a$10$QMvlKdQhnpyoaXPLuCTynuifns4VUtgLS8Zo55Iha66RKVY7fEi8G'),
('sandro', '$2a$10$QMvlKdQhnpyoaXPLuCTynuifns4VUtgLS8Zo55Iha66RKVY7fEi8G');

insert into user_profile(user_account_id, firstname, lastname, gender, phone_number, email) value
(1, 'reizdendy', 'akbar', 'L', '08181818', 'reizdendy@gmail.com'),
(2, 'sandro', 'dorooo', 'L', '0818288', 'sandro@gmail.com');

insert into user_location(user_account_id) value(1),(2);

USE `tublessin_montir` ;

INSERT INTO master_status_activity(status) VALUE('Standby'),('On Going'),('Working');

INSERT INTO montir_account(username,password) VALUE
('gerin','$2a$10$QMvlKdQhnpyoaXPLuCTynuifns4VUtgLS8Zo55Iha66RKVY7fEi8G'),
('vio','$2a$10$QMvlKdQhnpyoaXPLuCTynuifns4VUtgLS8Zo55Iha66RKVY7fEi8G'),
('bebek','$2a$10$QMvlKdQhnpyoaXPLuCTynuifns4VUtgLS8Zo55Iha66RKVY7fEi8G'),
('burung','$2a$10$QMvlKdQhnpyoaXPLuCTynuifns4VUtgLS8Zo55Iha66RKVY7fEi8G'),
('Kucing','$2a$10$QMvlKdQhnpyoaXPLuCTynuifns4VUtgLS8Zo55Iha66RKVY7fEi8G');

INSERT INTO montir_profile(montir_account_id, firstname, lastname, born_date, gender, ktp, address, city, phone_number, email) VALUE
(1,'Gerin','Prakoso','1990-01-24','L','123456789','Kec. Rawalumbu','Bekasi','08982279019','gerin@google.com'),
(2,'Viontina','Dea','1995-01-25','P','123444444','Kec. Malang','Malang','08982272727','vio@google.com'),
(3,'Bebek','Air','1992-01-23','P','58868334','Kec. Jakarta','Bogor','089846456','Bebek@google.com'),
(4,'Burung','Terbang','1991-02-22','L','634544','Kec. Bekasi','Jakarta','089887676','Burung@google.com'),
(5,'Kucing','Hitam','1990-02-22','L','632155534544','Kec. Depok','Depok','098323255','Kucing@google.com');

INSERT INTO montir_status(montir_account_id,status_operational) VALUE 
(1,"A"),(2,"A"),(3,"A"),(4,"A"),(5,"A");

INSERT INTO montir_location(montir_account_id,latitude,longitude) VALUE 
(1, -6.158658, 106.856724),
(2, -6.153943, 106.829343),
(3, -6.194667, 106.787583),
(4, -6.194050, 106.828470),
(5, -6.204413, 106.858610);

INSERT INTO montir_rating(montir_account_id, rating, rater_id, review) VALUE
(1, 1, 10, "Bagus Sekali"),(1, 4, 10, "Cakep Sekali"),(1, 3, 10, "Hehe Sekali"),
(2, 5, 11, "Bagus Sekali"),(2, 2, 11, "Cakep Sekali"),(2, 3, 11, "Hehe Sekali"),
(3, 5, 11, "Bagus Sekali"),(2, 2, 11, "Cakep Sekali"),(2, 5, 11, "Hehe Sekali"),
(4, 3, 11, "Bagus Sekali"),(2, 2, 11, "Cakep Sekali"),(2, 3, 11, "Hehe Sekali"),
(5, 3, 11, "Bagus Sekali"),(2, 2, 11, "Cakep Sekali"),(2, 3, 11, "Hehe Sekali");