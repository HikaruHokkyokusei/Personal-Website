<script lang="ts">
    let wrapperHeight = 0, wrapperWidth = 0;
</script>

<!--suppress HtmlUnknownAttribute -->
<div class="MovingStarsWrapper" bind:clientHeight={wrapperHeight} bind:clientWidth={wrapperWidth}
     style:--wrapper-height="{wrapperHeight}" style:--wrapper-width="{wrapperWidth}">
    <div class="MovingStars1"></div>
    <div class="MovingStars2"></div>
    <div class="MovingStars3"></div>

    <div class="MovingStarsContentWrapper">
        <slot></slot>
    </div>
</div>

<style lang="scss">
    /* For JetBrains, install the sass plugin and enable the file writer to remove any IDE errors */

    /* n is number of stars required */
    @function multiple-box-shadow ($n) {
        $value: 'calc(var(--wrapper-width) * #{random(10000)} / 10000 * 1px) calc(var(--wrapper-height) * #{random(10000)} / 10000 * 1px) #FFF';
        @for $i from 2 through $n {
            $value: '#{$value}, calc(var(--wrapper-width) * #{random(10000)} / 10000 * 1px) calc(var(--wrapper-height) * #{random(10000)} / 10000 * 1px) #FFF';
        }

        @return unquote($value);
    }

    $shadows-small: multiple-box-shadow(350);
    $shadows-medium: multiple-box-shadow(250);
    $shadows-big: multiple-box-shadow(125);

    .MovingStarsWrapper {
        height: 100%;
        width: 100%;

        background: linear-gradient(
                        transparent 20%,
                        rgba(136, 72, 180, 0.4) 60%,
                        rgba(126, 57, 124, 0.66) 77%,
                        rgba(215, 121, 139, 0.7) 100%
        ),
        radial-gradient(ellipse at bottom, #1f2c3b 0%, #090A0F 60%, #080a10 100%);

        overflow: hidden;
    }

    .MovingStars1 {
        height: 1px;
        width: 1px;

        background: transparent;
        box-shadow: $shadows-small;

        animation: BackgroundStarAnimation 50s linear infinite;

        &:after {
            content: "";
            height: 1px;
            width: 1px;

            position: absolute;
            top: calc(var(--wrapper-height) * 1px);

            background: transparent;
            box-shadow: $shadows-small;
        }
    }

    .MovingStars2 {
        height: 2px;
        width: 2px;

        background: transparent;
        box-shadow: $shadows-medium;

        animation: BackgroundStarAnimation 400s linear infinite reverse;

        &:after {
            content: "";
            height: 2px;
            width: 2px;

            position: absolute;
            top: calc(var(--wrapper-height) * 1px);

            background: transparent;
            box-shadow: $shadows-medium;
        }
    }

    .MovingStars3 {
        height: 3px;
        width: 3px;

        background: transparent;
        box-shadow: $shadows-big;

        animation: BackgroundStarAnimation 150s linear infinite;

        &:after {
            content: "";
            height: 3px;
            width: 3px;

            position: absolute;
            top: calc(var(--wrapper-height) * 1px);

            background: transparent;
            box-shadow: $shadows-big;
        }
    }

    .MovingStarsContentWrapper {
        height: 100%;
        width: 100%;

        position: absolute;
        top: 0;
        left: 0;
    }

    @keyframes BackgroundStarAnimation {
        from {
            transform: translateY(0px);
        }
        to {
            transform: translateY(calc(var(--wrapper-height) * -1px));
        }
    }
</style>