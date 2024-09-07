CREATE TABLE `customer` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `username` varchar(255) UNIQUE,
  `password` varchar(255),
  `full_name` varchar(255),
  `email` varchar(255) UNIQUE,
  `created_at` timestamp,
  `updated_at` datetime
);

CREATE TABLE `category` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `category_name` varchar(255) NOT NULL,
  `created_at` timestamp,
  `updated_at` datetime
);

CREATE TABLE `product` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `category_id` int,
  `product_name` varchar(255),
  `product_image` varchar(255),
  `product_description` longtext,
  `price` float,
  `created_at` timestamp,
  `updated_at` datetime
);

CREATE TABLE `transaction` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `id_product` int,
  `id_customer` int,
  `qty` int,
  `amount` float,
  `created_at` timestamp,
  `updated_at` datetime
);

CREATE TABLE `category_product` (
  `category_id` int,
  `product_category_id` int,
  PRIMARY KEY (`category_id`, `product_category_id`)
);

ALTER TABLE `category_product` ADD FOREIGN KEY (`category_id`) REFERENCES `category` (`id`);

ALTER TABLE `category_product` ADD FOREIGN KEY (`product_category_id`) REFERENCES `product` (`category_id`);

ALTER TABLE `transaction` ADD FOREIGN KEY (`id_product`) REFERENCES `product` (`id`);

ALTER TABLE `transaction` ADD FOREIGN KEY (`id_customer`) REFERENCES `customer` (`id`);
