
SET FOREIGN_KEY_CHECKS = 1;

CREATE TABLE IF NOT EXISTS `users` (
    `id` bigint PRIMARY KEY,
    `user_id` varchar(255) UNIQUE NOT NULL,
    `first_name` varchar(255) NOT NULL,
    `last_name` varchar(255) NOT NULL,
    `email` varchar(255) UNIQUE NOT NULL,
    `password` varchar(255) NOT NULL,
    `created_at` datetime
    );

CREATE TABLE IF NOT EXISTS `categories` (
    `id` bigint PRIMARY KEY,
    `category_id` varchar(255) UNIQUE NOT NULL,
    `category_name` varchar(255) UNIQUE NOT NULL
    );

CREATE TABLE IF NOT EXISTS `products` (
    `id` bigint PRIMARY KEY,
    `product_id` varchar(255) UNIQUE NOT NULL,
    `image_url_public_id` varchar(255) NOT NULL,
    `image_url_secure_id` varchar(255) NOT NULL,
    `product_name` varchar(255) NOT NULL,
    `product_description` varchar(255) NOT NULL,
    `price` bigint NOT NULL,
    `created_at` datetime
    );

CREATE TABLE IF NOT EXISTS `reviews` (
   `id` bigint PRIMARY KEY,
   `review_id` varchar(255) UNIQUE NOT NULL,
    `user_review_id` varchar(255),
    `review` text NOT NULL,
    `created_at` datetime
    );

CREATE TABLE IF NOT EXISTS `orders` (
    `id` bigint PRIMARY KEY,
    `order_id` varchar(255) UNIQUE NOT NULL,
    `quantity` int DEFAULT 1,
    `total_amount` bigint,
    `created_at` varchar(255) COMMENT 'When order created'
    );

CREATE TABLE IF NOT EXISTS `address` (
   `id` bigint PRIMARY KEY,
   `address_user_id` varchar(255) UNIQUE NOT NULL,
    `order_id` varchar(255),
    `email` varchar(255) NOT NULL,
    `street` varchar(255) NOT NULL,
    `address_line_1` varchar(255) NOT NULL,
    `address_line_2` varchar(255) NOT NULL,
    `city` varchar(255) NOT NULL,
    `state` varchar(255) NOT NULL,
    `zip_code` varchar(255) NOT NULL
    );

ALTER TABLE `products` ADD FOREIGN KEY (`product_id`) REFERENCES `categories` (`category_id`);

ALTER TABLE `reviews` ADD FOREIGN KEY (`review_id`) REFERENCES `products` (`product_id`);

ALTER TABLE `reviews` ADD FOREIGN KEY (`user_review_id`) REFERENCES `users` (`user_id`);

ALTER TABLE `orders` ADD FOREIGN KEY (`order_id`) REFERENCES `products` (`product_id`);

ALTER TABLE `address` ADD FOREIGN KEY (`address_user_id`) REFERENCES `users` (`user_id`);

ALTER TABLE `address` ADD FOREIGN KEY (`order_id`) REFERENCES `orders` (`order_id`);

CREATE INDEX `users_index_0` ON `users` (`user_id`);

CREATE INDEX `products_index_1` ON `products` (`product_id`);

CREATE INDEX `reviews_index_2` ON `reviews` (`review_id`);

CREATE INDEX `orders_index_3` ON `orders` (`order_id`);

CREATE INDEX `address_index_4` ON `address` (`address_user_id`);
