FROM golang:alpine AS golang-build

WORKDIR /golang

COPY ./api/go.mod ./api/go.sum ./

RUN go mod download

COPY ./api/ ./

RUN go build -o main .

RUN chmod +x ./main


FROM node:alpine AS svelte

WORKDIR /svelte

COPY ./web/package*.json ./

RUN npm install

COPY ./web/ .

RUN npm run build

FROM busybox:latest AS runtime
COPY --from=golang-build ./golang/main ./main

# COPY ./api/clean-data ./clean-data

COPY --from=svelte ./svelte/build .
CMD ./main & exec busybox httpd -f -v -p 5173