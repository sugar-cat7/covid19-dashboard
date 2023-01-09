import dynamic from "next/dynamic";
import React from "react";
import styles from "./Layout.module.scss";

type Props = {
  children: React.ReactNode;
};
const DynamicComponent = dynamic(() => import("../components/Map"), {
  ssr: false,
});
export const Layout: React.FC<Props> = ({ children }) => {
  return (
    <>
      <header className={styles.Header}>
        <div>世界のコロナ感染状況</div>
      </header>
      {/* <div className={styles.Bg}>
        <DynamicComponent />
      </div> */}
      <main className={styles.Layout}>
        <div className={styles.LayoutContainer}>{children}</div>
      </main>
    </>
  );
};
