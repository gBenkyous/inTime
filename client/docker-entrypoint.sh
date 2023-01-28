#!/bin/sh

WORK_DIR="/usr/src/app"
APP="react-app"
APP_DIR="${WORK_DIR}/${APP}"

if [ -z "$(ls ${APP_DIR})" ]; then
    echo "[START] create react app..."
    cd ${WORK_DIR}
    create-react-app ${APP} --template typescript
    rm ${APP}/package-lock.json
    echo "[END] create react app"
fi

cd ${APP_DIR}
yarn install
yarn start

exec "$@"