
CREATE TABLE IF NOT EXISTS  `user` (
    `id` INTEGER PRIMARY KEY,
    `name` TEXT NOT NULL,
    `email` TEXT NOT NULL UNIQUE,
    `password` TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS `list` (
    `id` INTEGER PRIMARY KEY,
    `title` TEXT NOT NULL,

    `user_id` INTEGER NOT NULL,

    FOREIGN KEY(`user_id`) REFERENCES `user`(`id`)
);

CREATE TABLE IF NOT EXISTS `list_item` (
    `id` INTEGER PRIMARY KEY,
    `title` TEXT NOT NULL,

    `user_id` INTEGER NOT NULL,
    `list_id` INTEGER NOT NULL,

    FOREIGN KEY(`user_id`) REFERENCES `user`(`id`),
    FOREIGN KEY(`list_id`) REFERENCES `list`(`id`)
);

CREATE TABLE IF NOT EXISTS `comment` (
    `id` INTEGER PRIMARY KEY,
    `message` TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS `comment_list_item` (
    `id` INTEGER PRIMARY KEY,
    `comment_id` INTEGER NOT NULL,
    `list_item_id` INTEGER NOT NULL,

    FOREIGN KEY(`comment_id`) REFERENCES `comment`(`id`),
    FOREIGN KEY(`list_item_id`) REFERENCES `list_item`(`id`)
);
