docker build -t gin-note:v1 --build-arg MYSQL_HOST=172.18.0.1 --build-arg MYSQL_USER=xxxx --build-arg MYSQL_PASSWORD=xxx --build-arg MYSQL_DBNAME=todo4 . | tee -a log.txt

# Rename Image
docker tag gin-note:v1 aswadwk/gin-note:v1 | tee -a log.txt

# Login Docker Hub
echo $PASSWORD_DOCKER_HUB | docker login -u aswadwk --password-stdin | tee -a log.txt

# Docker Push
docker push aswadwk/gin-note:v1 | tee -a log.txt