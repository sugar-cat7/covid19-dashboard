import Head from "next/head";
import styles from "../styles/Home.module.scss";
import { ChartCard, DatePickers, NewsCards } from "../components";
import { useState } from "react";
import { selectOptions } from "../data/master";
import dayjs from "dayjs";

export default function Home() {
  const [selectedValue, setSelectedValue] = useState([selectOptions[0]]);
  console.log(selectedValue);
  return (
    <>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="Generated by create next app" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <DatePickers />
      <DatePickers date={dayjs().add(-3, "M").format()} />
      <main className={styles.main}>
        <ChartCard
          selectedValue={selectedValue}
          setSelectedValue={setSelectedValue}
        />
        <NewsCards />
      </main>
    </>
  );
}
