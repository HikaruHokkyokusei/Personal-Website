<script lang="ts">
    let wrapperHeight = 0, wrapperWidth = 0;
</script>

<svelte:head>
    <link rel='stylesheet' href='https://fonts.googleapis.com/css?family=Lato:300,400,700'/>
</svelte:head>

<!--suppress HtmlUnknownAttribute -->
<div class="SOStarsWrapper" bind:clientHeight={wrapperHeight} bind:clientWidth={wrapperWidth}
     style:--wrapper-height="{wrapperHeight}" style:--wrapper-width="{wrapperWidth}">
    <div class="SOStars1"></div>
    <div class="SOStars2"></div>
    <div class="SOStars3"></div>
    <div class="SOTitle">
        <span>KON'NICHIWA</span>
        <br/>
        <span>GOJIN</span>
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

    .SOStarsWrapper {
        height: 100%;
        width: 100%;

        background: radial-gradient(ellipse at bottom, #1B2735 0%, #090A0F 100%);

        overflow: hidden;
    }

    .SOStars1 {
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

    .SOStars2 {
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

    .SOStars3 {
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

    .SOTitle {
        position: absolute;
        top: 50%;
        left: 0;
        right: 0;

        color: #FFF;
        text-align: center;
        font-family: 'lato', sans-serif;
        font-weight: 300;
        font-size: 50px;
        letter-spacing: 10px;

        margin-top: -60px;
        padding-left: 10px;

        span {
            background: -webkit-linear-gradient(white, #38495a);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }
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