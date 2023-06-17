<script lang="ts">
    import { tweened } from "svelte/motion";
    import { linear } from "svelte/easing";
    import { onMount } from "svelte";
    import MediaQuery from "$lib/components/generic/MediaQuery.svelte";
    import { RenderChartOnCanvas } from "$lib/services/ChartJsService";
    import { generateOdysseyMythosChartConfigs } from "$lib/configs/OdysseyMythos";
    import { genericDataStore } from "$lib/stores/GenericDataStore";

    let intervalId;
    let titleTextType: "en" | "jp" = "en";
    let titleOpacity = tweened(1, {
        duration: 500,
        easing: linear
    });

    onMount(() => {
        setTimeout(() => {
            switchTitleText();
            intervalId = setInterval(switchTitleText, 5000);
        }, 4500);

        return () => {
            clearInterval(intervalId);
        };
    });

    const switchTitleText = async () => {
        await titleOpacity.set(0);

        if (titleTextType === "en") {
            titleTextType = "jp";
        } else {
            titleTextType = "en";
        }

        await titleOpacity.set(1);
    };
</script>

<section id="chronicles">
    <div class="S3V1Wrapper" style:color="{$genericDataStore.theme.onBackground}">
        <div class="CenterRowFlex S3V1TitleHolder" style:opacity="{$titleOpacity}">
            {#if titleTextType === "en"}
                My odyssey's mythos
            {:else}
                <ruby style="font-size: 2.55rem">
                    俺様
                    <rt>&hairsp;Oresama</rt>
                    &nbsp;の
                    <rt>&nbsp;&thinsp;no</rt>
                    &nbsp;手記
                    <rt>&thinsp;&thinsp;Karte</rt>
                </ruby>
            {/if}
        </div>

        <div style="height: 30px; width: 100%;"></div>

        <MediaQuery query="only screen and (min-aspect-ratio: 20 / 17)" let:matches>
            <div class="CenterRowFlex TimelineWrapper">
                <div class="TimelineCanvasHolder" style="height: {matches ? '450px' : '1000px'}; width: 85%">
                    {#if matches}
                        <canvas use:RenderChartOnCanvas={generateOdysseyMythosChartConfigs("Horizontal")}></canvas>
                    {:else}
                        <canvas use:RenderChartOnCanvas={generateOdysseyMythosChartConfigs("Vertical")}></canvas>
                    {/if}
                </div>
            </div>
        </MediaQuery>

        <div style="height: 80px; width: 100%;"></div>
    </div>
</section>

<style>
    .S3V1Wrapper {
        height: fit-content;
        width: 100%;

        --font-base: 0;
    }

    .S3V1TitleHolder {
        height: 100px;
        width: 100%;

        overflow: hidden;

        font: 3rem "Noto Sans Mono", sans-serif;
        text-align: center;
    }

    .TimelineWrapper {
        width: 100%;
        min-width: 550px;
    }

    .TimelineCanvasHolder {
        background: rgba(75, 50, 125, 0.50);

        padding: 25px;
        border-radius: 25px;
    }
</style>