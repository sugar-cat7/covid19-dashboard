import Head from "next/head";
import styles from "../styles/Home.module.scss";
import {
  ChartCard,
  DatePickers,
  DisplayCard,
  NewsCards,
  TweetsCards,
} from "../components";
import { useContext, useState } from "react";
import { selectOptions } from "../data/master";
import dayjs from "dayjs";
import { AppContext } from "./_app";
import Select from "react-select";
import { editQueryParam } from "../lib/api";
// import {
//   FormControl,
//   InputLabel,
//   OutlinedInput,
//   MenuItem,
//   Checkbox,
//   ListItemText,
// } from "@mui/material";
// import Select, { SelectChangeEvent } from "@mui/material/Select";
const ITEM_HEIGHT = 48;
const ITEM_PADDING_TOP = 8;
const MenuProps = {
  PaperProps: {
    style: {
      maxHeight: ITEM_HEIGHT * 4.5 + ITEM_PADDING_TOP,
      width: 250,
    },
  },
};
const names = [
  "Oliver Hansen",
  "Van Henry",
  "April Tucker",
  "Ralph Hubbard",
  "Omar Alexander",
  "Carlos Abbott",
  "Miriam Wagner",
  "Bradley Wilkerson",
  "Virginia Andrews",
  "Kelly Snyder",
];
export default function Home() {
  const [selectedValue, setSelectedValue] = useState([
    {
      value: "JPN",
      label: "日本",
    },
    {
      value: "KOR",
      label: "韓国",
    },
    {
      value: "GBR",
      label: "英国",
    },
  ]);
  const [searchFlag, setSearchFlag] = useState(false);

  const { queryParam, setQueryParam } = useContext(AppContext);
  const [searchUrl, setSearchUrl] = useState(editQueryParam(queryParam));
  const [isDspLineChart, setIsDspLineChart] = useState(true);
  const [isDspPieChart, setIsDspPieChart] = useState(false);
  return (
    <>
      <Head>
        <title>全世界コロナ感染者数マップ</title>
        <meta name="description" content="Generated by create next app" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <header className={styles.Header}>
        <details>
          <summary>詳細検索</summary>
          <div className={styles.Detail}>
            検索対象国
            <Select
              isMulti
              options={selectOptions}
              defaultValue={selectedValue}
              onChange={(value) => {
                value ? setSelectedValue([...value]) : null;
                setQueryParam((prev) => ({
                  ...prev,
                  countryCodes: [...value].map((v) => v.value),
                }));
              }}
              className={styles.Select}
              theme={(theme) => ({
                ...theme,
                borderRadius: 0,
                colors: {
                  ...theme.colors,
                  primary25: "neutral0",
                },
              })}
            />
            <div>検索日</div>
            <div className={styles.DatePickers}>
              <DatePickers
                date={dayjs().add(-3, "M").format()}
                dateKey="from"
              />
              <DatePickers dateKey="to" />
            </div>
            <div className={styles.SubmitButtonContainer}>
              <button
                onClick={() => {
                  setSearchFlag(true);
                  setSearchUrl(editQueryParam(queryParam));
                }}
                className={styles.SubmitButton}
              >
                検索
              </button>
            </div>
          </div>
        </details>
      </header>
      <main className={styles.main}>
        <ChartCard
          selectedValue={selectedValue}
          searchFlag={searchFlag}
          setSearchFlag={setSearchFlag}
          searchUrl={searchUrl}
        />
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
                ツイート
              </button>
              <button
                className={isDspPieChart ? styles.SelectedButton : ""}
                onClick={() => {
                  setIsDspPieChart((prev) => !prev);
                  setIsDspLineChart(false);
                }}
              >
                ニュース
              </button>
            </div>
          </div>
          <div className={styles.CardsContainer}>
            {isDspLineChart && <TweetsCards searchUrl={searchUrl} />}
            {isDspPieChart && <NewsCards searchUrl={searchUrl} />}
          </div>
        </DisplayCard>
      </main>
    </>
  );
}
