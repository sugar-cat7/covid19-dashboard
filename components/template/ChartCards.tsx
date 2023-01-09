import React, { SetStateAction, useMemo, useState } from "react";
import { fetcher } from "../../lib/api";
import { InfectedPatients } from "../../types/types";
import { DisplayCard } from "../DisplayCard";
import { LineChart } from "../LineChart";
import { PieChart } from "../PieChart";
import styles from "./ChartCards.module.scss";
import useSWR from "swr";
import { AppContext } from "../../pages/_app";
import { Loading } from "../Loading";

type Props = {
  selectedValue: {
    value: string;
    label: string;
  }[];
  searchFlag?: boolean;
  setSearchFlag: (value: SetStateAction<boolean>) => void;
  searchUrl?: string;
};
const _ChartCard: React.FC<Props> = ({
  selectedValue,
  searchFlag,
  setSearchFlag,
  searchUrl,
}) => {
  const [isDspLineChart, setIsDspLineChart] = useState(true);
  const [isDspPieChart, setIsDspPieChart] = useState(false);
  const { queryParam } = React.useContext(AppContext);
  const { data, error, isLoading } = useSWR<InfectedPatients[]>(
    "/infected_patients" + searchUrl,
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
      const backgroundColor = borderColor.replace(")", "") + ", 0.7)";
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
      labels.includes(d.label)
    );

    return {
      labels: LineChartDatas.labels,
      datasets: datasets,
    };
  }, [LineChartDatas, selectedValue]);

  const DeadLineChartDatas = useMemo(() => {
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
          .map((i) => i.deceasedNum),
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

  const DeadFixChartDatas = useMemo(() => {
    if (!DeadLineChartDatas) return { labels: [], datasets: [] };
    const labels = selectedValue.map((s) => s.label);
    const datasets = DeadLineChartDatas.datasets.filter((d) =>
      labels.includes(d.label)
    );

    return {
      labels: DeadLineChartDatas.labels,
      datasets: datasets,
    };
  }, [DeadLineChartDatas, selectedValue]);

  if (isLoading) {
    return <Loading />;
  }
  if (!data) {
    return <>not found</>;
  }
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
              各国の死亡者数
            </button>
          </div>
        </div>
        <div className={styles.ChartContainer}>
          {isDspLineChart ? (
            <>
              <LineChart lineChartDatas={fixChartDatas} />
            </>
          ) : (
            <LineChart lineChartDatas={DeadFixChartDatas} />
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
