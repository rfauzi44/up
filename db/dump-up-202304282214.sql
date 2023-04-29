-- MySQL dump 10.13  Distrib 8.0.32, for Win64 (x86_64)
--
-- Host: localhost    Database: up
-- ------------------------------------------------------
-- Server version	8.0.32

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `like_review`
--

DROP TABLE IF EXISTS `like_review`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `like_review` (
  `id_review` varchar(100) NOT NULL,
  `id_member` varchar(100) NOT NULL,
  PRIMARY KEY (`id_review`,`id_member`),
  KEY `like_review_FK_1` (`id_member`),
  CONSTRAINT `like_review_FK` FOREIGN KEY (`id_review`) REFERENCES `review_product` (`id_review`),
  CONSTRAINT `like_review_FK_1` FOREIGN KEY (`id_member`) REFERENCES `member` (`id_member`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `like_review`
--

LOCK TABLES `like_review` WRITE;
/*!40000 ALTER TABLE `like_review` DISABLE KEYS */;
INSERT INTO `like_review` VALUES ('2a5a7d4c-0490-4de0-8f17-e5b03d555330','815fdeda-8ee1-4a87-bde9-d60a958646bc'),('2a5a7d4c-0490-4de0-8f17-e5b03d555330','a4dadc54-c8c0-43f8-8e39-e951b77ccb13'),('2a5a7d4c-0490-4de0-8f17-e5b03d555330','e6b3dfa1-3b61-432a-8bde-0def7f737a10');
/*!40000 ALTER TABLE `like_review` ENABLE KEYS */;
UNLOCK TABLES;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb3 */ ;
/*!50003 SET character_set_results = utf8mb3 */ ;
/*!50003 SET collation_connection  = utf8mb3_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'IGNORE_SPACE,ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`localhost`*/ /*!50003 TRIGGER `review_like` AFTER INSERT ON `like_review` FOR EACH ROW BEGIN
    UPDATE review_product SET like_count = (SELECT COUNT(*) FROM like_review WHERE id_review = NEW.id_review) WHERE id_review = NEW.id_review;
  END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb3 */ ;
/*!50003 SET character_set_results = utf8mb3 */ ;
/*!50003 SET collation_connection  = utf8mb3_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'IGNORE_SPACE,ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`localhost`*/ /*!50003 TRIGGER `review_dislike` AFTER DELETE ON `like_review` FOR EACH ROW BEGIN
    UPDATE review_product SET like_count = (SELECT COUNT(*) FROM like_review WHERE id_review = OLD.id_review) WHERE id_review = OLD.id_review;
  END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `member`
--

DROP TABLE IF EXISTS `member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `member` (
  `id_member` varchar(100) NOT NULL,
  `username` varchar(100) NOT NULL,
  `gender` varchar(100) NOT NULL,
  `skintype` varchar(100) NOT NULL,
  `skincolor` varchar(100) NOT NULL,
  PRIMARY KEY (`id_member`),
  UNIQUE KEY `member_un` (`id_member`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `member`
--

LOCK TABLES `member` WRITE;
/*!40000 ALTER TABLE `member` DISABLE KEYS */;
INSERT INTO `member` VALUES ('815fdeda-8ee1-4a87-bde9-d60a958646bc','johndoe','female','type','brown'),('a4dadc54-c8c0-43f8-8e39-e951b77ccb13','john','female','type','brown'),('e6b3dfa1-3b61-432a-8bde-0def7f737a10','doe','female','type','brown'),('f4b4c5e1-8973-4da4-bd3f-b2ef88ed5907','foo','female','type','brown');
/*!40000 ALTER TABLE `member` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product`
--

DROP TABLE IF EXISTS `product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product` (
  `id_product` varchar(100) NOT NULL,
  `name_product` varchar(100) NOT NULL,
  `price` int NOT NULL,
  PRIMARY KEY (`id_product`),
  UNIQUE KEY `product_un` (`id_product`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product`
--

LOCK TABLES `product` WRITE;
/*!40000 ALTER TABLE `product` DISABLE KEYS */;
INSERT INTO `product` VALUES ('5c4d96ad-ec09-4b6d-a8a3-f6715d452b63','product2',50000),('bb43a8c3-8892-47f3-8dc0-0978bd08477e','product1',50000),('c4b081e7-0bab-409b-bd8b-590fa0ac5af6','product3',50000);
/*!40000 ALTER TABLE `product` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `review_product`
--

DROP TABLE IF EXISTS `review_product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `review_product` (
  `id_review` varchar(100) NOT NULL,
  `id_product` varchar(100) NOT NULL,
  `id_member` varchar(100) NOT NULL,
  `desc_review` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `like_count` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id_review`),
  UNIQUE KEY `review_product_un` (`id_review`),
  KEY `review_product_FK` (`id_product`),
  KEY `review_product_FK_1` (`id_member`),
  CONSTRAINT `review_product_FK` FOREIGN KEY (`id_product`) REFERENCES `product` (`id_product`),
  CONSTRAINT `review_product_FK_1` FOREIGN KEY (`id_member`) REFERENCES `member` (`id_member`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `review_product`
--

LOCK TABLES `review_product` WRITE;
/*!40000 ALTER TABLE `review_product` DISABLE KEYS */;
INSERT INTO `review_product` VALUES ('2a5a7d4c-0490-4de0-8f17-e5b03d555330','c4b081e7-0bab-409b-bd8b-590fa0ac5af6','a4dadc54-c8c0-43f8-8e39-e951b77ccb13','I recently purchased this product and I have to say I am thoroughly impressed. The quality is top-notch and the design is both stylish and functional. The product exceeded my expectations and I would highly recommend it to anyone in need of a reliable solution.',3),('6340b79b-1050-46cf-9bef-81000d023354','c4b081e7-0bab-409b-bd8b-590fa0ac5af6','e6b3dfa1-3b61-432a-8bde-0def7f737a10','I was skeptical about this product at first, but I have to say I\'m pleasantly surprised. It performs exactly as advertised and the quality is impressive. The customer support team was also very responsive and helpful. Overall, a great experience.',0),('ac96b5f5-3381-4ab2-8dc0-cb3ce34976fb','c4b081e7-0bab-409b-bd8b-590fa0ac5af6','f4b4c5e1-8973-4da4-bd3f-b2ef88ed5907','I\'m very impressed with the quality of this product. It\'s clear that a lot of thought and attention to detail went into its design and production. The features are intuitive and easy to use, and the overall experience is great. Would definitely purchase again.',0),('c367b905-6070-471d-9840-d908c62cf3fc','c4b081e7-0bab-409b-bd8b-590fa0ac5af6','e6b3dfa1-3b61-432a-8bde-0def7f737a10','I\'ve been using this product for a few weeks now and it has completely changed my daily routine. The ease of use and durability of the product are both exceptional, and I love the sleek design. Highly recommend!',0),('c48bdc0b-67de-4a97-b20f-2fdd21115407','bb43a8c3-8892-47f3-8dc0-0978bd08477e','e6b3dfa1-3b61-432a-8bde-0def7f737a10','I\'m blown away by the performance of this product. It\'s fast, reliable, and has exceeded all of my expectations. The customer support team was also fantastic and provided quick and helpful responses to all of my inquiries. Highly recommend!',0),('f4e2b43a-b5ae-4196-a7c4-169c67ceb264','5c4d96ad-ec09-4b6d-a8a3-f6715d452b63','e6b3dfa1-3b61-432a-8bde-0def7f737a10','This product is a game-changer. It\'s versatile, easy to use, and has completely transformed the way I approach my work. The attention to detail in the design is apparent and the overall quality is exceptional. I\'m so glad I made this purchase',0);
/*!40000 ALTER TABLE `review_product` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'up'
--

--
-- Dumping routines for database 'up'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-04-28 22:14:35
