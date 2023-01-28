FROM node:18.13.0-alpine

COPY ./docker-entrypoint.sh /
RUN npm cache clean --force && \
    yarn global add create-react-app && \
    yarn cache clean && \
    chmod +x /docker-entrypoint.sh

WORKDIR /usr/src/app/react-app
ENTRYPOINT ["/docker-entrypoint.sh"]