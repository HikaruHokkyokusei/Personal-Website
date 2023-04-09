<script lang="ts">
    import { onMount } from "svelte";
    import { personalData } from "../lib/configs/PersonalData";
    import { genericDataStore } from "../lib/stores/GenericDataStore";
    import MainLoader from "../lib/components/generic/MainLoader.svelte";
    import Section1v1 from "../lib/components/main-page/Section1v1.svelte";
    import Section1v2 from "../lib/components/main-page/Section1v2.svelte";
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
        <Section1v1 on:videoLoading={switchToLoader} on:videoLoaded={switchToContent}></Section1v1>
    {:else}
        <Section1v2></Section1v2>
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
