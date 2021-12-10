#! /bin/bash

DOCKER_REPO=asia.gcr.io/lcwp-ray/private

DOCKER_VERSION=1.0.0

docker build -f ./main.dockerfile -t "$DOCKER_REPO"-stocker:"$DOCKER_VERSION" .
docker push -a "$DOCKER_REPO"-stocker

echo ""
echo "@ Build Done!"