<script lang="ts">
    import { genericDataStore } from "../../stores/GenericDataStore.js";

    export let hamburgerSizeEm = 3;

    let isHovering = false;
    const setHover = (value: boolean) => {
        isHovering = value;
    };
</script>

<div class="CenterRowFlex HamburgerHolder" class:HamburgerHover={isHovering}
     on:mouseenter={() => setHover(true)} on:mouseleave={() => {setHover(false)}}
     style="--hamburger-size: {hamburgerSizeEm}; background: {$genericDataStore.theme.primary};">
    {#if isHovering}
        <!--        TODO: Fill options here...-->
    {:else}
        <i class="fa-solid fa-bars" style="font-size: 20px; color: {$genericDataStore.theme.onPrimary};"></i>
    {/if}
</div>

<style>
    .HamburgerHolder {
        height: calc(var(--hamburger-size) * 1em);
        width: calc(var(--hamburger-size) * 1em);

        position: absolute;
        inset: 30px 30px auto auto;
        z-index: 100;

        border-radius: calc(var(--hamburger-size) * 0.5em);

        transition: 800ms width cubic-bezier(0.25, 0.8, 0.6, 1); /* https://cubic-bezier.com/ */
    }

    .HamburgerHover {
        width: 55%;

        padding: 0 calc(var(--hamburger-size) * 0.5em);
    }
</style>