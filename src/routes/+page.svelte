<script lang="ts">
    import MediaQuery from "$lib/services/MediaQuery.svelte";
    import { personalData } from "../lib/configs/PersonalData";

    let s1ContentHolderWidth = 0, s1ContentHolderHeight = 0, s1ContentWrapperWidth;

    $: {
        s1ContentWrapperWidth = Math.floor(s1ContentHolderWidth - (s1ContentHolderHeight * 8 / 9));
    }
</script>

<svelte:head>
    <title>{`${personalData.firstName} ${personalData.lastName}`}</title>
</svelte:head>

<div class="MainPageWrapper">
    <MediaQuery query="only screen and (min-aspect-ratio: 16 / 9)" let:matches>
        <svelte:fragment>
            <div class="MainPageSectionOneHolder" style="{matches ? 'height' : 'width' }: 100%;">
                <video muted loop autoPlay>
                    <source src="/videos/main-page-bg.webm" type="video/webm">
                </video>
                <div class="SectionOneContentHolder"
                     bind:clientWidth={s1ContentHolderWidth} bind:clientHeight={s1ContentHolderHeight}>
                    <div class="CenterColumnFlex SectionOneContentWrapper"
                         style="--s1-content-wrapper-width: {s1ContentWrapperWidth}px;">
                        <h1>Kon'nichiwa Gojin</h1>
                    </div>
                </div>
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
        scrollbar-gutter: stable;

        --scroll-track-color: rgba(255, 255, 255, 0.17);
        --scroll-thumb-image: linear-gradient(45deg, rgba(0, 175, 255, 0.5), rgba(166, 142, 255, 0.5));
        --scroll-thumb-hover-image: linear-gradient(45deg, rgba(0, 175, 255, 0.9), rgba(166, 142, 255, 0.9));
    }

    .MainPageSectionOneHolder {
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

        pointer-events: none;
    }

    .SectionOneContentHolder {
        height: 100%;
        width: calc(100vw - var(--scroll-bar-size));

        position: absolute;
        top: 0;
        left: 0;
        z-index: 3;
    }

    .SectionOneContentWrapper {
        width: var(--s1-content-wrapper-width);
        height: 100%;

        position: absolute;
        top: 0;
        right: 0;
    }
</style>
