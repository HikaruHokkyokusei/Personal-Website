<script lang="ts">
    import MediaQuery from "$lib/components/generic/MediaQuery.svelte";
    import { personalData } from "$lib/configs/PersonalData.js";
    import { genericDataStore } from "$lib/stores/GenericDataStore.js";
    import Animated5ColorBorderText from "$lib/components/generic/Animated5ColorBorderText.svelte";

    const vidHeight = 9, vidWidth = 16;
    let mpSOHolderWidth = 0, greetSlideHeight = 0, iAmSlideHeight = 0, endHolderHeight = 0;
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

                <Animated5ColorBorderText text="{personalData.firstName}"
                                          holderStyle="grid-row: 6 / 7; grid-column: 3 / 4;">
                </Animated5ColorBorderText>
                <Animated5ColorBorderText text="{personalData.lastName}"
                                          holderStyle="grid-row: 8 / 9; grid-column: 3 / 4;">
                </Animated5ColorBorderText>

                <div class="SOEndHolder" bind:clientHeight={endHolderHeight}>
                    <div class="CenterRowFlex SOSliderTextWrapper SOEndTextWrapper"
                         style="--end-holder-height: {endHolderHeight}">
                        <ruby>
                            で
                            <rt>de</rt>
                            &nbsp;ござる。
                            <rt>Gozaru</rt>
                        </ruby>
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
        grid-template: calc(2.5 * var(--local-value))  var(--local-value)
                0.27em var(--local-value)
                1em calc(1.3 * var(--local-value))
                0.25em calc(1.3 * var(--local-value))
                1.4em var(--local-value)
                auto / var(--s1-cw-gap) calc(var(--vid-width) * 0.86%) auto;

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

    .SOEndHolder {
        grid-row: 10 / 11;
        grid-column: 3 / 4;

        overflow: hidden;
    }

    .SOEndTextWrapper {
        height: 100%;

        font-size: calc(var(--end-holder-height) / 2 * 1px);

        animation: pulse-opacity 10s ease-in infinite;
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

    @keyframes pulse-opacity {
        0% {
            opacity: 0;
        }

        45% {
            opacity: 0;
        }

        50% {
            opacity: 1;
        }

        95% {
            opacity: 1;
        }

        100% {
            opacity: 0;
        }
    }
</style>