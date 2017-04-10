#!/bin/bash

# Only build docker container for master branch
if [ "$TRAVIS_PULL_REQUEST" == "false" ]; then echo 'upload to Docker Hub'; else echo 'skip docker upload for PR'; fi

export REPO=jaegertracing/all-in-one
export BRANCH=$(if [ "$TRAVIS_PULL_REQUEST" == "false" ]; then echo $TRAVIS_BRANCH; else echo $TRAVIS_PULL_REQUEST_BRANCH; fi)
export TAG=`if [ "$BRANCH" == "master" ]; then echo "latest"; else echo "${BRANCH///}"; fi`
echo "TRAVIS_BRANCH=$TRAVIS_BRANCH, REPO=$REPO, PR=$PR, BRANCH=$BRANCH, TAG=$TAG"

nvm use 4
make build-all-in-one-linux || { echo 'failed build-all-in-one-linux' ; exit 0; }
docker login -u $DOCKER_USER -p $DOCKER_PASS
docker build -f cmd/standalone/Dockerfile -t $REPO:$COMMIT .
docker tag $REPO:$COMMIT $REPO:$TAG
docker tag $REPO:$COMMIT $REPO:travis-$TRAVIS_BUILD_NUMBER
docker push $REPO
