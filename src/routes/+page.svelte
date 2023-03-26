<script lang="ts">
    import { personalData } from "../lib/configs/PersonalData";
    import { genericDataStore } from "../lib/stores/GenericDataStore";
    import SectionOne from "$lib/components/main-page/SectionOne.svelte";
    import SectionTwo from "$lib/components/main-page/SectionTwo.svelte";
    import MainLoader from "$lib/components/generic/MainLoader.svelte";

    let hideContent = true;

    const videoLoading = () => {
        hideContent = true;
        $genericDataStore.showHamburger = false;
    };
    const videoLoaded = () => {
        hideContent = false;
        $genericDataStore.showHamburger = true;
    };
</script>

<svelte:head>
    <title>{`${personalData.firstName} ${personalData.lastName}`}</title>
</svelte:head>

{#if hideContent}
    <MainLoader ringGapEm="1.5"></MainLoader>
{/if}
<div class="MainPageWrapper" class:DisplayHidden={hideContent}>
    <SectionOne on:videoLoading={videoLoading} on:videoLoaded={videoLoaded}></SectionOne>
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

    .DisplayHidden {
        display: none;
    }
</style>
