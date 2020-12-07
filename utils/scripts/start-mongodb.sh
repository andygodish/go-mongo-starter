set -e

SERVER="my_database_server"
PORT="27017"
USERNAME="app"
PASSWORD="mysecretpassword"
PWD=$(pwd)

echo "echo stop & remove old docker [$SERVER] and starting new fresh instance of [$SERVER] --- $PWD"
(docker kill $SERVER || :) && \
  (docker rm $SERVER || :) && \
  docker run --name $SERVER -p $PORT:$PORT \
  -e MONGO_INITDB_ROOT_USERNAME=$USERNAME \
  -e MONGO_INITDB_ROOT_PASSWORD=$PASSWORD \
  -d mongo:latest

# Add the below line after line 14 above (-e MONGO_INITDB_TOOT_PASSWORD) - this will look for a volume mount locallaly, you will need to add a data-volumes/demo dir in your project root dir
# -v $PWD/data-volumes/demo:/data/db \

# Acutal mongo image being used by P1 in our stagin environment
# -d registry.il2.dsop.io/spacecamp/battledrill/battledrill-manifests/mongo:latest

# wait for pg to start
echo "sleep wait for server [$SERVER] to start"
sleep 10

# create the db 
echo "FINISHED"