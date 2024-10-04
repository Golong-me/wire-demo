create table user
(
    id         bigint auto_increment comment 'id'
        primary key,
    name       varchar(15)   null comment '用户名称',
    phone      bigint        null comment '手机号',
    birth      datetime(3)   null comment '生日',
    created_at datetime(3)   null comment '创建时间',
    deleted_at int default 0 not null
)
    comment '用户表';

insert into user(name, phone, birth, created_at, deleted_at) values ('张三', 18888888888, '1999-01-01 00:00:00.000', '2023-01-01 00:00:00',0)