#!/bin/bash

until docker exec mysql-master sh -c 'export MYSQL_PWD=masterpassword; mysql -u root -e";"'
do
	echo "Waiting for mysql-master database connection..."
	sleep 4
done

priv_stmt='CREATE USER "mydb_slave_user"@"%" IDENTIFIED BY "mysql_slave_pwd"; GRANT REPLICATION SLAVE ON *.* TO "mydb_slave_user"@"%"; FLUSH PRIVILAGES;'
docker exec mysql-master sh -c "export MYSQL_PWD=masterdatabase; mysql -u root -e '$priv_stmt'"