CREATE TABLE `users` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `username` VARCHAR(255) UNIQUE,
  `hashed_password` VARCHAR(255) NOT NULL,
  `full_name` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) UNIQUE,
  `password_changed_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  -- Additional columns if needed
);

CREATE TABLE `rent` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `bookName` VARCHAR(255),
  `user_id` INT, -- Assuming `user_name` is a foreign key referencing `users.id`
  `days_due` INT,
  `fine` INT,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  -- Additional columns if needed
);

CREATE TABLE `Books` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255),
  `author` VARCHAR(255),
  `quantity` INT,
  `ISBN` INT,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  -- Additional columns if needed
);

ALTER TABLE `rent` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

