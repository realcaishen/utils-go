CREATE TABLE `t_token_info` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `update_timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `insert_timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `token_name` varchar(128) NOT NULL,
    `chain_name` varchar(64) NOT NULL,
    `token_address` varchar(128) NOT NULL,
    `decimals` int NOT NULL,
    `full_name` varchar(128) NOT NULL DEFAULT '',
    `total_supply` DECIMAL(64, 0) NOT NULL DEFAULT 0,
    `discover_timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `icon` varchar(1024) NOT NULL DEFAULT '',
    `mcap` DOUBLE NOT NULL DEFAULT 0,
    `fdv` DOUBLE NOT NULL DEFAULT 0,
    `volume24h` DOUBLE NOT NULL DEFAULT 0,
    `pricechg24h` DOUBLE NOT NULL DEFAULT 0,
    `pricechg6h` DOUBLE NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_chain_name_token_address` (`chain_name`, `token_address`),
    KEY `idx_token_name` (`token_name`),
    KEY `idx_insert_timestamp` (`insert_timestamp`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `t_tag` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `update_timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `insert_timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `tag_name` varchar(128) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_tag_name` (`tag_name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `t_object_tag` (
    `object_table` varchar(64),
    `object_id` bigint NOT NULL,
    `tag_id` bigint NOT NULL,
    `update_timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `insert_timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`object_table`, `tag_id`, `object_id`),
    KEY `idx_object` (`object_table`, `object_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;