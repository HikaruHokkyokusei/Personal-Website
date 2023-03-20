FROM alpine AS base
WORKDIR /personal-website
RUN apk upgrade


FROM base as install-nodejs
RUN apk add nodejs npm
RUN npm install -g --upgrade npm
RUN npm install -g --upgrade pnpm


FROM install-nodejs as install-golang
COPY --from=golang:1.20.2-alpine3.17 /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"
RUN go version


FROM install-golang as install-node-deps
COPY package.json .
RUN pnpm install


FROM install-node-deps as install-go-deps
WORKDIR /personal-website/server
COPY ./server/. .
RUN go get -u
WORKDIR /personal-website


FROM install-go-deps as copy-code
COPY . .


FROM copy-code as build-app
RUN pnpm run svelte:build
WORKDIR /personal-website/server
RUN go build ./main.go


FROM build-app as final
ENV PORT=6969
ENV EnvName=prd
EXPOSE 6969
CMD ["./main"]
