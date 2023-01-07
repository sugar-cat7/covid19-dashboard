import React from "react";
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  ChartData,
} from "chart.js";
import { Line } from "react-chartjs-2";
ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
);

export const options = {
  responsive: true,
  plugins: {
    legend: {
      position: "top" as const,
    },
    title: {
      display: true,
      text: "Chart.js Line Chart",
    },
  },
};

type Props = {
  lineChartDatas: ChartData<"line", number[], string>;
};
const _LineChart: React.FC<Props> = ({ lineChartDatas }) => {
  return <Line options={options} data={lineChartDatas} />;
};

export const LineChart = React.memo(_LineChart);
