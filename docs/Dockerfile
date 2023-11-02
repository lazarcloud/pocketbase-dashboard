FROM node:lts-alpine AS build

COPY ./docs/package*.json ./

RUN npm i

COPY ./docs/ .
RUN npm run build

FROM busybox:latest AS runtime
COPY --from=build /dist .
CMD ["busybox", "httpd", "-f", "-v", "-p", "80"]