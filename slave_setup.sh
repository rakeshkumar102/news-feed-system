#!/bin/bash

until docker exec mysql-master sh -c 'export MYSQL_PWD=masterpassword; mysql -u root -e";"'
do
	echo "Waiting for mysql-master database connection..."
	sleep 4
done

MS_STATUS=`docker exec mysql-master sh -c 'export MYSQL_PWD=masterpassword; mysql -u root -e "show master status"'`
CURRENT_LOG=`echo $MS_STATUS | awk '{print $6}'`
CURRENT_POS=`echo $MS_STATUS | awk '{print $7}'`

echo $CURRENT_LOG
echo $CURRENT_POS

stop_slave="stop slave;"
reset_slave="reset slave;"
start_slave_stmt="CHANGE MASTER TO MASTER_HOST='mysql_master',MASTER_USER='mydb_slave_user',MASTER_PASSWORD='mydb_slave_pwd',MASTER_LOG_FILE='$CURRENT_LOG',MASTER_LOG_POS=$CURRENT_POS;"
start_slave="start slave;"
show_status="show slave status\G"

cmd="export MYSQL_PWD=masterpassword; mysql -u root -e'"

stop_slave_cmd="export MYSQL_PWD=masterpassword; mysql -u root -e '$stop_slave'"

reset_slave_cmd="export MYSQL_PWD=masterpassword; mysql -u root -e '$reset_slave'"

start_slave_stmt_cmd='export MYSQL_PWD=masterpassword; mysql -u root -e "'
start_slave_stmt_cmd+="$start_slave_stmt"
start_slave_stmt_cmd+='"'

start_slave_cmd="export MYSQL_PWD=masterpassword; mysql -u root -e '$start_slave'"

show_status_cmd="export MYSQL_PWD=masterpassword; mysql -u root -e '$show_status'"

docker exec mysql-slave-1 sh -c "$stop_slave_cmd"
docker exec mysql-slave-2 sh -c "$stop_slave_cmd"
docker exec mysql-slave-3 sh -c "$stop_slave_cmd"

docker exec mysql-slave-1 sh -c "$reset_slave_cmd"
docker exec mysql-slave-2 sh -c "$reset_slave_cmd"
docker exec mysql-slave-3 sh -c "$reset_slave_cmd"

docker exec mysql-slave-1 sh -c "$start_slave_stmt_cmd"
docker exec mysql-slave-2 sh -c "$start_slave_stmt_cmd"
docker exec mysql-slave-3 sh -c "$start_slave_stmt_cmd"

docker exec mysql-slave-1 sh -c "$start_slave_cmd"
docker exec mysql-slave-2 sh -c "$start_slave_cmd"
docker exec mysql-slave-3 sh -c "$start_slave_cmd"

echo "MYSQL SLAVE 1"
docker exec mysql-slave-1 sh -c "$show_status_cmd"
echo "MYSQL SLAVE 2"
docker exec mysql-slave-2 sh -c "$show_status_cmd"
echo "MYSQL SLAVE 3"
docker exec mysql-slave-3 sh -c "$show_status_cmd"

