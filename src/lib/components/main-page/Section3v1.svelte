<script lang="ts">
    import { tweened } from "svelte/motion";
    import { linear } from "svelte/easing";
    import { onMount } from "svelte";
    import MediaQuery from "../../components/generic/MediaQuery.svelte";
    import { RenderChartOnCanvas } from "../../services/ChartJsService";
    import { OdysseyMythos } from "../../configs/OdysseyMythos";

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

    const chartPlugins = {
        todayLine: {
            id: "todayLine",
            afterDatasetsDraw: (chart, args, pluginOptions) => {
                const { ctx, chartArea: { top, bottom }, scales: { x } } = chart;
                const xPos = x.getPixelForValue(new Date());

                ctx.save();

                ctx.beginPath();
                ctx.lineWidth = 2;
                ctx.strokeStyle = "rgba(255,129,0,0.4)";
                ctx.setLineDash([6, 6]);
                ctx.moveTo(xPos, top);
                ctx.lineTo(xPos, bottom);
                ctx.stroke();
                ctx.restore();

                ctx.setLineDash([]);

                ctx.beginPath();
                ctx.lineWidth = 1;
                ctx.strokeStyle = "rgba(255,129,0,0.4)";
                ctx.fillStyle = "rgba(255,129,0,0.4)";
                ctx.moveTo(xPos, bottom + 5);
                ctx.lineTo(xPos - 6, bottom + 20);
                ctx.lineTo(xPos + 6, bottom + 20);
                ctx.closePath();
                ctx.stroke();
                ctx.fill();
                ctx.restore();

                ctx.font = "bold 12px sans-serif";
                ctx.fillStyle = "rgba(255,129,0,0.4)";
                ctx.textAlign = "center";
                ctx.fillText("Today", xPos, top - 5);
                ctx.restore();
            }
        }
    };
    const optionsPlugins = {
        tooltip: {
            callbacks: {
                title: (ctx) => {
                    const startDate = new Date(ctx[0].raw.x[0]).toLocaleString([], {
                        year: "numeric",
                        month: "short",
                        day: "numeric",
                        hour12: true
                    });
                    const endDate = new Date(ctx[0].raw.x[1]).toLocaleString([], {
                        year: "numeric",
                        month: "short",
                        day: "numeric",
                        hour12: true
                    });

                    return [
                        ctx[0].raw.title,
                        `${startDate} - ${endDate}`
                    ];
                },
                label: (ctx) => {
                    return null;
                }
            },
            displayColors: false
        }
    };

    const chartConfigs = {
        type: "bar",
        data: OdysseyMythos,
        options: {
            indexAxis: "y",
            layout: {
                padding: {
                    top: 20
                }
            },
            maintainAspectRatio: false,
            plugins: {
                legend: {
                    display: false
                },
                tooltip: optionsPlugins["tooltip"]
            },
            scales: {
                x: {
                    min: "2019-04-01",
                    max: ((dt: Date) => {
                        return `${dt.getUTCFullYear() + 1}-01-01`;
                    })(new Date(Date.now())),
                    type: "time",
                    time: {
                        unit: "quarter"
                    },
                    grid: {
                        color: "rgba(191, 255, 169, 0.15)"
                    },
                    ticks: {
                        padding: 15,
                        font: {
                            size: 15,
                            weight: "bold"
                        }
                    }
                },
                y: {
                    labels: ["Education", "Freelancing", "Job"],
                    min: 0,
                    max: 2,
                    grid: {
                        display: false
                    },
                    position: "left",
                    ticks: {
                        font: {
                            size: 20,
                            weight: "bold"
                        }
                    }
                },
                y2: {
                    labels: ["Education", "Freelancing", "Job"],
                    min: 0,
                    max: 2,
                    grid: {
                        display: false
                    },
                    position: "right",
                    ticks: {
                        font: {
                            size: 20,
                            weight: "bold"
                        }
                    }
                }
            }
        },
        plugins: [chartPlugins["todayLine"]]
    };
</script>

<section id="chronicles">
    <div class="S3V1Wrapper">
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

        <MediaQuery query="(min-width: 1600px)" let:matches>
            <div class="CenterRowFlex TimelineWrapper">
                <div class="TimelineCanvasHolder">
                    <canvas use:RenderChartOnCanvas={chartConfigs}></canvas>
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

        background: radial-gradient(ellipse at top, #1d2a3a 0%, #091023 100%);

        --font-base: 0;
    }

    .S3V1TitleHolder {
        height: 125px;
        width: 100%;

        overflow: hidden;

        font: 3rem "Noto Sans Mono", sans-serif;
        text-align: center;
    }

    .TimelineWrapper {
        width: 100%;
    }

    .TimelineCanvasHolder {
        width: 75%;
        aspect-ratio: 20 / 5;

        background: #002a26;

        padding: 25px;
        border-radius: 25px;
    }
</style>