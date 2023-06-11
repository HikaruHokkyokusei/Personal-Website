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

# Running Servers

## Go Server

To start the go server, use the following commands starting from the `root` directory: -

```bash
cd server
go build ./main.go
./main
```

By default, the server will run on `PORT` `42069`. Look at the environment
variables [below](#environment-variables-setup) for customizations.

## Vite Dev Server: -

1. Run the go server in dev env. as mentioned in the above steps.
2. Run the following command in the `root` directory: -

```bash
pnpm run svelte:dev
```

Now the dev server will serve the svelte app on `PORT` `5173` and will be able to make the WebSocket
connection with the go server(by default, with host `http://localhost:42069`). Look at the environment
variables [below](#environment-variables-setup) for customizations.

<style>
    table, tr, th, td {
        border: 1px solid gray;
        border-collapse: collapse;
        text-align: center;
    }
</style>

# Environment Variables Setup

<table>
    <thead>
        <tr>
            <th>Key</th>
            <th>Allowed Values</th>
            <th>Example Value</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td colspan="3" style="text-align: center; background-color: rgba(91,129,129,0.75)">GO Server</td>
        </tr>
        <tr>
            <td><code>EnvName</code></td>
            <td><code>dev|prd</code></td>
            <td><code>dev</code></td>
        </tr>
        <tr>
            <td><code>EnableRouteLogs</code></td>
            <td><code>boolean</code></td>
            <td><code>true</code></td>
        </tr>
        <tr>
            <td><code>PORT</code></td>
            <td><code>number</code></td>
            <td><code>42069</code></td>
        </tr>
        <tr>
            <td><code>AllowedOrigins</code></td>
            <td><code>string</code></td>
            <td><code>http://localhost:5173, http://127.0.0.1:5173, ws://localhost:5173, ws://127.0.0.1:5173</code></td>
        </tr>
        <tr>
            <td><code>StripePublicKey</code></td>
            <td><code>string</code></td>
            <td><code>pk_test_************************************</code></td>
        </tr>
        <tr>
            <td><code>StripePrivateKey</code></td>
            <td><code>string</code></td>
            <td><code>sk_test_************************************</code></td>
        </tr>
        <tr>
            <td colspan="3" style="text-align: center; background-color: rgba(145,93,74,0.75)">Vite Dev Server &nbsp; or &nbsp; Vite App Build</td>
        </tr>
        <tr>
            <td><code>PUBLIC_SERVER_LOCATION_ORIGIN</code></td>
            <td><code>string</code></td>
            <td><code>http://localhost:42069</code></td>
        </tr>
    </tbody>
</table>
