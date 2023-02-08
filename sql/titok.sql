CREATE TABLE `user`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户id唯一表示',
  `name` varchar(255) NOT NULL COMMENT '用户名称',
  `password` varchar(255) NOT NULL COMMENT '用户密码',
  `follow_count` int NULL COMMENT '关注数量',
  `follower_count` int NULL COMMENT '粉丝数量',
  `create_at` timestamp NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '用户注册时间',
  PRIMARY KEY (`id`)
);

CREATE TABLE `video`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
  `paly_url` varchar(255) NOT NULL COMMENT '视频播放地址',
  `cover_url` varchar(255) NULL COMMENT '视频封面地址',
  `favorite_count` int NULL COMMENT '视频点赞数量',
  `comment_count` int NULL COMMENT '视频评论数量',
  `title` varchar(255) NULL COMMENT '视频标题',
  `create_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '视频创建时间',
  `update_at` timestamp NULL COMMENT '视频更新时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '视频删除时间',
  PRIMARY KEY (`id`)
);

