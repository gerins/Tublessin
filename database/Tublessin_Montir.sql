-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema tublessin_montir
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema tublessin_montir
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `tublessin_montir` DEFAULT CHARACTER SET utf8 ;
-- -----------------------------------------------------
-- Schema tublessin_user
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema tublessin_chat
-- -----------------------------------------------------
USE `tublessin_montir` ;

-- -----------------------------------------------------
-- Table `tublessin_montir`.`montir_account`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_montir`.`montir_account` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(15) NOT NULL DEFAULT '',
  `password` VARCHAR(300) NOT NULL DEFAULT '',
  `status_account` VARCHAR(1) NOT NULL DEFAULT 'A',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `username_UNIQUE` (`username` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tublessin_montir`.`montir_profile`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_montir`.`montir_profile` (
  `montir_account_id` INT NOT NULL,
  `firstname` VARCHAR(45) NOT NULL DEFAULT '',
  `lastname` VARCHAR(45) NOT NULL DEFAULT '',
  `born_date` DATE NOT NULL DEFAULT '1900-01-01',
  `gender` VARCHAR(1) NOT NULL DEFAULT 'L',
  `ktp` VARCHAR(45) NOT NULL DEFAULT '',
  `address` VARCHAR(120) NOT NULL DEFAULT '',
  `city` VARCHAR(45) NOT NULL DEFAULT '',
  `email` VARCHAR(45) NOT NULL DEFAULT '',
  `phone_number` VARCHAR(15) NOT NULL DEFAULT '',
  `imageURL` VARCHAR(45) NOT NULL DEFAULT 'default_profile.jpg',
  `verified_account` VARCHAR(1) NOT NULL DEFAULT 'N',
  `date_updated` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `date_created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`montir_account_id`),
  UNIQUE INDEX `phone_number_UNIQUE` (`phone_number` ASC) VISIBLE,
  UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE,
  CONSTRAINT `fk_montir_profile_montir_account`
    FOREIGN KEY (`montir_account_id`)
    REFERENCES `tublessin_montir`.`montir_account` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tublessin_montir`.`montir_rating`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_montir`.`montir_rating` (
  `montir_account_id` INT NOT NULL,
  `rating` INT NOT NULL,
  `rater_id` VARCHAR(45) NOT NULL,
  `review` VARCHAR(200) NOT NULL DEFAULT '',
  `date_created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT `fk_montir_rating_montir_account1`
    FOREIGN KEY (`montir_account_id`)
    REFERENCES `tublessin_montir`.`montir_account` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tublessin_montir`.`montir_location`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_montir`.`montir_location` (
  `montir_account_id` INT NOT NULL,
  `latitude` DOUBLE NOT NULL DEFAULT -6.175439,
  `longitude` DOUBLE NOT NULL DEFAULT 106.827227,
  `date_updated` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`montir_account_id`),
  CONSTRAINT `fk_montir_location_montir_account1`
    FOREIGN KEY (`montir_account_id`)
    REFERENCES `tublessin_montir`.`montir_account` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tublessin_montir`.`master_status_activity`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_montir`.`master_status_activity` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `status` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tublessin_montir`.`montir_status`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_montir`.`montir_status` (
  `montir_account_id` INT NOT NULL,
  `status_operational` VARCHAR(1) NOT NULL DEFAULT 'N',
  `status_activity_id` INT NOT NULL DEFAULT 1,
  `date_updated` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`montir_account_id`, `status_activity_id`),
  INDEX `fk_montir_status_status_activity1_idx` (`status_activity_id` ASC) VISIBLE,
  CONSTRAINT `fk_montir_status_montir_account1`
    FOREIGN KEY (`montir_account_id`)
    REFERENCES `tublessin_montir`.`montir_account` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_montir_status_status_activity1`
    FOREIGN KEY (`status_activity_id`)
    REFERENCES `tublessin_montir`.`master_status_activity` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

USE `tublessin_montir` ;

-- -----------------------------------------------------
-- Placeholder table for view `tublessin_montir`.`montir_rating_location_view`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_montir`.`montir_rating_location_view` (`id` INT, `status_account` INT, `firstname` INT, `lastname` INT, `imageURL` INT, `status_operational` INT, `status` INT, `latitude` INT, `longitude` INT, `date_updated` INT, `total_rating` INT, `average_rating` INT);

-- -----------------------------------------------------
-- View `tublessin_montir`.`montir_rating_location_view`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `tublessin_montir`.`montir_rating_location_view`;
USE `tublessin_montir`;
CREATE  OR REPLACE VIEW `montir_rating_location_view` AS
SELECT ma.id, ma.status_account,mp.firstname, mp. lastname, mp.imageURL, ms.status_operational, msa.status, 
ml.latitude, ml.longitude, ml.date_updated, mr.total_rating , mr.average_rating
FROM montir_account ma 
JOIN montir_status ms ON ma.id = ms.montir_account_id
JOIN master_status_activity msa ON ms.status_activity_id = msa.id
JOIN montir_location ml ON ma.id = ml.montir_account_id
JOIN montir_profile mp ON ma.id = mp.montir_account_id
JOIN (SELECT montir_account_id as id, count(montir_account_id) as total_rating, AVG(rating) average_rating FROM montir_rating 
GROUP BY montir_account_id) mr
ON mr.id = ma.id;

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

INSERT INTO master_status_activity(status) VALUE('Standby'),('On Going'),('Working');

INSERT INTO montir_account(username,password) VALUE('gerin','admin'),('vio','admin');

INSERT INTO montir_profile(montir_account_id, firstname, lastname, born_date, gender, ktp, address, city, phone_number, email) VALUE
(1,'Gerin','Prakoso','1990-01-24','L','123456789','Kec. Rawalumbu','Bekasi','08982279019','gerin@google.com'),
(2,'Viontina','Dea','1995-01-25','P','123444444','Kec. Malang','Malang','08982272727','vio@google.com');

INSERT INTO montir_status(montir_account_id) VALUE (1),(2);
INSERT INTO montir_location(montir_account_id) VALUE (1),(2);

INSERT INTO montir_rating(montir_account_id, rating, rater_id, review) VALUE
(1, 5, 10, "Bagus Sekali"),(1, 4, 10, "Cakep Sekali"),(1, 3, 10, "Hehe Sekali"),
(2, 5, 11, "Bagus Sekali"),(2, 4, 11, "Cakep Sekali"),(2, 3, 11, "Hehe Sekali");
