<script lang="ts">
    import { onMount } from "svelte";

    onMount(() => {
        fetch("http://localhost:6969/healthCheck").then(async (res) => {
            console.log(new TextDecoder().decode((await res.body.getReader().read()).value));
        }).catch((err) => {
            console.log(err);
        });

        const socket = new WebSocket("ws://localhost:6969/ws/");
        socket.onerror = (err) => {
            console.log("Socket Error");
            console.log(err);
        }
        socket.onopen = () => {
            console.log("Opened...");
            socket.send("message");
        };
    });
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
