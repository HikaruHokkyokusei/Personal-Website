<script lang="ts">
    import MediaQuery from "$lib/components/generic/MediaQuery.svelte";
    import MovingStars from "$lib/components/generic/MovingStars.svelte";
    import AnimatedColorWaveText from "$lib/components/generic/AnimatedColorWaveText.svelte";
    import { personalData } from "$lib/configs/PersonalData";

    let wrapperHeight = 0, wrapperWidth = 0;
    let greetSliderHeight = 0, iAmSliderHeight = 0, endHolderHeight = 0;
</script>

<MovingStars>
    <MediaQuery query="only screen and (min-aspect-ratio: 8 / 10)" let:matches>
        <div class="S1V2ContentWrapper" bind:clientHeight={wrapperHeight} bind:clientWidth={wrapperWidth}
             style:--font-base="{matches ? wrapperHeight : wrapperWidth}"
             style:font-weight="{matches ? 500 : 800}">
            <div class="S1V2GreetHolder">
                <div class="S1V2Slider" bind:clientHeight={greetSliderHeight}
                     style="--greet-slider-height: {greetSliderHeight};">
                    <div class="CenterRowFlex S1V2SliderTextWrapper S1V2GreetWrapper">
                        <span>
                            &nbsp;Hello Monsieur,
                        </span>
                    </div>
                    <div class="CenterRowFlex S1V2SliderTextWrapper S1V2GreetWrapper">
                        <span>
                            &nbsp;
                            <ruby>
                                こんにちは
                                <rt><span>Kon'nichiwa</span></rt>
                                &nbsp;御仁
                                <rt><span>Gojin</span></rt>
                                、
                            </ruby>
                        </span>
                    </div>
                </div>
            </div>

            <div class="S1V2IAmHolder">
                <div class="S1V2Slider" bind:clientHeight={iAmSliderHeight}
                     style="--i-am-slider-height: {iAmSliderHeight};">
                    <div class="CenterRowFlex S1V2SliderTextWrapper S1V2IAmWrapper">
                        <span>
                            I am
                        </span>
                    </div>
                    <div class="CenterRowFlex S1V2SliderTextWrapper S1V2IAmWrapper">
                        <span>
                            <ruby>
                                拙者
                                <rt><span>Sessha</span></rt>
                                &nbsp;は
                                <rt><span>wa</span></rt>
                            </ruby>
                        </span>
                    </div>
                </div>
            </div>

            <div class="CenterRowFlex" style="grid-row: 5 / 6; grid-column: 1 / 2; opacity: 0.75;">
                <AnimatedColorWaveText fontSizeEm="{(matches ? wrapperHeight : wrapperWidth) / 175}"
                                       data="{personalData.firstName}">
                </AnimatedColorWaveText>
            </div>
            <div class="CenterRowFlex" style="grid-row: 6 / 7; grid-column: 1 / 2; opacity: 0.75;">
                <AnimatedColorWaveText fontSizeEm="{(matches ? wrapperHeight : wrapperWidth) / 175}"
                                       data="{personalData.lastName}">
                </AnimatedColorWaveText>
            </div>

            <div class="S1V2EndHolder" bind:clientHeight={endHolderHeight}>
                <div class="CenterRowFlex S1V2SliderTextWrapper S1V2EndTextWrapper"
                     style="--end-holder-height: {endHolderHeight}">
                    <span>
                        &nbsp;
                        <ruby>
                            で
                            <rt><span>de</span></rt>
                            &nbsp;ござる
                            <rt><span>Gozaru</span></rt>
                            。
                        </ruby>
                    </span>
                </div>
            </div>
        </div>
    </MediaQuery>
</MovingStars>

<style lang="scss">
    .S1V2ContentWrapper {
        height: 100%;

        margin: auto;

        display: grid;
        grid-template: auto auto
                       max(65px, 8%) max(65px, 8%)
                       max(100px, 10%) max(100px, 10%)
                       max(65px, 8%)
                       auto auto / 100%;


        font-family: "Noto Sans Mono", sans-serif;
    }

    .S1V2GreetHolder {
        height: 100%;
        width: 100%;

        grid-row: 3 / 4;
        grid-column: 1 / 2;

        overflow: hidden;
    }

    .S1V2Slider {
        height: 200%;
        width: 100%;

        animation: slide-up-down 10s linear infinite;
    }

    .S1V2SliderTextWrapper {
        height: 50%;
        width: 100%;

        overflow: hidden;

        text-align: center;
    }

    .S1V2GreetWrapper {
        font-size: calc(var(--font-base) / 366 * 1em);
    }

    .S1V2IAmHolder {
        height: 100%;
        width: 100%;

        grid-row: 4 / 5;
        grid-column: 1 / 2;

        overflow: hidden;
    }

    .S1V2IAmWrapper {
        font-size: calc(var(--font-base) / 366 * 1em);
    }

    .S1V2EndHolder {
        height: 100%;
        width: 100%;

        grid-row: 7 / 8;
        grid-column: 1 / 2;

        overflow: hidden;
    }

    .S1V2EndTextWrapper {
        height: 100%;

        font-size: calc(var(--font-base) / 366 * 1em);

        animation: pulse-opacity 10s linear infinite;
    }

    span {
        background: -webkit-linear-gradient(white, #38495a);
        -webkit-background-clip: text;
        -webkit-text-fill-color: transparent;
    }

    ruby {
        font-family: "ヒラギノ角ゴ Pro W3", "Hiragino Kaku Gothic Pro", Osaka, メイリオ, Meiryo, "ＭＳ Ｐゴシック", "MS PGothic", "ＭＳ ゴシック", "MS Gothic", "Noto Sans CJK JP", TakaoPGothic, sans-serif;
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