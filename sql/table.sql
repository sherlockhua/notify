CREATE TABLE `task` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `task_id` bigint unsigned  NOT NULL COMMENT '唯一任务ID',
    `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
    `task_type` tinyint NOT NULL COMMENT '任务类型',
    `task_desc` varchar(1024) NOT NULL COMMENT '任务描述',
    `task_name` varchar(256)  NOT NULL DEFAULT '' COMMENT '任务名称',
    `task_status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
    `notify_channel` tinyint NOT NULL DEFAULT '0' COMMENT '通知渠道',
    `strategy_data` varchar(1024) NOT NULL DEFAULT '' COMMENT '策略数据',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_idx_task_id` (`task_id`),
    KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='任务表';

CREATE TABLE `task_template` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `template_id` bigint unsigned  NOT NULL COMMENT '唯一任务ID',
    `template_name` varchar(256)  NOT NULL DEFAULT '' COMMENT '任务名称',
    `template_desc` varchar(1024) NOT NULL COMMENT '任务模板描述',
    `task_type` tinyint NOT NULL COMMENT '任务类型',
    `notify_channel` tinyint NOT NULL DEFAULT '0' COMMENT '通知渠道',
    `template_status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
    `strategy_data` varchar(1024) NOT NULL DEFAULT '' COMMENT '策略数据',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_idx_task_id` (`template_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='任务表';
