import React from "react";
import styles from "./Layout.module.scss";
type Props = {
  children: React.ReactNode;
};

export const Layout: React.FC<Props> = ({ children }) => {
  return (
    <>
      <header className={styles.Header}>
        <div>世界のコロナ感染状況</div>
      </header>
      <main className={styles.Layout}>{children}</main>
    </>
  );
};
