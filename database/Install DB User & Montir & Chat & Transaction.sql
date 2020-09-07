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
-- Schema tublessin_user
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `tublessin_user` ;
-- -----------------------------------------------------
-- Schema tublessin_chat
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema tublessin_chat
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `tublessin_chat` ;
-- -----------------------------------------------------
-- Schema tublessin_transaction
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema tublessin_transaction
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `tublessin_transaction` ;
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

USE `tublessin_user` ;

-- -----------------------------------------------------
-- Table `tublessin_user`.`user_account`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_user`.`user_account` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(15) NOT NULL DEFAULT '',
  `password` VARCHAR(300) NOT NULL DEFAULT '',
  `status_account` VARCHAR(1) NOT NULL DEFAULT 'A',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `username_UNIQUE` (`username` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tublessin_user`.`user_profile`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_user`.`user_profile` (
  `user_account_id` INT NOT NULL,
  `firstname` VARCHAR(45) NOT NULL DEFAULT '',
  `lastname` VARCHAR(45) NOT NULL DEFAULT '',
  `gender` VARCHAR(1) NOT NULL DEFAULT 'L',
  `phone_number` VARCHAR(15) NOT NULL DEFAULT '',
  `email` VARCHAR(45) NOT NULL DEFAULT '',
  `imageURL` VARCHAR(45) NOT NULL DEFAULT 'default_profile.jpg',
  `date_updated` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `date_created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE INDEX `phone_number_UNIQUE` (`phone_number` ASC) VISIBLE,
  PRIMARY KEY (`user_account_id`),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE,
  CONSTRAINT `fk_user_profile_user_account1`
    FOREIGN KEY (`user_account_id`)
    REFERENCES `tublessin_user`.`user_account` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tublessin_user`.`user_location`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_user`.`user_location` (
  `user_account_id` INT NOT NULL,
  `latitude` DOUBLE NOT NULL DEFAULT -6.175439,
  `longitude` DOUBLE NOT NULL DEFAULT 106.827227,
  `date_updated` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_account_id`),
  CONSTRAINT `fk_user_location_user_account`
    FOREIGN KEY (`user_account_id`)
    REFERENCES `tublessin_user`.`user_account` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

USE `tublessin_chat` ;

-- -----------------------------------------------------
-- Table `tublessin_chat`.`conversations`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_chat`.`conversations` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `sender_id` INT NOT NULL,
  `receiver_id` INT NOT NULL,
  `message` VARCHAR(300) NOT NULL,
  `status` VARCHAR(1) NOT NULL DEFAULT 'A',
  `date_created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

USE `tublessin_transaction` ;

-- -----------------------------------------------------
-- Table `tublessin_transaction`.`master_status_transaction`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_transaction`.`master_status_transaction` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `status` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tublessin_transaction`.`transaction_history`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_transaction`.`transaction_history` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `id_montir` INT NOT NULL,
  `id_user` INT NOT NULL,
  `montir_firstname` VARCHAR(45) NOT NULL,
  `user_firstname` VARCHAR(45) NOT NULL,
  `master_status_transaction_id` INT NOT NULL DEFAULT 1,
  `date_created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `fk_transaction_history_master_status_transaction1_idx` (`master_status_transaction_id` ASC) VISIBLE,
  CONSTRAINT `fk_transaction_history_master_status_transaction1`
    FOREIGN KEY (`master_status_transaction_id`)
    REFERENCES `tublessin_transaction`.`master_status_transaction` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tublessin_transaction`.`transaction_location`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_transaction`.`transaction_location` (
  `transaction_history_id` INT NOT NULL,
  `latitude` DOUBLE NOT NULL,
  `longitude` DOUBLE NOT NULL,
  PRIMARY KEY (`transaction_history_id`),
  CONSTRAINT `fk_transaction_location_transaction_history1`
    FOREIGN KEY (`transaction_history_id`)
    REFERENCES `tublessin_transaction`.`transaction_history` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

USE `tublessin_montir` ;

-- -----------------------------------------------------
-- Placeholder table for view `tublessin_montir`.`montir_rating_location_view`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_montir`.`montir_rating_location_view` (`id` INT, `status_account` INT, `firstname` INT, `lastname` INT, `imageURL` INT, `status_operational` INT, `status` INT, `latitude` INT, `longitude` INT, `date_updated` INT, `total_rating` INT, `average_rating` INT);

-- -----------------------------------------------------
-- Placeholder table for view `tublessin_montir`.`overview_montir_view`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_montir`.`overview_montir_view` (`id` INT, `firstname` INT, `lastname` INT, `imageURL` INT, `gender` INT, `phone_number` INT, `city` INT, `username` INT, `status_account` INT, `verified_account` INT, `status_operational` INT, `status` INT, `total_rating` INT, `average_rating` INT);

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

-- -----------------------------------------------------
-- View `tublessin_montir`.`overview_montir_view`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `tublessin_montir`.`overview_montir_view`;
USE `tublessin_montir`;
CREATE  OR REPLACE VIEW `overview_montir_view` AS
SELECT ma.id,mp.firstname, mp. lastname, mp.imageURL,mp.gender,mp.phone_number,mp.city,ma.username,ma.status_account,mp.verified_account, ms.status_operational, msa.status, 
 mr.total_rating , mr.average_rating
FROM montir_account ma 
JOIN montir_status ms ON ma.id = ms.montir_account_id
JOIN master_status_activity msa ON ms.status_activity_id = msa.id
JOIN montir_location ml ON ma.id = ml.montir_account_id
JOIN montir_profile mp ON ma.id = mp.montir_account_id
JOIN (SELECT montir_account_id as id, count(montir_account_id) as total_rating, AVG(rating) average_rating FROM montir_rating 
GROUP BY montir_account_id) mr
ON mr.id = ma.id;
USE `tublessin_user` ;

-- -----------------------------------------------------
-- Placeholder table for view `tublessin_user`.`overview_user_view`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_user`.`overview_user_view` (`id` INT, `firstname` INT, `lastname` INT, `gender` INT, `phone_number` INT, `email` INT, `imageURL` INT, `date_updated` INT, `date_created` INT, `username` INT, `status_account` INT);

-- -----------------------------------------------------
-- View `tublessin_user`.`overview_user_view`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `tublessin_user`.`overview_user_view`;
USE `tublessin_user`;
CREATE  OR REPLACE VIEW `overview_user_view` AS
select ua.id,up.firstname, up.lastname,up.gender,up.phone_number,up.email,up.imageURL,up.date_updated,up.date_created, ua.username,ua.status_account from user_profile up JOIN
user_account ua ON up.user_account_id = ua.id;
USE `tublessin_transaction` ;

-- -----------------------------------------------------
-- Placeholder table for view `tublessin_transaction`.`transaction_history_view`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_transaction`.`transaction_history_view` (`id` INT, `id_montir` INT, `id_user` INT, `montir_firstname` INT, `user_firstname` INT, `status` INT, `date_created` INT, `latitude` INT, `longitude` INT);

-- -----------------------------------------------------
-- View `tublessin_transaction`.`transaction_history_view`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `tublessin_transaction`.`transaction_history_view`;
USE `tublessin_transaction`;
CREATE  OR REPLACE VIEW `transaction_history_view` AS
SELECT th.id,th.id_montir,th.id_user,th.montir_firstname,th.user_firstname,mst.status,th.date_created, tl.latitude,tl.longitude 
FROM transaction_history th JOIN master_status_transaction mst ON
th.master_status_transaction_id = mst.id JOIN transaction_location tl ON
th.id = tl.transaction_history_id;

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
