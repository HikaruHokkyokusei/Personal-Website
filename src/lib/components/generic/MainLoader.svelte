<script lang="ts">
    export let sideLengthEm: number = 10;
    export let tiltAngleDeg: number = 20;
    export let rotationDurationS: number = 8;
    export let ringThicknessEm: number = 1;
    export let ringGapEm: number = 1;

    let ringRadius: number;
    $: ringRadius = sideLengthEm / 2;
</script>

<div class="CenterRowFlex LoaderWrapper">
    <div class="CenterColumnFlex LoaderRingHolder"
         style="--sideLength: {sideLengthEm}em; --tiltAngle: {tiltAngleDeg}deg; --rotationDuration: {rotationDurationS}s;
         --ringThickness: {ringThicknessEm}em; --ringRadius: {ringRadius}em; --ringGap: {ringGapEm}em;"
    >
        <div class="LoaderRing">
            <div class="LoaderRingSector" style="--n: 1;">L</div>
            <div class="LoaderRingSector" style="--n: 2;">o</div>
            <div class="LoaderRingSector" style="--n: 3;">a</div>
            <div class="LoaderRingSector" style="--n: 4;">d</div>
            <div class="LoaderRingSector" style="--n: 5;">i</div>
            <div class="LoaderRingSector" style="--n: 6;">n</div>
            <div class="LoaderRingSector" style="--n: 7;">g</div>
            <div class="LoaderRingSector" style="--n: 8;">.</div>
            <div class="LoaderRingSector" style="--n: 9;">.</div>
            <div class="LoaderRingSector" style="--n: 10;">.</div>
            <div class="LoaderRingSector" style="--n: 11;"></div>
            <div class="LoaderRingSector" style="--n: 12;"></div>
            <div class="LoaderRingSector" style="--n: 13;"></div>
            <div class="LoaderRingSector" style="--n: 14;"></div>
            <div class="LoaderRingSector" style="--n: 15;"></div>
            <div class="LoaderRingSector" style="--n: 16;"></div>
            <div class="LoaderRingSector" style="--n: 17;"></div>
            <div class="LoaderRingSector" style="--n: 18;"></div>
            <div class="LoaderRingSector" style="--n: 19;"></div>
            <div class="LoaderRingSector" style="--n: 20;"></div>
            <div class="LoaderRingSector" style="--n: 21;"></div>
            <div class="LoaderRingSector" style="--n: 22;"></div>
            <div class="LoaderRingSector" style="--n: 23;"></div>
            <div class="LoaderRingSector" style="--n: 24;"></div>
            <div class="LoaderRingSector" style="--n: 25;"></div>
            <div class="LoaderRingSector" style="--n: 26;"></div>
            <div class="LoaderRingSector" style="--n: 27;"></div>
            <div class="LoaderRingSector" style="--n: 28;"></div>
            <div class="LoaderRingSector" style="--n: 29;"></div>
            <div class="LoaderRingSector" style="--n: 30;"></div>
            <div class="LoaderRingSector" style="--n: 31;"></div>
            <div class="LoaderRingSector" style="--n: 32;"></div>
            <div class="LoaderRingSector" style="--n: 33;"></div>
            <div class="LoaderRingSector" style="--n: 34;"></div>
            <div class="LoaderRingSector" style="--n: 35;"></div>
            <div class="LoaderRingSector" style="--n: 36;"></div>
        </div>
        <div style="width: 100%; height: {ringGapEm}em;"></div>
        <div class="LoaderRing">
            <div class="LoaderRingSector" style="--n: 1;">L</div>
            <div class="LoaderRingSector" style="--n: 2;">o</div>
            <div class="LoaderRingSector" style="--n: 3;">a</div>
            <div class="LoaderRingSector" style="--n: 4;">d</div>
            <div class="LoaderRingSector" style="--n: 5;">i</div>
            <div class="LoaderRingSector" style="--n: 6;">n</div>
            <div class="LoaderRingSector" style="--n: 7;">g</div>
            <div class="LoaderRingSector" style="--n: 8;">.</div>
            <div class="LoaderRingSector" style="--n: 9;">.</div>
            <div class="LoaderRingSector" style="--n: 10;">.</div>
            <div class="LoaderRingSector" style="--n: 11;"></div>
            <div class="LoaderRingSector" style="--n: 12;"></div>
            <div class="LoaderRingSector" style="--n: 13;"></div>
            <div class="LoaderRingSector" style="--n: 14;"></div>
            <div class="LoaderRingSector" style="--n: 15;"></div>
            <div class="LoaderRingSector" style="--n: 16;"></div>
            <div class="LoaderRingSector" style="--n: 17;"></div>
            <div class="LoaderRingSector" style="--n: 18;"></div>
            <div class="LoaderRingSector" style="--n: 19;"></div>
            <div class="LoaderRingSector" style="--n: 20;"></div>
            <div class="LoaderRingSector" style="--n: 21;"></div>
            <div class="LoaderRingSector" style="--n: 22;"></div>
            <div class="LoaderRingSector" style="--n: 23;"></div>
            <div class="LoaderRingSector" style="--n: 24;"></div>
            <div class="LoaderRingSector" style="--n: 25;"></div>
            <div class="LoaderRingSector" style="--n: 26;"></div>
            <div class="LoaderRingSector" style="--n: 27;"></div>
            <div class="LoaderRingSector" style="--n: 28;"></div>
            <div class="LoaderRingSector" style="--n: 29;"></div>
            <div class="LoaderRingSector" style="--n: 30;"></div>
            <div class="LoaderRingSector" style="--n: 31;"></div>
            <div class="LoaderRingSector" style="--n: 32;"></div>
            <div class="LoaderRingSector" style="--n: 33;"></div>
            <div class="LoaderRingSector" style="--n: 34;"></div>
            <div class="LoaderRingSector" style="--n: 35;"></div>
            <div class="LoaderRingSector" style="--n: 36;"></div>
        </div>
    </div>
