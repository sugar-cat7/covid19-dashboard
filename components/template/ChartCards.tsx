import { ChartData } from "chart.js";
import React, {
  Dispatch,
  SetStateAction,
  useCallback,
  useMemo,
  useState,
} from "react";
import { fetcher } from "../../lib/api";
import { InfectedPatients } from "../../types/types";
import { DisplayCard } from "../DisplayCard";
import { LineChart } from "../LineChart";
import { PieChart } from "../PieChart";
import styles from "./ChartCards.module.scss";
import useSWR from "swr";
import Select from "react-select";
import { selectOptions } from "../../data/master.js";

type Props = {
  selectedValue: {
    value: string;
    label: string;
  }[];
  setSelectedValue: Dispatch<
    SetStateAction<
      {
        value: string;
        label: string;
      }[]
    >
  >;
};
const _ChartCard: React.FC<Props> = ({ selectedValue, setSelectedValue }) => {
  const [isDspLineChart, setIsDspLineChart] = useState(true);
  const [isDspPieChart, setIsDspPieChart] = useState(false);

  const { data, error, isLoading } = useSWR<InfectedPatients[]>(
    "/infected_patients?from=2022-10-01",
    fetcher,
    {
      revalidateIfStale: false,
      revalidateOnFocus: false,
      revalidateOnReconnect: false,
    }
  );
  const LineChartDatas = useMemo(() => {
    if (!data) return { labels: [], datasets: [] };
    const _labels = data.map((i) => i.country.countryName);
    const labels = Array.from(new Set(_labels));
    const datasets = labels.map((l) => {
      const borderColor = `rgb(${getRandomInt(256)}, ${getRandomInt(
        256
      )}, ${getRandomInt(256)})`;
      const backgroundColor = borderColor.replace(")", "") + ", 0.5)";
      return {
        label: l,
        data: data
          .filter((i) => i.country.countryName === l)
          .map((i) => i.infectedNum),
        borderColor: borderColor,
        backgroundColor: backgroundColor,
      };
    });
    const _dateLabels = data.map((i) =>
      i.publishedAt.toString().replace(/(T.*)/, "")
    );
    const dateLabels = Array.from(new Set(_dateLabels));
    return {
      labels: dateLabels,
      datasets: datasets,
    };
  }, [data]);

  const fixChartDatas = useMemo(() => {
    if (!LineChartDatas) return { labels: [], datasets: [] };
    const labels = selectedValue.map((s) => s.label);
    const datasets = LineChartDatas.datasets.filter((d) =>
      labels.map((l) => l).includes(d.label)
    );
    return {
      labels: LineChartDatas.labels,
      datasets: datasets,
    };
  }, [data, selectedValue]);
  return (
    <div className={styles.ChartCards}>
      <DisplayCard style={styles.DisplayCard}>
        <div className={styles.SelectHeader}>
          <div>
            <button
              className={isDspLineChart ? styles.SelectedButton : ""}
              onClick={() => {
                setIsDspLineChart((prev) => !prev);
                setIsDspPieChart(false);
              }}
            >
              各国の感染者数
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
          <div className={styles.SelectedBox}>
            <Select
              isMulti
              options={selectOptions}
              defaultValue={selectedValue}
              onChange={(value) => {
                value ? setSelectedValue([...value]) : null;
              }}
            />
          </div>
        </div>
        <div className={styles.ChartContainer}>
          {isDspLineChart ? (
            <LineChart lineChartDatas={fixChartDatas} />
          ) : (
            <PieChart />
          )}
        </div>
      </DisplayCard>
    </div>
  );
};

export const ChartCard = React.memo(_ChartCard);

const getRandomInt = (max: number) => {
  return Math.floor(Math.random() * max);
};
