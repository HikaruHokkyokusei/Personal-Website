<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import MediaQuery from "../../components/generic/MediaQuery.svelte";
    import Animated5ColorBorderText from "../../components/generic/Animated5ColorBorderText.svelte";
    import { personalData } from "../../configs/PersonalData.js";
    import { genericDataStore } from "../../stores/GenericDataStore.js";

    const dispatch = createEventDispatcher();
    let canPlayVideo = false;

    const vidHeight = 9, vidWidth = 16;
    let mpSOHolderWidth = 0, greetSliderHeight = 0, iAmSliderHeight = 0, endHolderHeight = 0;

    const loadStarted = () => {
        if (!canPlayVideo) {
            dispatch('videoLoading');
        }
    };

    const loadCompleted = () => {
        canPlayVideo = true;
        dispatch('videoLoaded');
    };
</script>

<MediaQuery query="only screen and (min-aspect-ratio: {vidWidth} / {vidHeight})" let:matches>
    <svelte:fragment>
        <div class="MPS1V1Holder" bind:clientWidth={mpSOHolderWidth}
             style="--vid-height: {vidHeight}; --vid-width: {vidWidth}; {matches ? 'height' : 'width' }: 100%;">
            <video class="S1V1BgVideo" src="/videos/main-page-bg.webm" muted loop autoplay playsinline
                   on:loadstart={loadStarted} on:loadeddata={loadCompleted}>
            </video>
            <div class="S1V1VideoOverlay"
                 style="--s1-cw-gap: {mpSOHolderWidth / 2}px; color: {$genericDataStore.theme.onBackground};">
                <div class="S1V1GreetHolder">
                    <div class="S1V1Slider" bind:clientHeight={greetSliderHeight}
                         style="--greet-slider-height: {greetSliderHeight};">
                        <div class="CenterRowFlex S1V1SliderTextWrapper S1V1GreetWrapper">
                            Hello Monsieur,
                        </div>
                        <div class="CenterRowFlex S1V1SliderTextWrapper S1V1GreetWrapper">
                            <ruby>
                                こんにちは
                                <rt>Kon'nichiwa</rt>
                                &nbsp;御仁
                                <rt>Gojin</rt>
                                、
                            </ruby>
                        </div>
                    </div>
                </div>

                <div class="S1V1IAmHolder">
                    <div class="S1V1Slider" bind:clientHeight={iAmSliderHeight}
                         style="--i-am-slider-height: {iAmSliderHeight};">
                        <div class="CenterRowFlex S1V1SliderTextWrapper S1V1IAmWrapper">
                            I am
                        </div>
                        <div class="CenterRowFlex S1V1SliderTextWrapper S1V1IAmWrapper">
                            <ruby>
                                拙者
                                <rp>(</rp>
                                <rt>Sessha</rt>
                                <rp>)</rp>
                                &nbsp;は
                                <rp>(</rp>
                                <rt>wa</rt>
                                <rp>)</rp>
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

                <div class="S1V1EndHolder" bind:clientHeight={endHolderHeight}>
                    <div class="CenterRowFlex S1V1SliderTextWrapper S1V1EndTextWrapper"
                         style="--end-holder-height: {endHolderHeight}">
                        <ruby>
                            で
                            <rp>(</rp>
                            <rt>de</rt>
                            <rp>)</rp>
                            &nbsp;ござる
                            <rp>(</rp>
                            <rt>Gozaru</rt>
                            <rp>)</rp>
                            。
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

    .MPS1V1Holder {
        /* Default values for the vars passed */
        --vid-height: 9;
        --vid-width: 16;

        aspect-ratio: var(--vid-width) / var(--vid-height);

        position: relative;
    }

    .S1V1BgVideo {
        width: 100%;
        height: 100%;

        position: absolute;
        inset: 0 auto auto 0;
        z-index: 2;

        pointer-events: none;
    }

    .S1V1VideoOverlay {
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

    .S1V1GreetHolder {
        grid-row: 2 / 3;
        grid-column: 3 / 4;

        overflow: hidden;
    }

    .S1V1Slider {
        height: 200%;
        width: 100%;

        animation: slide-up-down 10s linear infinite;
    }

    .S1V1SliderTextWrapper {
        height: 50%;
        width: 100%;

        overflow: hidden;
        justify-content: flex-start;

        font-family: "Noto Sans Mono", sans-serif;
        text-align: left;
    }

    .S1V1GreetWrapper {
        font-size: calc(var(--greet-slider-height) / 4 * 1px);
    }

    .S1V1IAmHolder {
        grid-row: 4 / 5;
        grid-column: 3 / 4;

        overflow: hidden;
    }

    .S1V1IAmWrapper {
        font-size: calc(var(--i-am-slider-height) / 4 * 1px);
    }

    .S1V1EndHolder {
        grid-row: 10 / 11;
        grid-column: 3 / 4;

        overflow: hidden;
    }

    .S1V1EndTextWrapper {
        height: 100%;

        font-size: calc(var(--end-holder-height) / 2 * 1px);

        animation: pulse-opacity 10s linear infinite;
    }

    ruby {
        font-family: Meiryo, "Noto Serif JP", sans-serif;
    }

    ruby > rt {
        font-family: "Noto Sans Mono", sans-serif;
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