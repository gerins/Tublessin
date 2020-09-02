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
-- Schema tublessin_chat
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema tublessin_chat
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `tublessin_chat` ;
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


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
