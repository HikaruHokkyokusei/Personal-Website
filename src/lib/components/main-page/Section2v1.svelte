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
            <div class="S2V1ContentHolder" style:width="{matches ? '72.5%' : '100%'}">
                <div class="CenterRowFlex" style:width="{matches ? '45%' : '100%'}">
                    <img src="/images/profile-pic-ganyu.png" alt="Ganyu" style="width: 400px; height: 400px;"
                         height="2278" width="2278">
                </div>
                {#if !matches}
                    <div style="height: 50px; width: 100%;"></div>
                {/if}
                <div class="CenterRowFlex" style:width="{matches ? '55%' : '100%'}">
                    <div class="CenterRowFlex S2V1AboutContent">
                        I am a highly experienced software developer, possessing a diverse set of skills that enable me
                        to create innovative and efficient solutions for any project.
                        <br><br>
                        My expertise extends to a wide range of programming languages, frameworks and
                        libraries, allowing me to tackle any challenge with ease. Additionally, I have extensive
                        experience working with the AWS Cloud platform and Networking, empowering me to develop scalable
                        and reliable applications that meet the needs of modern businesses and end-users alike.
                        <br><br>
                        My exceptional analytical and problem-solving skills, combined with my ability to work
                        independently or as part of a team, make me an ideal candidate for any software development
                        role.
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

        font-size: 1.6rem;
        text-align: justify;
    }
</style>