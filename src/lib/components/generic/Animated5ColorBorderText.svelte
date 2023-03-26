<script lang="ts">
    import { v4 as uuidV4 } from "uuid";

    const uniqueId = uuidV4();

    export let text: string;
    export let holderStyle: string = "";

    let wrapperHeight = 0, wrapperWidth = 0;
</script>

<div class="CenterRowFlex A5BCTWrapper" bind:clientHeight={wrapperHeight} bind:clientWidth={wrapperWidth}
     style="--font-size-base: {wrapperHeight}; {holderStyle}">
    <svg width="{wrapperWidth}" height="{wrapperHeight}">
        <symbol id="{uniqueId}">
            <text alignment-baseline="central" dominant-baseline="central" text-anchor="start" x="0" y="50%">
                {text}
            </text>
        </symbol>
        <use class="QuadColorTextAnimation" xlink:href="#{uniqueId}"></use>
        <use class="QuadColorTextAnimation" xlink:href="#{uniqueId}"></use>
        <use class="QuadColorTextAnimation" xlink:href="#{uniqueId}"></use>
        <use class="QuadColorTextAnimation" xlink:href="#{uniqueId}"></use>
        <use class="QuadColorTextAnimation" xlink:href="#{uniqueId}"></use>
    </svg>
</div>

<style>
    .A5BCTWrapper {
        justify-content: flex-start;

        font: calc((var(--font-size-base) / 1.15) * 1px) Open Sans, Impact;
        font-weight: bold;
        text-transform: uppercase;

        overflow: hidden;
    }

    .QuadColorTextAnimation {
        fill: none;

        stroke-width: 4;
        stroke-linejoin: round;

        stroke-dasharray: 70 330;

        -webkit-animation: stroke 6s infinite linear;
        animation: stroke 6s infinite linear;
    }

    .QuadColorTextAnimation:nth-child(5n + 1) {
        stroke: #f2385a;
        -webkit-animation-delay: -1.2s;
        animation-delay: -1.2s;
    }

    .QuadColorTextAnimation:nth-child(5n + 2) {
        stroke: #f5a503;
        -webkit-animation-delay: -2.4s;
        animation-delay: -2.4s;
    }

    .QuadColorTextAnimation:nth-child(5n + 3) {
        stroke: #e9f1df;
        -webkit-animation-delay: -3.6s;
        animation-delay: -3.6s;
    }

    .QuadColorTextAnimation:nth-child(5n + 4) {
        stroke: #56d9cd;
        -webkit-animation-delay: -4.8s;
        animation-delay: -4.8s;
    }

    .QuadColorTextAnimation:nth-child(5n + 5) {
        stroke: #3aa1bf;
        -webkit-animation-delay: -6s;
        animation-delay: -6s;
    }

    @-webkit-keyframes stroke {
        0% {
            stroke-dashoffset: 0;
        }

        100% {
            stroke-dashoffset: -400;
        }
    }

    @keyframes stroke {
        0% {
            stroke-dashoffset: 0;
        }

        100% {
            stroke-dashoffset: -400;
        }
    }
</style>