</div>

<style>
    :root {
        --sideLength: 10em;
        --tiltAngle: 20deg;
        --rotationDuration: 8s;

        --ringThickness: 1em;
        --ringRadius: 10em;
        --ringGap: 0.5em;

        --n: 0;
    }

    .LoaderWrapper {
        width: 100%;
        height: 100%;

        color: #3df1f1;
        font: 1em Dosis, sans-serif;
        font-size: calc((16px + (24 - 16) * (100vw - 320px) / (1280 - 320) * 10 / 7));
    }

    .LoaderWrapper > * {
        box-sizing: border-box;
        border: 0;
        margin: 0;
        padding: 0;
    }

    .LoaderRingHolder {
        height: calc((2 * var(--ringThickness)) + var(--ringGap));
        width: var(--sideLength);

        background: transparent;

        line-height: 1.5;

        perspective: 40em;
        transform-style: preserve-3d;

        animation: tiltSpin var(--rotationDuration) linear infinite;
    }

    .LoaderRing {
        height: var(--ringThickness);
        width: calc((2 * 3.1416 * var(--ringRadius)) / 30);

        font-size: 1.5em;
        position: relative;

        transform-style: preserve-3d;

        animation: spin calc(var(--rotationDuration) / 2) linear infinite;
    }

    .LoaderRing:nth-child(1) {
        animation-direction: normal;
    }

    .LoaderRing:nth-child(3) {
        animation-direction: reverse;
    }

    .LoaderRingSector, .LoaderRingSector:empty:before {
        display: inline-block;
    }

    .LoaderRingSector:empty:before {
        content: "";
        height: 100%;
        width: 100%;
        background: linear-gradient(transparent 45%, currentColor 45%, currentColor 55%, transparent 55%);
    }

    .LoaderRingSector {
        height: 100%;
        width: 100%;

        position: absolute;
        top: 0;
        left: 0;

        font-weight: 600;
        text-align: center;
        text-transform: uppercase;

        transform: rotateY(calc((var(--n) - 1) * 10deg)) translateZ(var(--ringRadius));
    }

    @keyframes tiltSpin {
        from {
            transform: rotateY(0) rotateX(var(--tiltAngle));
        }
        to {
            transform: rotateY(1turn) rotateX(var(--tiltAngle));
        }
    }

    @keyframes spin {
        from {
            transform: rotateY(0);
        }
        to {
            transform: rotateY(1turn);
        }
    }
</style>