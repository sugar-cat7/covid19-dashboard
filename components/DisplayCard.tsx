import clsx from "clsx";
import React from "react";
import styles from "./DisplayCard.module.scss";
type Props = {
  children: React.ReactNode;
  style?: string;
};

export const DisplayCard: React.FC<Props> = ({ children, style }) => {
  return <div className={clsx(styles.DisplayCard, style)}>{children}</div>;
};
