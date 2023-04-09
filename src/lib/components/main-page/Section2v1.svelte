<script lang="ts">
    import { tweened } from "svelte/motion";
    import { linear } from "svelte/easing";
    import { onMount } from "svelte";
    import MediaQuery from "../../components/generic/MediaQuery.svelte";

    let intervalId;
    let titleTextType: "en" | "jp" = "en";
    let titleOpacity = tweened(1, {
        duration: 500,
        easing: linear
    });

    const switchTitleText = async () => {
        await titleOpacity.set(0);

        if (titleTextType === "en") {
            titleTextType = "jp";
        } else {
            titleTextType = "en";
        }

        await titleOpacity.set(1);
    }

    onMount(() => {
        setTimeout(() => {
            switchTitleText();
            intervalId = setInterval(switchTitleText, 5000);
        }, 4500);

        return () => {
            clearInterval(intervalId);
        };
    });
</script>

<section id="about">
    <div class="S2V1Wrapper">
        <div class="CenterRowFlex S2V1TitleHolder" style:opacity="{$titleOpacity}">
            {#if titleTextType === "en"}
                About me
            {:else}
                <ruby style="font-size: 2.55rem">
                    吾輩
                    <rt>&hairsp;Wagahai</rt>
                    &nbsp;に
                    <rt>&nbsp;&thinsp;ni</rt>
                    &nbsp;ついて
                    <rt>&thinsp;&thinsp;tsuite</rt>
                </ruby>
            {/if}
        </div>

        <div style="height: 30px; width: 100%;"></div>

        <MediaQuery query="(min-width: 1650px)" let:matches>
            <div class="S2V1ContentHolder" style:width="{matches ? '60%' : '100%'}">
                <div class="CenterRowFlex" style:width="{matches ? '50%' : '100%'}">
                    <img src="/images/profile-pic-ganyu.png" alt="Ganyu" style="width: 400px; height: 400px;"
                         height="2278" width="2278">
                </div>
                {#if !matches}
                    <div style="height: 25px; width: 100%;"></div>
                {/if}
                <div class="CenterRowFlex" style:width="{matches ? '50%' : '100%'}">
                    <div class="CenterRowFlex S2V1AboutContent">
                        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut
                        labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco
                        laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in
                        voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat
                        non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
                    </div>
                </div>
            </div>
        </MediaQuery>

        <div style="height: 80px; width: 100%;"></div>
    </div>
</section>

<style>
    .S2V1Wrapper {
        height: fit-content;
        width: 100%;

        background: radial-gradient(ellipse at top, #1d2a3a 0%, #091023 100%);

        --font-base: 0;
    }

    .S2V1TitleHolder {
        height: 125px;
        width: 100%;

        overflow: hidden;

        font: 3rem "Noto Sans Mono", sans-serif;
        text-align: center;
    }

    .S2V1ContentHolder {
        width: 100%;
        max-width: 1500px;

        margin: 0 auto;

        display: flex;
        flex-wrap: wrap;
    }

    .S2V1AboutContent {
        width: 85%;

        font-size: 1.75rem;
        text-align: justify;
    }
</style>