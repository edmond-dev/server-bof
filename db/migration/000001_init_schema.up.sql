SET FOREIGN_KEY_CHECKS = 1;

CREATE TABLE IF NOT EXISTS `customers` (
     `customer_id` varchar(255) UNIQUE PRIMARY KEY NOT NULL,
     `first_name` varchar(255) NOT NULL,
     `last_name` varchar(255) NOT NULL,
     `email` varchar(255) UNIQUE NOT NULL,
     `password` blob NOT NULL,
     `created_at` timestamp default current_timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS `categories` (
      `category_id` varchar(255) UNIQUE PRIMARY KEY NOT NULL,
      `category_name` varchar(255) UNIQUE NOT NULL,
      `category_html_description` text NOT NULL ,
      `image` varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS `products` (
    `product_id` varchar(255) UNIQUE PRIMARY KEY NOT NULL,
    `category_id` varchar(255) NOT NULL ,
    `image_url_public_id` varchar(255) NOT NULL,
    `image_url_secure_id` varchar(255) NOT NULL,
    `product_name` varchar(255) NOT NULL,
    `product_description` varchar(255) NOT NULL,
    `price` bigint NOT NULL,
    `quantity_in_stock` varchar(20) NOT NULL ,
    `created_at` timestamp default current_timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS `reviews` (
   `review_id` varchar(255) UNIQUE PRIMARY KEY NOT NULL ,
   `product_review_id` varchar(255) UNIQUE NOT NULL,
   `customer_review_id` varchar(255) NOT NULL ,
   `review` text NOT NULL,
   `created_at` timestamp default current_timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS `orders` (
      `order_number` varchar(255) UNIQUE PRIMARY KEY NOT NULL,
      `customer_id`  varchar(255) NOT NULL ,
      `product_id` varchar(255) NOT NULL ,
      `order_category_name` varchar(255) NOT NULL,
      `quantity_ordered` int not null,
      `price_each` int not null,
      `customer_comments` varchar(255) NOT NULL,
      `ordered_date` timestamp default current_timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS `orderDetails` (
    `order_number` varchar(255) UNIQUE PRIMARY KEY NOT NULL,
    `shipped_date` varchar(255) NOT NULL ,
    `status` varchar(255) NOT NULL,
    `comments` varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS `payment` (
   `customer_id` varchar(255) UNIQUE NOT NULL,
   `payment_date` timestamp NOT NULL,
   `payment_method_id` varchar(255) NOT NULL,
   `amount` bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS `address` (
   `address_customer_id` varchar(255) UNIQUE NOT NULL,
   `email` varchar(255) NOT NULL,
   `street` varchar(255) NOT NULL,
   `address_line_1` varchar(255) NOT NULL,
   `address_line_2` varchar(255) NOT NULL,
   `phone` varchar(255) NOT NULL ,
   `city` varchar(255) NOT NULL,
   `state` varchar(255) NOT NULL,
   `zip_code` varchar(255) NOT NULL
);

ALTER TABLE `products` ADD FOREIGN KEY (`category_id`) REFERENCES `categories` (`category_id`);

ALTER TABLE `orders` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`product_id`);

ALTER TABLE `orders` ADD FOREIGN KEY (`customer_id`) REFERENCES `customers` (`customer_id`);

ALTER TABLE `reviews` ADD FOREIGN KEY (`product_review_id`) REFERENCES `products` (`product_id`);

ALTER TABLE `reviews` ADD FOREIGN KEY (`customer_review_id`) REFERENCES `customers` (`customer_id`);

ALTER TABLE `orderDetails` ADD FOREIGN KEY (`order_number`) REFERENCES `orders` (`order_number`);

ALTER TABLE `payment` ADD FOREIGN KEY (`customer_id`) REFERENCES `customers` (`customer_id`);

ALTER TABLE `address` ADD FOREIGN KEY (`address_customer_id`) REFERENCES `customers` (`customer_id`);

 CREATE INDEX `customers_index_0` ON `customers` (`email`);

 CREATE INDEX `products_index_1` ON `products` (`product_name`);

 CREATE INDEX `reviews_index_2` ON `reviews` (`customer_review_id`);

 CREATE INDEX `orders_index_3` ON `orders` (`ordered_date`);

 CREATE INDEX `address_index_4` ON `address` (`zip_code`);

