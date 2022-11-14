CREATE TABLE `students` (
  `student_id` int unsigned NOT NULL AUTO_INCREMENT,
  `student_name` varchar(45) NOT NULL,
  `student_email` varchar(45) NOT NULL,
  `student_phone` varchar(45) DEFAULT NULL,
  `student_pass` varchar(45) NOT NULL,
  PRIMARY KEY (`student_id`),
  UNIQUE KEY `id_UNIQUE` (`student_id`),
  UNIQUE KEY `student_email_UNIQUE` (`student_email`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
