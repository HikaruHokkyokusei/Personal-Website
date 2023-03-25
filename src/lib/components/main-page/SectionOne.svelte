<script lang="ts">
    import MediaQuery from "$lib/components/generic/MediaQuery.svelte";
    import { genericDataStore } from "$lib/stores/GenericDataStore.js";

    let s1CHWidth = 0, s1CHHeight = 0, s1CWWidth;

    $: {
        s1CWWidth = Math.floor(s1CHWidth - (s1CHHeight * 8 / 9));
    }
</script>

<MediaQuery query="only screen and (min-aspect-ratio: 16 / 9)" let:matches>
    <svelte:fragment>
        <div class="MPSOHolder" style="{matches ? 'height' : 'width' }: 100%;">
            <video class="SOBgVideo" muted loop autoPlay playsinline>
                <source src="/videos/main-page-bg.webm" type="video/webm">
            </video>

            <div class="SOContentHolder" style="color: {$genericDataStore.theme.onBackground}"
                 bind:clientWidth={s1CHWidth} bind:clientHeight={s1CHHeight}>
                <div class="CenterColumnFlex SOContentWrapper" style="--s1-cw-width: {s1CWWidth}px;">
                    <div class="SOGreetHolder">
                        <div class="SOGreetSlider">
                            <div class="CenterRowFlex SOGreetWrapper">
                                Bonjour Monsieur,
                            </div>
                            <div class="CenterRowFlex SOGreetWrapper">
                                <ruby>
                                    こんにちは
                                    <rt>Kon'nichiwa</rt>
                                    &nbsp;御仁、
                                    <rt>Gojin</rt>
                                </ruby>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </svelte:fragment>
</MediaQuery>

<style>
    /* MP => Main page
     * SO => Section One
    */

    .MPSOHolder {
        aspect-ratio: 16 / 9;

        position: relative;
    }

    .SOBgVideo {
        width: 100%;
        height: 100%;

        position: absolute;
        top: 0;
        left: 0;
        z-index: 2;

        pointer-events: none;
    }

    .SOContentHolder {
        height: 100%;
        width: calc(100vw - var(--scroll-bar-size));

        position: absolute;
        top: 0;
        left: 0;
        z-index: 3;
    }

    .SOContentWrapper {
        width: var(--s1-cw-width);
        height: 100%;

        position: absolute;
        top: 0;
        right: 0;
    }

    .SOGreetHolder {
        height: 90px;
        width: 600px;

        overflow: hidden;
    }

    .SOGreetSlider {
        height: 200%;
        width: 100%;

        animation: slide-up-down 10s ease-in infinite;
    }

    .SOGreetWrapper {
        height: 50%;
        width: 100%;

        font: 3rem "Noto Sans Mono", sans-serif;
        text-align: center;
    }

    @keyframes slide-up-down {
        0% {
            transform: translateY(0%);
        }

        45% {
            transform: translateY(0%);
        }

        50% {
            transform: translateY(-50%);
        }

        95% {
            transform: translateY(-50%);
        }

        100% {
            transform: translateY(0%);
        }
    }
</style>