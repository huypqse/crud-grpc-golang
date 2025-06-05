CREATE TABLE `user`
(
    `id`        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'User ID',
    `passport`  varchar(45) NOT NULL COMMENT 'User Passport',
    `password`  varchar(45) NOT NULL COMMENT 'User Password',
    `nickname`  varchar(45) NOT NULL COMMENT 'User Nickname',
    `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
    `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS recharge_log (
    id          INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Recharge ID',
    user_id     int(10) unsigned NOT NULL COMMENT 'User ID',
    amount      DECIMAL(18,2) NOT NULL COMMENT 'Recharge amount',
    `before`    DECIMAL(18,2) NOT NULL COMMENT 'Balance before recharge',
    `after`     DECIMAL(18,2) NOT NULL COMMENT 'Balance after recharge',
    created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    remark      VARCHAR(255) DEFAULT NULL COMMENT 'Remark',
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Recharge log table';