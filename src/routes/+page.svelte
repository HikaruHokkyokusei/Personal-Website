<script lang="ts">
    import { onMount } from "svelte";
    import { personalData } from "../lib/configs/PersonalData";
    import { genericDataStore } from "../lib/stores/GenericDataStore";
    import MainLoader from "../lib/components/generic/MainLoader.svelte";
    import SectionOneV1 from "../lib/components/main-page/SectionOneV1.svelte";
    import SectionOneV2 from "../lib/components/main-page/SectionOneV2.svelte";
    import SectionTwo from "../lib/components/main-page/SectionTwo.svelte";

    let hideContent = true;

    const switchToLoader = () => {
        hideContent = true;
        $genericDataStore.showHamburger = false;
    };
    const switchToContent = () => {
        hideContent = false;
        $genericDataStore.showHamburger = true;
    };

    onMount(() => {
        if ($genericDataStore.mainPageSectionOneVersion !== 1) {
            switchToContent();
        }
    });
</script>

<svelte:head>
    <title>{`${personalData.firstName} ${personalData.lastName}`}</title>
</svelte:head>

{#if hideContent}
    <MainLoader ringGapEm="1.5"></MainLoader>
{/if}
<div class="MainPageWrapper" class:MainPageContentHidden="{hideContent}">
    {#if $genericDataStore.mainPageSectionOneVersion === 1}
        <SectionOneV1 on:videoLoading={switchToLoader} on:videoLoaded={switchToContent}></SectionOneV1>
    {:else}
        <SectionOneV2></SectionOneV2>
    {/if}

    <SectionTwo></SectionTwo>
</div>

<style>
    .MainPageWrapper {
        height: 100%;
        width: 100%;

        max-height: 100%;
        max-width: 100%;

        overflow: hidden auto;
        scrollbar-gutter: stable;

        --scroll-track-color: rgba(255, 255, 255, 0.17);
        --scroll-thumb-image: linear-gradient(45deg, rgba(0, 175, 255, 0.5), rgba(166, 142, 255, 0.5));
        --scroll-thumb-hover-image: linear-gradient(45deg, rgba(0, 175, 255, 0.9), rgba(166, 142, 255, 0.9));
    }

    .MainPageContentHidden {
        display: none;
    }
</style>
