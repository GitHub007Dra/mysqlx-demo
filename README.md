# mysqlx-demo
mysql and go sqlx demo

make -C mysqlx-demo
docker-compose up

# 登录mysql查看
docker exec -it mysqlx-demo_database_1 bash

//CREATE USER 'user_sqlx'@'%' IDENTIFIED BY 'password_sqlx';
//GRANT ALL PRIVILEGES ON *.* TO 'user_sqlx'@'%' WITH GRANT OPTION;
//FLUSH PRIVILEGES;
