CREATE DATABASE  IF NOT EXISTS `my_db` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `my_db`;
-- MySQL dump 10.13  Distrib 8.0.22, for Linux (x86_64)
--
-- Host: 149.28.143.216    Database: my_db
-- ------------------------------------------------------
-- Server version	5.7.32

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `courses`
--

DROP TABLE IF EXISTS `courses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `courses` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `code` varchar(45) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `courses`
--

LOCK TABLES `courses` WRITE;
/*!40000 ALTER TABLE `courses` DISABLE KEYS */;
INSERT INTO `courses` VALUES (1,'Course 101',1,'101',1),(2,'Course 102',1,'102',NULL),(6,'Course 106',1,'106',NULL),(7,'Course 107',1,'107',NULL),(8,'Course 108',1,'108',NULL),(11,'Course 111',1,'111',0),(12,'Course 112',1,'112',0),(13,'Course 113',1,'113',0),(14,'Course 114',1,'114',1),(15,'Course 115',1,'115',0),(16,'Course 116',1,'116',0),(23,'Course 123',1,'123',0),(24,'Course 124',1,'124',0),(25,'Course 125',1,'125',0),(26,'Course 126',1,'126',0),(27,'Course 127',1,'127',0),(28,'Course 128',1,'128',0),(29,'Course 129',1,'Course 129',0),(30,'Course 30',1,'000030',0),(31,'Course 31',1,'000031',0),(32,'Course 32',1,'000032',0),(33,'Course 33',1,'000033',0),(34,'Course 34',1,'000034',0),(35,'Course 35',1,'000035',0),(36,'Course 36',1,'000036',0);
/*!40000 ALTER TABLE `courses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `keys`
--

DROP TABLE IF EXISTS `keys`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `keys` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `key` varchar(45) DEFAULT NULL,
  `status` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `keys`
--

LOCK TABLES `keys` WRITE;
/*!40000 ALTER TABLE `keys` DISABLE KEYS */;
INSERT INTO `keys` VALUES (1,'c4304d1efcd9bfa812780fbcdd8b7ccf','1'),(2,'48f4be4548e70ee14ba3d884e29f9f1b','1'),(3,'3c0504482629d3b1ae7dfcdc7b469a53','1'),(8,'79ca7652b3ea24a5cf5444b56e48f041','1'),(9,'cc800869635c0b5119c6b7d3294dbf50','1'),(11,'59875c0e4981506b84c15b14f5d6aeb3','1'),(12,'812cea4186868f003f2cc0ac69fb46a6','1'),(13,'22d07e1d8e2ca7bcc41ea040f04ae793','1'),(14,'78fc5d3dd1ea446f3ea40edd1a350461','1'),(15,'da921d48e83257c7a21f621d8b6c59b1','1');
/*!40000 ALTER TABLE `keys` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(45) DEFAULT NULL,
  `last_name` varchar(45) DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'User','One',5),(2,'User','Two',5),(3,'User','Three',5),(23,'User','TwentyThree',50),(24,'User','TwentyFour',50),(25,'User','TwentyFive',50),(26,'User','TwentySix',50),(27,'User','TwentySeven',50),(28,'User','TwentyEight',52),(29,'User','TwentyNine',53),(30,'User','Thirty',53),(31,'User','ThirtyOne',53),(32,'User','ThirtyTwo',50),(33,'User','ThirtyThree',50),(34,'User','ThirtyFour',50),(35,'User','ThirtyFive',50),(36,'User','ThirtySix',50);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-11-20 10:22:30
