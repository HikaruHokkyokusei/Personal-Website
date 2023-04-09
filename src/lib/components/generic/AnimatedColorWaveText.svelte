<script lang="ts">
    export let data = "";
    export let fontWeight = 700;
    export let fontSizeEm = 4;
</script>

<span style="font: {fontWeight} {fontSizeEm}em/1 'Oswald', sans-serif;">
    {data}
</span>

<style lang="scss">
    span {
        text-transform: uppercase;
        letter-spacing: 0;
        text-shadow: 0 0 45px rgba(255, 255, 255, .4);
        -webkit-text-fill-color: transparent;

        background: url(/images/animated-text-fill.png) repeat-y;
        background-size: cover;
        -webkit-background-clip: text;
        background-clip: text;

        -webkit-animation: bg-animation 80s linear infinite;
        -webkit-backface-visibility: hidden;
    }

    @-webkit-keyframes bg-animation {
        0% {
            background-position: 0 50%;
        }
        100% {
            background-position: 100% 50%;
        }
    }
</style>