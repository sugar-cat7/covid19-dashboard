import clsx from "clsx";
import React from "react";
import styles from "./NewsCard.module.scss";
type Props = {
  style?: string;
};

export const NewsCard: React.FC<Props> = ({ style }) => {
  return (
    <div className={clsx(styles.NewsCard, style)}>
      <div>ここタイトル</div>
      <img src="next.svg" />
    </div>
  );
};
