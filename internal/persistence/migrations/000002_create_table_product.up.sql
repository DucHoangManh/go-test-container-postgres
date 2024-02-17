create table product
(
    id    serial8      not null primary key,
    name  varchar(100) not null,
    type  varchar(50)  not null,
    code  varchar(50)  not null,
    price int4         not null
);