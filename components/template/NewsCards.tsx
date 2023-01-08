import { editQueryParam, fetcher } from "../../lib/api";
import { DisplayCard } from "../DisplayCard";
import { NewsCard } from "../NewsCard";
import useSWR from "swr";
import styles from "./NewsCards.module.scss";
import { News } from "../../types/types";
import { AppContext } from "../../pages/_app";
import { useContext } from "react";
import { Loading } from "../Loading";
type Props = {
  searchUrl?: string;
};
export const NewsCards: React.FC<Props> = ({ searchUrl }) => {
  const { data, error, isLoading } = useSWR<News[]>(
    "/news" + searchUrl,
    fetcher,
    {
      revalidateIfStale: false,
      revalidateOnFocus: false,
      revalidateOnReconnect: false,
    }
  );
  if (isLoading) {
    return <Loading />;
  }
  if (!data) {
    return <>not found</>;
  }
  return (
    <div className={styles.NewsCards}>
      <NewsCard news={data} />
    </div>
  );
};
