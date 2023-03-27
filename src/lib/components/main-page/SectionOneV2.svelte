<script lang="ts">

</script>

<svelte:head>
    <link rel='stylesheet' href='https://fonts.googleapis.com/css?family=Lato:300,400,700'/>
</svelte:head>

<div class="SOStarsWrapper">
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
        $value: '#{random(2000)}px #{random(2000)}px #FFF';
        @for $i from 2 through $n {
            $value: '#{$value} , #{random(2000)}px #{random(2000)}px #FFF';
        }

        @return unquote($value);
    }

    $shadows-small: multiple-box-shadow(1000);
    $shadows-medium: multiple-box-shadow(400);
    $shadows-big: multiple-box-shadow(200);

    .SOStarsWrapper {
        width: 100%;
        height: 100%;
        background: radial-gradient(ellipse at bottom, #1B2735 0%, #090A0F 100%);
        overflow: hidden;
    }

    .SOStars1 {
        width: 1px;
        height: 1px;
        background: transparent;
        box-shadow: $shadows-small;
        animation: BackgroundStarAnimation 50s linear infinite;

        &:after {
            content: "";
            position: absolute;
            top: 2000px;
            width: 1px;
            height: 1px;
            background: transparent;
            box-shadow: $shadows-small;
        }
    }

    .SOStars2 {
        width: 2px;
        height: 2px;
        background: transparent;
        box-shadow: $shadows-medium;
        animation: BackgroundStarAnimation 100s linear infinite;

        &:after {
            content: "";
            position: absolute;
            top: 2000px;
            width: 2px;
            height: 2px;
            background: transparent;
            box-shadow: $shadows-medium;
        }
    }

    .SOStars3 {
        width: 3px;
        height: 3px;
        background: transparent;
        box-shadow: $shadows-big;
        animation: BackgroundStarAnimation 150s linear infinite;

        &:after {
            content: "";
            position: absolute;
            top: 2000px;
            width: 3px;
            height: 3px;
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
            transform: translateY(-2000px);
        }
    }
</style>