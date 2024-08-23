ARG PORT


FROM node:22-alpine AS node-base
WORKDIR /personal-website
RUN apk upgrade


FROM golang:1.23-alpine AS go-base
WORKDIR /personal-website


FROM go-base as install-go-deps-and-build
COPY ./server/. .
RUN go get -u
RUN CGO_ENABLED=0 go build -ldflags="-s -w" ./main.go


FROM node-base AS get-go-exec
WORKDIR /personal-website-temp
COPY --from=install-go-deps-and-build /personal-website/main .


FROM get-go-exec as install-node-deps
WORKDIR /personal-website
COPY ./package.json .
RUN npm install


FROM install-node-deps as copy-code-and-files
COPY . .
RUN mv /personal-website-temp/main ./server/main
RUN rmdir /personal-website-temp


FROM copy-code-and-files as build-node-app
ARG PUBLIC_SERVER_LOCATION_ORIGIN
ENV NODE_ENV=production
ENV PUBLIC_SERVER_LOCATION_ORIGIN=${PUBLIC_SERVER_LOCATION_ORIGIN}
RUN npm run svelte:build
RUN npm prune --omit=dev


FROM build-node-app as final
WORKDIR /personal-website/server
EXPOSE ${PORT}
CMD ["./main"]
