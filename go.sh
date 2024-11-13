#!/bin/bash

# Variables
PROJECT_ROOT="."

IMAGE_NAME="tig"
CLI_NAME="tig"
CONTAINER_NAME=$IMAGE_NAME
DOCKERFILE_PATH=$PROJECT_ROOT

HOST_BIND_MOUNT=$PROJECT_ROOT
CONTAINER_BIND_MOUNT="/app"

remove_container() {
  docker rm -f $CONTAINER_NAME
}

remove_image() {
  docker rmi -f $IMAGE_NAME
}

build_container() {
  remove_container
  remove_image

  docker build -t $IMAGE_NAME $DOCKERFILE_PATH
  docker run -d --name $CONTAINER_NAME -v $HOST_BIND_MOUNT:$CONTAINER_BIND_MOUNT $IMAGE_NAME

  echo "Container $CONTAINER_NAME is built and running!"
}

start_container() {
  echo "Running Docker container: $CONTAINER_NAME..."
  docker start $CONTAINER_NAME
}

stop_container() {
  echo "Stopping Docker container: $CONTAINER_NAME..."
  docker stop $CONTAINER_NAME
}

remove_workspace() {
  echo "Removing Docker container: $CONTAINER_NAME..."
  remove_container
  remove_image
}

compile_project() {
  GOOS=linux GOARCH=amd64 go build -buildvcs=false -o tig-$VERSION
}

package_project() {
  compile_project
  tar -czvf tig-$VERSION.tar.gz ./$CLI_NAME
  rm ./$CLI_NAME
}

if [ "$1" == "build" ]; then
  build_container
elif [ "$1" == "start" ]; then
  start_container
elif [ "$1" == "remove" ]; then
  remove_workspace
elif [ "$1" == "stop" ]; then
  stop_container
elif [ "$1" == "compile" ]; then
  compile_project
elif [ "$1" == "package" ]; then
  package_project
else
  echo "Usage [command]"
fi
