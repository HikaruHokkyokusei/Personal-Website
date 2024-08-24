<script lang="ts">
    import { PUBLIC_SERVER_LOCATION_ORIGIN } from "$env/static/public";
    import "../global.css";
    import { onMount } from "svelte";
    import { WebSocketService } from "$lib/services/WebSocketService";
    import { genericDataStore } from "$lib/stores/GenericDataStore";
    import MainLoader from "$lib/components/generic/MainLoader.svelte";

    let isLoading = true;

    onMount(async () => {
        const res = await fetch(`${PUBLIC_SERVER_LOCATION_ORIGIN}/healthCheck`);
        let showValue = "❌", showError = null;
        if (res.body != null) {
            try {
                if (new TextDecoder().decode((await res.body.getReader().read()).value).toLowerCase() === "ok") {
                    showValue = "✔️";
                }
            } catch (err) {
                showError = err;
            }
        } else {
            showError = "Error: Health Check has no body!"
        }
        console.log(`Health Check: ${showValue}`);
        if (showError) {
            console.log(showError);
        } else {
            WebSocketService.connect(PUBLIC_SERVER_LOCATION_ORIGIN.replace(/^http/, "ws"), () => {
                isLoading = false;
            });
        }
    });
</script>

<div class="CenterRowFlex MainHolder" style="background: {$genericDataStore.theme.background};">
    {#if isLoading}
        <MainLoader ringGapEm={1.5}></MainLoader>
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
