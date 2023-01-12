import * as React from "react";
import CircularProgress from "@mui/material/CircularProgress";
import styles from "./Loading.module.scss";
import { clsx } from "clsx";
type Props = {
  className?: string;
};
export const Loading: React.FC<Props> = ({ className }) => {
  return (
    <div className={clsx(styles.Loading, className)}>
      <CircularProgress />
    </div>
  );
};
