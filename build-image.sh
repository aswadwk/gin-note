docker build -t gin-note --build-arg MYSQL_HOST=172.18.0.1 --build-arg MYSQL_USER=xxxx --build-arg MYSQL_PASSWORD=xxx --build-arg MYSQL_DBNAME=todo4 .
docker run -e MYSQL_HOST=192.168.0.11 -e MYSQL_USER=root -e MYSQL_PASSWORD=root -e MYSQL_DBNAME=go_note_db -p 8090:3030 gin-note
