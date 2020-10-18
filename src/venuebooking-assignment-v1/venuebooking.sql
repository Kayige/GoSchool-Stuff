--
-- Dumping data for table `bookings`
--
DROP TABLE IF EXISTS `bookings`;

--
-- Table structure for table `bookings`
--
CREATE TABLE `bookings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `v_name` varchar(30) DEFAULT NULL,
  `st` varchar(50) DEFAULT NULL,
  `et` varchar(50) DEFAULT NULL,
  `booked_by` int(11) NOT NULL,
  `customer` varchar(30) DEFAULT NULL,
  `phone` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `booked_by` (`booked_by`),
  CONSTRAINT `bookings_ibfk_1` FOREIGN KEY (`booked_by`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--
DROP TABLE IF EXISTS `users`;

--
-- Table structure for table `users`
--
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `fname` varchar(20) DEFAULT NULL,
  `lname` varchar(20) DEFAULT NULL,
  `email` varchar(50) NOT NULL,
  `password` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `role` int(2) NOT NULL DEFAULT '2',
  `session` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

/* INSERTING ADMIN DATA INTO table `users` */;
INSERT INTO `users` VALUES (1,'Admin','User','admin@gmail.com','$P$BwQZDcQaNbAWNrROacEREiu4BLty1N1','2020-10-13 11:55:25',1,'');

--
-- Dumping data for table `venues`
--
DROP TABLE IF EXISTS `venues`;
--
-- Table structure for table `venues`
--
CREATE TABLE `venues` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `img` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

