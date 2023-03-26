<script lang="ts">
    import "../global.css";
    import { onMount } from "svelte";
    import { WebSocketService } from "../lib/services/WebSocketService";
    import { UtilService } from "../lib/services/UtilService";
    import { genericDataStore } from "../lib/stores/GenericDataStore";
    import MainLoader from "$lib/components/generic/MainLoader.svelte";
    import Hamburger from "$lib/components/generic/Hamburger.svelte";

    let isLoading = true;

    onMount(async () => {
        const delayer = UtilService.delay(1250);
        const websiteOrigin = document.location.origin.replace(/localhost:[0-9]{4,}$/, "localhost:6969");

        const res = await fetch(`${websiteOrigin}/healthCheck`);
        let showValue = "❌", showError = null;
        try {
            if (new TextDecoder().decode((await res.body.getReader().read()).value).toLowerCase() === "ok") {
                showValue = "✔️";
            }
        } catch (err) {
            showError = err;
        }
        console.log(`Health Check: ${showValue}`);
        if (showError) {
            console.log(showError);
        }

        WebSocketService.connect(websiteOrigin.replace(/^http/, "ws"), async () => {
            await delayer;
            isLoading = false;
        });
    });
</script>

<div class="CenterRowFlex MainHolder" style="background: {$genericDataStore.theme.background};">
    {#if isLoading}
        <MainLoader></MainLoader>
    {:else}
        <Hamburger></Hamburger>
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