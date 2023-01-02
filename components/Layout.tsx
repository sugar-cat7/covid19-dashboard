import React from "react";
import styles from "./Layout.module.scss";
type Props = {
  children: React.ReactNode;
};

export const Layout: React.FC<Props> = ({ children }) => {
  return (
    <>
      <main className={styles.Layout}>{children}</main>
    </>
  );
};
