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
import styles from "./LineChart.module.scss";

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
      text: "各国の累計感染者数",
      color: "rgb(255, 255, 255, 0.8)",
    },
    labels: {
      color: "rgb(255, 255, 255, 0.8)",
      title: {
        color: "rgb(255, 255, 255, 0.8)",
      },
    },
  },
  scales: {
    x: {
      grid: {
        color: "rgb(255, 255, 255, 0.1)",
      },
      ticks: {
        color: "rgb(255, 255, 255, 0.8)",
      },
    },
    y: {
      grid: {
        color: "rgb(255, 255, 255, 0.1)",
      },
      ticks: {
        color: "rgb(255, 255, 255, 0.8)",
      },
    },
  },
};

type Props = {
  lineChartDatas: ChartData<"line", number[], string>;
};
const _LineChart: React.FC<Props> = ({ lineChartDatas }) => {
  return (
    <Line
      options={options}
      data={lineChartDatas}
      className={styles.LineChart}
    />
  );
};

export const LineChart = React.memo(_LineChart);
