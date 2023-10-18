-- Create database myblog.
CREATE DATABASE `myblog` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

-- myblog.BlogPosts definition
CREATE TABLE `BlogPosts` (
  `PostID` int NOT NULL AUTO_INCREMENT,
  `Title` varchar(255) NOT NULL,
  `Content` text NOT NULL,
  `Author` varchar(100) NOT NULL,
  `CreatedAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `UpdatedAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`PostID`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- myblog.`User` definition
CREATE TABLE `User` (
  `UserID` int NOT NULL AUTO_INCREMENT,
  `Username` varchar(255) NOT NULL,
  `Password` varchar(255) NOT NULL,
  PRIMARY KEY (`UserID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- myblog.`User` sample user data 
INSERT INTO myblog.`User` (Username,Password) VALUES
	 ('jprakash','6e624291d9ae74b0d3811ba1a7860095');

-- myblog.`BlogPosts` sample data
INSERT INTO myblog.BlogPosts (Title,Content,Author,CreatedAt,UpdatedAt) VALUES
	 ('The Art of Cooking','In this blog post, we explore the culinary world and discover the secrets of creating delicious dishes that tantalize the taste buds.','Jane Smith','2023-10-17 20:12:05','2023-10-17 20:12:05'),
	 ('Exploring the Wonders of Nature','Join us on an adventure into the heart of nature''s beauty.In this blog post, we''ll share breathtaking photos and stories of our recent travels to exotic locations','John Doe','2023-10-17 20:13:08','2023-10-17 20:13:08'),
	 ('Exploring the Wonders of Nature','Join us on an adventure into the heart of nature''s beauty. In this blog post, we''ll share breathtaking photos and stories of our recent travels to exotic locations','John Doe','2023-10-18 07:23:40','2023-10-18 07:23:40'),
	 ('The Art of Cooking','In the culinary world, cooking is both a science and an art. From mastering the perfect soufflé to creating mouthwatering dishes that tantalize the taste buds, cooking is a journey of creativity and precision','Jane Smith','2023-10-18 07:24:57','2023-10-18 07:24:57'),
	 ('The Art of Cooking by Michel','In the culinary world, cooking is both a science and an art. From mastering the perfect soufflé to creating mouthwatering dishes that tantalize the taste buds, cooking is a journey of creativity and precision','Michel','2023-10-18 07:25:19','2023-10-18 07:25:19'),
	 ('A Journey Through History','Travel back in time with us as we explore the rich history of ancient civilizations. From the majestic pyramids of Egypt to the grandeur of the Roman Empire, this blog post is a virtual time machine that takes you on a historical adventure','Michael Anderson','2023-10-18 07:27:18','2023-10-18 07:27:18'),
	 ('Discovering the Cosmos','The universe is a vast and mysterious place, and our thirst for knowledge knows no bounds. Join us as we embark on a cosmic journey through the galaxies, stars, and planets','Sarah Johnson','2023-10-18 07:27:51','2023-10-18 07:27:51'),
	 ('Discovering the Cosmos','The universe is a vast and mysterious place, and our thirst for knowledge knows no bounds. Join us as we embark on a cosmic journey through the galaxies, stars, and planets','Michael','2023-10-18 07:27:59','2023-10-18 07:27:59'),
	 ('Unveiling the World of Art','Art is a reflection of human expression and creativity. In this blog post, we delve into the world of art, from classical masterpieces to contemporary innovations','Michael','2023-10-18 07:30:12','2023-10-18 07:30:12'),
	 ('Unveiling the World of Art','Art is a reflection of human expression and creativity. In this blog post, we delve into the world of art, from classical masterpieces to contemporary innovations','David','2023-10-18 07:30:25','2023-10-18 07:30:25');
INSERT INTO myblog.BlogPosts (Title,Content,Author,CreatedAt,UpdatedAt) VALUES
	 ('Unveiling the World of Art','Art is a reflection of human expression and creativity. In this blog post, we delve into the world of art, from classical masterpieces to contemporary innovations','David Martinez','2023-10-18 07:30:33','2023-10-18 07:30:33');
