<p align="center">
  <img src="PW-Info-Img.jpg" alt="Information about my personal website">
</p>

<h1 align="center">Personal Website</h1>

# Setting up the project

To install the node dependencies, run the following command in the `root` directory: -

```bash
pnpm install
```

To install the go dependencies, run the following commands starting from the `root` directory: -

```bash
cd server
go get -u
```

# Building the svelte app

To build the svelte app for production, run the following command in the `root` directory: -

```bash
pnpm run svelte:build
```

# Running the go server

Setting up the environment variables: -

<ul>
    <li><code>PORT</code>: <code>number</code></li>
    <li><code>EnvName</code>: <code>dev|prd</code></li>
</ul>

To start the go server, use the following commands starting from the `root` directory: -

```bash
cd server
go build ./main.go
./main
```

By default, the server will run on `PORT` `6971`.

## Running the dev server for svelte: -

1. Run the go server in dev env. as mentioned in the above steps.
2. Run the following command in the `root` directory: -

```bash
pnpm run svelte:dev
```

Now the dev server will serve the svelte app on `PORT` `5173` and will be able to make the WebSocket
connection with the go server(by default, at `PORT` `6971`).
