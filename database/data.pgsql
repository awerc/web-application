INSERT INTO clients (login, password, telephone, birthday, fullname) VALUES 
('login1', 'pass', '88005553535', '1991-01-22', 'client1'),
('login2', 'pass', '88005553535', '1992-04-03', 'client2'),
('login3', 'pass', '88005553535', '1993-07-12', 'client3'),
('login4', 'pass', '88005553535', '1994-02-08', 'client4'),
('login5', 'pass', '88005553535', '1995-12-23', 'client5'),
('login6', 'pass', '88005553535', '1996-10-11', 'client6');

INSERT INTO products (type, cost, name, weight) VALUES 
('pizza', 200, 'Pizza aglio e olio', 200),
('sushi', 123, '握り寿司', 50),
('pizza', 321, 'Pizza Regina', 300),
('bakery', 90, 'Devil’s food cake', 90),
('dessert', 100, 'Sorbet', 100);

INSERT INTO orders (client, order_time) VALUES 
(1, '2017-12-12 10:23:54'),
(2, '2017-6-05 10:23:54'),
(3, '2017-9-03 10:23:54'),
(4, '2017-2-19 10:23:54');

INSERT INTO products_in_orders (order_id, product_id, count) VALUES
(1, 1, 3),
(1, 2, 12),
(1, 5, 5),
(3, 2, 9),
(3, 5, 5),
(3, 3, 31),
(4, 3, 2),
(4, 1, 3);
 INSERT INTO clients (login, password, telephone, birthday, fullname) VALUES  ('login1', 'pass', '88005553535', '1991-01-22', 'client1'), ('login2', 'pass', '88005553535', '1992-04-03', 'client2'), ('login3', 'pass', '88005553535', '1993-07-12', 'client3'), ('login4', 'pass', '88005553535', '1994-02-08', 'client4'), ('login5', 'pass', '88005553535', '1995-12-23', 'client5'), ('login6', 'pass', '88005553535', '1996-10-11', 'client6');  
 INSERT INTO products (type, cost, name, weight) VALUES  ('pizza', 200, 'Pizza aglio e olio', 200), ('sushi', 123, '握り寿司', 50), ('pizza', 321, 'Pizza Regina', 300), ('bakery', 90, 'Devil’s food cake', 90), ('dessert', 100, 'Sorbet', 100);  
 INSERT INTO orders (client, order_time) VALUES  (1, '2017-12-12 10:23:54'), (2, '2017-6-05 10:23:54'), (3, '2017-9-03 10:23:54'), (4, '2017-2-19 10:23:54');  
 INSERT INTO products_in_orders (order_id, product_id, count) VALUES (1, 1, 3), (1, 2, 12), (1, 5, 5), (3, 2, 9), (3, 5, 5), (3, 3, 31), (4, 3, 2), (4, 1, 3); 