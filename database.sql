CREATE TABLE IF NOT EXISTS `customer` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `username` varchar(255) UNIQUE,
  `password` varchar(255),
  `full_name` varchar(255),
  `email` varchar(255) UNIQUE,
  `created_at` timestamp,
  `updated_at` datetime
);

CREATE TABLE IF NOT EXISTS  `category` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `category_name` varchar(255) NOT NULL,
  `created_at` timestamp,
  `updated_at` datetime
);

CREATE TABLE IF NOT EXISTS  `product` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `category_id` int,
  `product_name` varchar(255),
  `product_image` varchar(255),
  `product_description` longtext,
  `price` float,
  `created_at` timestamp,
  `updated_at` datetime
);

CREATE TABLE IF NOT EXISTS  `transaction` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `id_product` int,
  `id_customer` int,
  `qty` int,
  `amount` float,
  `status` int
  `created_at` timestamp,
  `updated_at` datetime
);

CREATE TABLE IF NOT EXISTS  `category_product` (
  `category_id` int,
  `product_category_id` int,
  PRIMARY KEY (`category_id`, `product_category_id`)
);

-- Insert random data into `customer` table
INSERT INTO `customer` (`username`, `password`, `full_name`, `email`, `created_at`, `updated_at`)
VALUES
    ('johndoe', 'password123', 'John Doe', 'john.doe@example.com', NOW(), NOW()),
    ('janedoe', 'securepassword', 'Jane Doe', 'jane.doe@example.com', NOW(), NOW()),
    ('michaelb', 'password456', 'Michael Brown', 'michael.brown@example.com', NOW(), NOW()),
    ('emilyw', 'password789', 'Emily White', 'emily.white@example.com', NOW(), NOW()),
    ('davidj', 'password321', 'David Johnson', 'david.johnson@example.com', NOW(), NOW());

-- Insert random data into `category` table
INSERT INTO `category` (`category_name`, `created_at`, `updated_at`)
VALUES
    ('Electronics', NOW(), NOW()),
    ('Books', NOW(), NOW()),
    ('Clothing', NOW(), NOW()),
    ('Sports', NOW(), NOW()),
    ('Home & Garden', NOW(), NOW());

-- Insert random data into `product` table
INSERT INTO `product` (`category_id`, `product_name`, `product_image`, `product_description`, `price`, `created_at`, `updated_at`)
VALUES
    (1, 'Smartphone', 'smartphone.jpg', 'Latest model smartphone with all the new features.', 699.99, NOW(), NOW()),
    (2, 'Novel', 'novel.jpg', 'A captivating novel that keeps you hooked until the end.', 19.99, NOW(), NOW()),
    (3, 'T-shirt', 'tshirt.jpg', 'Comfortable cotton t-shirt available in multiple sizes.', 14.99, NOW(), NOW()),
    (4, 'Basketball', 'basketball.jpg', 'Official size basketball for outdoor play.', 29.99, NOW(), NOW()),
    (5, 'Lawn Mower', 'lawn_mower.jpg', 'Electric lawn mower with adjustable cutting height.', 249.99, NOW(), NOW());

-- Insert random data into `category_product` table
INSERT INTO `category_product` (`category_id`, `product_category_id`)
VALUES
    (1, 1),
    (2, 2),
    (3, 3),
    (4, 4),
    (5, 5);

ALTER TABLE `category_product` ADD FOREIGN KEY (`category_id`) REFERENCES `category` (`id`);

ALTER TABLE `category_product` ADD FOREIGN KEY (`product_category_id`) REFERENCES `product` (`id`);

ALTER TABLE `transaction` ADD FOREIGN KEY (`id_product`) REFERENCES `product` (`id`);

ALTER TABLE `transaction` ADD FOREIGN KEY (`id_customer`) REFERENCES `customer` (`id`);
