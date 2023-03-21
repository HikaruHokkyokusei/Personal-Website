<script lang="ts">
    import MediaQuery from "$lib/services/MediaQuery.svelte";
    import { personalData } from "../lib/configs/PersonalData";
</script>

<svelte:head>
    <title>{`${personalData.firstName} ${personalData.lastName}`}</title>
</svelte:head>

<div class="MainPageWrapper">
    <MediaQuery query="only screen and (min-aspect-ratio: 16 / 9)" let:matches>
        <svelte:fragment>
            <div class="MainPageBgHolder" style="{matches ? 'height' : 'width' }: 100%;">
                <video src="/videos/main-page-bg.mkv" muted loop autoPlay>
                </video>
            </div>
        </svelte:fragment>
    </MediaQuery>
</div>

<style>
    .MainPageWrapper {
        height: 100%;
        width: 100%;

        max-height: 100%;
        max-width: 100%;

        overflow-y: auto;
        scrollbar-gutter: stable both-edges;

        --scroll-track-color: rgba(255, 255, 255, 0.17);
        --scroll-thumb-image: linear-gradient(45deg, rgba(0, 175, 255, 0.5), rgba(166, 142, 255, 0.5));
        --scroll-thumb-hover-image: linear-gradient(45deg, rgba(0, 175, 255, 0.9), rgba(166, 142, 255, 0.9));
    }

    .MainPageBgHolder {
        aspect-ratio: 16 / 9;

        position: relative;
    }

    video {
        width: 100%;
        height: 100%;

        position: absolute;
        top: 0;
        left: 0;
        z-index: 2;
    }
</style>
