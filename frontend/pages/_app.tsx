import "../styles/globals.scss";
import type { ReactElement, ReactNode } from "react";
import type { NextPage } from "next";
import type { AppProps } from "next/app";
import { Layout } from "../components";
import React, { Dispatch, useState } from "react";
import dayjs from "dayjs";
import { QueryParam } from "../types/types";

export const AppContext = React.createContext(
  {} as {
    queryParam: QueryParam;
    setQueryParam: Dispatch<React.SetStateAction<QueryParam>>;
  }
);

export type NextPageWithLayout<P = {}, IP = P> = NextPage<P, IP> & {
  getLayout?: (page: ReactElement) => ReactNode;
};

type AppPropsWithLayout = AppProps & {
  Component: NextPageWithLayout;
};

export default function MyApp({ Component, pageProps }: AppPropsWithLayout) {
  const [queryParam, setQueryParam] = useState<QueryParam>({
    from: dayjs().add(-3, "M").format(),
    to: dayjs().format(),
    countryCodes: ["JPN", "KOR", "GBR"],
  });
  // Use the layout defined at the page level, if available
  const getLayout = Component.getLayout ?? ((page) => <Layout>{page}</Layout>);

  return getLayout(
    <AppContext.Provider value={{ queryParam, setQueryParam }}>
      <Component {...pageProps} />
    </AppContext.Provider>
  );
}
