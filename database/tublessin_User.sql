-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema tublessin_montir
-- -----------------------------------------------------
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
USE `tublessin_user` ;

-- -----------------------------------------------------
-- Table `tublessin_user`.`user_account`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tublessin_user`.`user_account` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(15) NOT NULL DEFAULT '',
  `password` VARCHAR(300) NOT NULL DEFAULT '',
  `date_created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
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


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

insert into user_account(username, password) value
('reizdendy', '$2a$10$fmWS.ALzAB4RTi66KkyDTOxSmHZOODW11TRI6jnBKZ.aL/m.UsWV2'),
('sandro', '$2a$10$fmWS.ALzAB4RTi66KkyDTOxSmHZOODW11TRI6jnBKZ.aL/m.UsWV2');

insert into user_profile(user_account_id, firstname, lastname, gender, phone_number, email) value
(1, 'reizdendy', 'akbar', 'L', '08181818', 'reizdendy@gmail.com'),
(2, 'sandro', 'dorooo', 'L', '0818288', 'sandro@gmail.com');

insert into user_location(user_account_id) value(1),(2);