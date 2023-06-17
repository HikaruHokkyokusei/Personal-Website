// noinspection JSUnusedGlobalSymbols
import type { ChartConfiguration, ChartData } from "chart.js";

const commonProps = {
    "grouped": false,
    "barPercentage": 1,
    "borderRadius": 25,
    "borderSkipped": false,
    "borderWidth": 1
};

const OdysseyMythos: ChartData<"bar", any, string> = {
    labels: ["Education", "Freelancing", "Job"],
    datasets: [
        {
            data: [
                { x: ["2019-06-25", "2023-06-01"], y: "Education", title: "KIIT D. University", content: "" }
            ],
            backgroundColor: "rgb(253, 98, 131, 0.45)",
            borderColor: "rgb(253, 98, 131, 1)",
            ...commonProps
        },
        {
            data: [
                { x: ["2023-08-27", "2025-06-01"], y: "Education", title: "Stony Brook University", content: "" }
            ],
            backgroundColor: "rgb(253, 98, 131, 0.45)",
            borderColor: "rgb(253, 98, 131, 1)",
            ...commonProps
        },
        {
            data: [
                { x: ["2020-08-01", "2022-06-01"], y: "Freelancing", title: "Upwork Global", content: "" }
            ],
            backgroundColor: "rgb(74, 190, 190, 0.45)",
            borderColor: "rgb(74, 190, 190, 1)",
            ...commonProps
        },
        {
            data: [
                { x: ["2021-10-01", "2022-03-01"], y: "Job", title: "Highradius", content: "" }
            ],
            backgroundColor: "rgb(253, 158, 63, 0.45)",
            borderColor: "rgb(253, 158, 63, 1)",
            ...commonProps
        },
        {
            data: [
                { x: ["2022-06-06", "2023-07-01"], y: "Job", title: "Autodesk India Pvt. Ltd.", content: "" }
            ],
            backgroundColor: "rgb(253, 158, 63, 0.45)",
            borderColor: "rgb(253, 158, 63, 1)",
            ...commonProps
        }
    ]
};

const chartPlugins = {
    todayLineHorizontal: {
        id: "todayLine",
        afterDatasetsDraw: (chart: any) => {
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
            ctx.fillStyle = "rgba(255,129,0,0.66)";
            ctx.textAlign = "center";
            ctx.fillText("Today", xPos, top - 5);
            ctx.restore();
        }
    },
    todayLineVertical: {
        id: "todayLine",
        afterDatasetsDraw: (chart: any) => {
            const { ctx, chartArea: { left, right }, scales: { y } } = chart;
            const yPos = y.getPixelForValue(new Date());

            ctx.save();

            ctx.beginPath();
            ctx.lineWidth = 2;
            ctx.strokeStyle = "rgba(255,129,0,0.4)";
            ctx.setLineDash([6, 6]);
            ctx.moveTo(left, yPos);
            ctx.lineTo(right, yPos);
            ctx.stroke();
            ctx.restore();

            ctx.setLineDash([]);

            ctx.beginPath();
            ctx.lineWidth = 1;
            ctx.strokeStyle = "rgba(255,129,0,0.4)";
            ctx.fillStyle = "rgba(255,129,0,0.4)";
            ctx.moveTo(left - 5, yPos);
            ctx.lineTo(left - 15, yPos - 5);
            ctx.lineTo(left - 15, yPos + 5);
            ctx.closePath();
            ctx.stroke();
            ctx.fill();
            ctx.restore();

            ctx.font = "bold 12px sans-serif";
            ctx.fillStyle = "rgba(255,129,0,0.66)";
            ctx.textAlign = "center";
            ctx.fillText("Today", right + 20, yPos + 4);
            ctx.restore();
        }
    }
};
const chartOptionsPlugins = {
    tooltip: {
        callbacks: {
            title: (ctx: any) => {
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
            label: () => ""
        },
        displayColors: false
    }
};
const chartBaseScales: { x: any, y: any } = {
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
            color: "rgba(255, 255, 255, 0.4)"
        }
    },
    y: {
        labels: ["Education", "Freelancing", "Job"],
        beginAtZero: true,
        min: 0,
        max: 2,
        grid: {
            display: false
        },
        ticks: {
            color: "rgba(255, 255, 255, 0.4)"
        }
    }
};

export const generateOdysseyMythosChartConfigs = (orientation: "Horizontal" | "Vertical" = "Horizontal"): ChartConfiguration => {
    let isYIndex = orientation === "Horizontal";
    let scales: { x: any, x2?: any, y: any, y2?: any } = structuredClone(chartBaseScales);

    if (isYIndex) {
        scales["y2"] = { position: "right", ...scales.y };

        scales.x.ticks["padding"] = 15;
        scales.x.ticks["font"] = { "size": 15, "weight": "bold" };
        scales.y.ticks["font"] = { "size": 20 };
    } else {
        [scales.x, scales.y] = [scales.y, scales.x];
        scales["x2"] = { position: "top", ...scales.x };

        scales.y.ticks["padding"] = 15;
        scales.x.ticks["font"] = { "size": 15 };
        scales.y.ticks["font"] = { "size": 15 };
    }
    scales.x["position"] = "bottom";
    scales.y["position"] = "left";

    return {
        type: "bar",
        data: OdysseyMythos,
        options: {
            indexAxis: isYIndex ? "y" : "x",
            layout: {
                padding: isYIndex ? {} : { right: 40 }
            },
            maintainAspectRatio: false,
            plugins: {
                legend: {
                    display: false
                },
                tooltip: chartOptionsPlugins["tooltip"]
            },
            scales: scales,
            parsing: {
                xAxisKey: isYIndex ? "x" : "y",
                yAxisKey: isYIndex ? "y" : "x"
            }
        },
        plugins: [chartPlugins[`todayLine${orientation}`]]
    };
};