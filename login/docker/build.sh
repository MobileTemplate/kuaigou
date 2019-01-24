#!/usr/bin/env bash

set -e

echo "docker build start ..."

# 设置 appname 分支名称，git 提交记录等信息
PROJECT_NAME=$(echo $(dirname $(pwd)) | rev | awk -F \/ '{print $1}' | rev)
BRNAME=$(git symbolic-ref --short HEAD)
BRNAME=${BRNAME/\//-}
VERSION=$(git log --pretty=format:"%h-%cd" --date=format:"%m-%d" | head -1  | awk '{print $1}')
IMAGE_NAME=$PROJECT_NAME
if [ "$BRNAME"x = "master"x ]; then
	IMAGE_NAME=$PROJECT_NAME
else
	IMAGE_NAME=$PROJECT_NAME-$BRNAME
fi
APP_NAME=$VERSION-$IMAGE_NAME

# 编译 go 程序
echo "go build $APP_NAME ..."
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ..

PACKAGE_NAME=qianuuu.com/$PROJECT_NAME
LC_PATH=$PWD/../..
PJ_PATH=/go/src/qianuuu.com/kuaigou
docker run --rm -e CGO_ENABLED=0 \
	   -v $LC_PATH:$PJ_PATH \
	   -w $PJ_PATH/$PROJECT_NAME/docker qianuuu.cn/golang go build -o app ..

# 创建 Dockerfile 文件
rm -f ./Dockerfile
touch ./Dockerfile
echo 'FROM qianuuu.cn/kuaigou' > Dockerfile
echo ADD app /$APP_NAME >> Dockerfile
echo CMD [\"/${APP_NAME}\"] >> Dockerfile

# 构建 docker image
REGISTRY=139.224.239.128:5000
DOCKER_IMAGE=$REGISTRY/$IMAGE_NAME
if [ -n "$1" ]; then
	DOCKER_IMAGE=$DOCKER_IMAGE:$1
fi

echo 'docker build $DOCKER_IMAGE'
docker build --force-rm=true -t $DOCKER_IMAGE .

rm ./app
rm ./Dockerfile

echo ''
echo "docker build $DOCKER_IMAGE succeed"

