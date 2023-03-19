<script lang="ts">
    import { onMount } from "svelte";
    import { WebSocketService } from "../lib/services/WebSocketService";
    import { UtilService } from "../lib/services/UtilService";

    let isLoading = true;

    onMount(async () => {
        const hostUrl = document.location.host.replace("5173", "6969");

        const res = await fetch(`http://${hostUrl}/healthCheck`);
        const response = new TextDecoder().decode((await res.body.getReader().read()).value);
        console.log(`Health Check: ${response}`);

        WebSocketService.connect(hostUrl);
        await UtilService.delay(500);

        isLoading = false;
    });
</script>

{#if isLoading}
    Loading...
{:else}
    <slot></slot>
{/if}