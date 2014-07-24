
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `user` (
    `uid` varchar(48) NOT NULL PRIMARY KEY,
    `username` varchar(255) NOT NULL UNIQUE,
    `email` varchar(255) NOT NULL UNIQUE,
    `password` varchar(255) NOT NULL,
    `verified` bool NOT NULL DEFAULT FALSE,
    `enabled` bool NOT NULL DEFAULT FALSE, 
    `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `last_login` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP    
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `user`;
