import { useState } from "react";
import { DisplayCard } from "../DisplayCard";
import { LineChart } from "../LineChart";
import { PieChart } from "../PieChart";
import styles from "./ChartCards.module.scss";

export const ChartCard: React.FC = () => {
  const [isDspLineChart, setIsDspLineChart] = useState(true);
  const [isDspPieChart, setIsDspPieChart] = useState(false);
  return (
    <div className={styles.ChartCards}>
      <DisplayCard style={styles.DisplayCard}>
        <div>
          <button
            className={isDspLineChart ? styles.SelectedButton : ""}
            onClick={() => {
              setIsDspLineChart((prev) => !prev);
              setIsDspPieChart(false);
            }}
          >
            感染者数
          </button>
          <button
            className={isDspPieChart ? styles.SelectedButton : ""}
            onClick={() => {
              setIsDspPieChart((prev) => !prev);
              setIsDspLineChart(false);
            }}
          >
            世間の関心
          </button>
        </div>
        <div className={styles.ChartContainer}>
          {isDspLineChart ? <LineChart /> : <PieChart />}
        </div>
      </DisplayCard>
    </div>
  );
};
