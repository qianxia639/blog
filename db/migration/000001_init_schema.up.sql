CREATE TABLE `users` (
    `id` bigint PRIMARY KEY AUTO_INCREMENT,
    `username` varchar(20) UNIQUE NOT NULL COMMENT '用户名',
    `email` varchar(30) UNIQUE NOT NULL COMMENT '用户邮箱',
    `password` varchar(210)  NOT NULL COMMENT '用户密码',
    `nickname` varchar(20) UNIQUE NOT NULL COMMENT '用户昵称',
    `avatar` varchar(255) NOT NULL DEFAULT 'default.jpg' COMMENT '用户头像'
);