import { BarController, BarElement, CategoryScale, Chart, LinearScale, TimeScale, Tooltip } from "chart.js";
import "chartjs-adapter-luxon";

Chart.register(BarController, BarElement, CategoryScale, LinearScale, TimeScale, Tooltip);

export const RenderChartOnCanvas = (node: HTMLCanvasElement, parameters: any) => {
    const _chart = new Chart(node, parameters);

    return {
        "update": (updatedParameters: any) => {
            _chart.data = updatedParameters.data;
            _chart.update();
        },
        "destroy": () => {
            _chart.destroy();
        }
    };
};