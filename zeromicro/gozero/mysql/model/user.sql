CREATE TABLE `user` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `name` varchar(255) NOT NULL DEFAULT '' COMMENT 'The username',
                        `password` varchar(255) NOT NULL DEFAULT '' COMMENT 'The user password',
                        `mobile` varchar(255) NOT NULL DEFAULT '' COMMENT 'The mobile phone number',
                        `gender` char(10) NOT NULL DEFAULT 'male' COMMENT 'gender,male|female|unknown',
                        `nickname` varchar(255) DEFAULT '' COMMENT 'The nickname',
                        `type` tinyint(1) DEFAULT '0' COMMENT 'The user type, 0:normal,1:vip, for test golang keyword',
                        `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `mobile_index` (`mobile`),
                        UNIQUE KEY `name_index` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='user table';