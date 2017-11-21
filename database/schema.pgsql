CREATE TABLE clients (
    id serial primary key not null check(id > 0),
    login text not null check(length(login) > 0),
    password text not null check(length(password) > 0),
    telephone text not null check(length(telephone) > 0),
    birthday date not null default now(),
    fullname text not null check(length(fullname) > 0),
);

CREATE TABLE clients ( id serial primary key not null check(id > 0), login text not null check(length(login) > 0), password text not null check(length(password) > 0), telephone text not null check(length(telephone) > 0), birthday date not null default now(), fullname text not null check(length(fullname) > 0));


CREATE TABLE products (
    id serial primary key not null check(id > 0),
    type text not null check(type in ('pizza', 'sushi', 'burger', 'bakery', 'dessert')),
    cost numeric(7, 2) not null,
    name text not null check(length(name) > 0),
    weight int not null 
);

CREATE TABLE products (     id serial primary key not null check(id > 0),     type text not null check(type in ('pizza', 'sushi', 'burger', 'bakery', 'dessert')),     cost numeric(7, 2) not null,     name text not null check(length(name) > 0),     weight int not null  );

CREATE TABLE orders (
    id serial primary key not null check(id > 0),
    client int not null references clients(id) on update cascade,
    order_time timestamp not null
);

CREATE TABLE orders (     id serial primary key not null check(id > 0),     client int not null references clients(id) on update cascade,     order_time timestamp not null );

CREATE TABLE products_in_orders (
    order_id int not null references orders(id) on update cascade,
    product_id int not null references products(id) on update cascade,
    count int not null check(count > 0)
);

 CREATE TABLE products_in_orders (     order_id int not null references orders(id) on update cascade,     product_id int not null references products(id) on update cascade,     count int not null check(count > 0) );

(SELECT id, type, cost, name, weight  FROM products   WHERE id in ( SELECT product_id FROM products_in_orders WHERE order_id = 1 )) INNER JOIN (SELECT product_id, count FROM products_in_orders WHERE order_id = 1) on id = product_id;
