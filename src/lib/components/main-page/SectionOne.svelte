<script lang="ts">
    import MediaQuery from "$lib/components/generic/MediaQuery.svelte";
    import { genericDataStore } from "$lib/stores/GenericDataStore.js";

    const vidHeight = 9, vidWidth = 16;
    let mpSOHolderWidth = 0, greetSlideHeight = 0, iAmSlideHeight = 0;
</script>

<MediaQuery query="only screen and (min-aspect-ratio: {vidWidth} / {vidHeight})" let:matches>
    <svelte:fragment>
        <div class="MPSOHolder" bind:clientWidth={mpSOHolderWidth}
             style="--vid-height: {vidHeight}; --vid-width: {vidWidth}; {matches ? 'height' : 'width' }: 100%;">
            <video class="SOBgVideo" muted loop autoPlay playsinline>
                <source src="/videos/main-page-bg.webm" type="video/webm">
            </video>

            <div class="SOVideoOverlay"
                 style="--s1-cw-gap: {mpSOHolderWidth / 2}px; color: {$genericDataStore.theme.onBackground};">
                <div class="SOGreetHolder">
                    <div class="SOSlider" bind:clientHeight={greetSlideHeight}
                         style="--greet-slider-height: {greetSlideHeight};">
                        <div class="CenterRowFlex SOSliderTextWrapper GreetWrapper">
                            Bonjour Monsieur,
                        </div>
                        <div class="CenterRowFlex SOSliderTextWrapper GreetWrapper">
                            <ruby>
                                こんにちは
                                <rt>Kon'nichiwa</rt>
                                &nbsp;御仁、
                                <rt>Gojin</rt>
                            </ruby>
                        </div>
                    </div>
                </div>

                <div class="SOIAmHolder">
                    <div class="SOSlider" bind:clientHeight={iAmSlideHeight}
                         style="--i-am-slider-height: {iAmSlideHeight};">
                        <div class="CenterRowFlex SOSliderTextWrapper IAmWrapper">
                            I am
                        </div>
                        <div class="CenterRowFlex SOSliderTextWrapper IAmWrapper">
                            <ruby>
                                拙者
                                <rt>Sessha</rt>
                                &nbsp;は
                                <rt>wa</rt>
                            </ruby>
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
        /* Default values for the vars passed */
        --vid-height: 9;
        --vid-width: 16;

        aspect-ratio: var(--vid-width) / var(--vid-height);

        position: relative;
    }

    .SOBgVideo {
        width: 100%;
        height: 100%;

        position: absolute;
        inset: 0 auto auto 0;
        z-index: 2;

        pointer-events: none;
    }

    .SOVideoOverlay {
        height: 100%;
        width: calc(100vw - var(--scroll-bar-size));

        position: absolute;
        inset: 0 auto auto 0;
        z-index: 3;

        --local-value: calc((var(--vid-width) * 5%) / var(--vid-height));

        display: grid;
        grid-template:
                6.5em var(--local-value) 0.25em var(--local-value) auto /
                var(--s1-cw-gap) calc(var(--vid-width) * 0.86%) auto;

        overflow: auto;
    }

    .SOGreetHolder {
        grid-row: 2 / 3;
        grid-column: 3 / 4;

        overflow: hidden;
    }

    .SOSlider {
        height: 200%;
        width: 100%;

        animation: slide-up-down 10s ease-in infinite;
    }

    .SOSliderTextWrapper {
        height: 50%;
        width: 100%;

        overflow: hidden;
        justify-content: flex-start;

        font-family: "Noto Sans Mono", sans-serif;
        text-align: left;
    }

    .GreetWrapper {
        font-size: calc(var(--greet-slider-height) / 4 * 1px);
    }

    .SOIAmHolder {
        grid-row: 4 / 5;
        grid-column: 3 / 4;

        overflow: hidden;
    }

    .IAmWrapper {
        font-size: calc(var(--i-am-slider-height) / 4 * 1px);
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