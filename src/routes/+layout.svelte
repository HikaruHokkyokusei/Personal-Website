<script lang="ts">
    import "../global.css";
    import { onMount } from "svelte";
    import { WebSocketService } from "../lib/services/WebSocketService";
    import { UtilService } from "../lib/services/UtilService";
    import MainLoader from "../lib/components/MainLoader.svelte";

    let isLoading = true;

    onMount(async () => {
        const delayer = UtilService.delay(1250);
        const websiteOrigin = document.location.origin.replace(/localhost:[0-9]{4,}$/, "localhost:6969");

        const res = await fetch(`${websiteOrigin}/healthCheck`);
        const response = new TextDecoder().decode((await res.body.getReader().read()).value);
        console.log(`Health Check: ${response}`);

        WebSocketService.connect(websiteOrigin.replace(/^http/, "ws"));

        await delayer;
        isLoading = false;
    });
</script>

<div class="CenterRowFlex MainHolder">
    {#if isLoading}
        <MainLoader></MainLoader>
    {:else}
        <slot></slot>
    {/if}
</div>

<style>
    .MainHolder {
        height: 100%;
        width: 100%;

        position: relative;
        z-index: 0;
    }
</style>