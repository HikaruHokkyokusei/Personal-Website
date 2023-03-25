<script lang="ts">
    import { personalData } from "../lib/configs/PersonalData";
    import SectionOne from "$lib/components/main-page/SectionOne.svelte";
    import SectionTwo from "$lib/components/main-page/SectionTwo.svelte";
</script>

<svelte:head>
    <title>{`${personalData.firstName} ${personalData.lastName}`}</title>
</svelte:head>

<div class="MainPageWrapper">
    <SectionOne></SectionOne>
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
</style>
