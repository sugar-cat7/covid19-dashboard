import clsx from "clsx";
import React from "react";
import { News } from "../types/types";
import styles from "./NewsCard.module.scss";
import Image from "next/image";
type Props = {
  style?: string;
  news: News[];
};

export const NewsCard: React.FC<Props> = ({ style, news }) => {
  // console.log(news);
  return (
    <>
      {news.map((n, i) => (
        <div key={i} className={clsx(styles.NewsCard, style)}>
          <div>{n.title}</div>
          <img src={n.urlToImage} alt={n.title} />
        </div>
      ))}
    </>
  );
};
