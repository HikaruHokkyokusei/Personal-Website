FROM node:18-alpine AS node-base
WORKDIR /personal-website
RUN apk upgrade
RUN npm install -g --upgrade pnpm


FROM golang:1.20.2-alpine3.17 AS go-base
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
RUN pnpm install


FROM install-node-deps as copy-code-and-files
COPY . .
RUN mv /personal-website-temp/main ./server/main
RUN rmdir /personal-website-temp


FROM copy-code-and-files as build-node-app
ENV NODE_ENV=production
RUN pnpm run svelte:build
RUN pnpm prune --prod


FROM build-node-app as final
WORKDIR /personal-website/server
ENV PORT=6969
ENV EnvName=prd
EXPOSE 6969
CMD ["./main"]
