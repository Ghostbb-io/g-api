-- casbin_rule
drop table if exists casbin_rule
go

create table casbin_rule
(
    id    bigint identity
        primary key,
    ptype nvarchar(100),
    v0    nvarchar(100),
    v1    nvarchar(100),
    v2    nvarchar(100),
    v3    nvarchar(100),
    v4    nvarchar(100),
    v5    nvarchar(100)
)
go

create unique index idx_casbin_rule
    on casbin_rule (ptype, v0, v1, v2, v3, v4, v5)
go

insert into casbin_rule (ptype, v0, v1, v2, v3, v4, v5)
values ('g', '1', 'root', '', '', '', '');

insert into casbin_rule (ptype, v0, v1, v2, v3, v4, v5)
values ('p', 'user', '/v1/casbin', 'DELETE', '', '', '');

insert into casbin_rule (ptype, v0, v1, v2, v3, v4, v5)
values ('p', 'user', '/v1/casbin', 'GET', '', '', '');

insert into casbin_rule (ptype, v0, v1, v2, v3, v4, v5)
values ('p', 'user', '/v1/casbin/:role', 'GET', '', '', '');

insert into casbin_rule (ptype, v0, v1, v2, v3, v4, v5)
values ('p', 'user', '/v1/casbin/:role', 'POST', '', '', '');

insert into casbin_rule (ptype, v0, v1, v2, v3, v4, v5)
values ('p', 'user', '/v1/casbin/:role', 'PUT', '', '', '');

insert into casbin_rule (ptype, v0, v1, v2, v3, v4, v5)
values ('p', 'user', '/v1/role', 'GET', '', '', '');

insert into casbin_rule (ptype, v0, v1, v2, v3, v4, v5)
values ('p', 'user', '/v1/user/:id/roles', 'PUT', '', '', '');

insert into casbin_rule (ptype, v0, v1, v2, v3, v4, v5)
values ('p', 'user', '/v1/cache/db', 'DELETE', '', '', '');

-- sys_roles
drop table if exists sys_roles
go

create table sys_roles
(
    role       nvarchar(191) not null
        primary key,
    created_at datetimeoffset,
    updated_at datetimeoffset,
    deleted_at datetimeoffset,
    role_name  nvarchar(max),
    status     bit    default 1,
    data_scope bigint default 3,
    remark     nvarchar(max)
)
go

create index idx_sys_roles_deleted_at
    on sys_roles (deleted_at)
go

insert into sys_roles (role, created_at, updated_at, role_name, data_scope)
values ('root', getdate(), getdate(), 'manager', 0);

insert into sys_roles (role, created_at, updated_at, role_name, data_scope)
values ('user', getdate(), getdate(), 'common user', 0);

-- sys_roles
drop table if exists sys_roles
go

create table sys_user_role
(
    sys_role_role nvarchar(191) not null,
    sys_user_id   bigint        not null,
    primary key (sys_role_role, sys_user_id)
)
go

insert into sys_user_role (sys_role_role, sys_user_id)
values ('root', '1');

-- sys_users
drop table if exists sys_users
go

create table sys_users
(
    id         bigint identity
        primary key,
    created_at datetimeoffset,
    updated_at datetimeoffset,
    deleted_at datetimeoffset,
    uuid       nvarchar(191),
    username   nvarchar(191),
    password   nvarchar(max),
    nick_name  nvarchar(max),
    email      nvarchar(max),
    remark     nvarchar(max)
)
go

create index idx_sys_users_username
    on sys_users (username)
go

create index idx_sys_users_uuid
    on sys_users (uuid)
go

create index idx_sys_users_deleted_at
    on sys_users (deleted_at)
go

-- username: admin
-- password: 123456
insert into sys_users (created_at, updated_at, uuid, username, password, nick_name, email)
values (getdate(), getdate(), '9bc03cd0-4b9e-11ee-9206-047c16cd0333', 'admin', '$2a$10$73MqmZWWqkihjwnPlP4Tq.Yau4lNp1.RhI4zhG7a3ZIjwivzQtRVa', 'Manager', 'xxx@xxx.com');

