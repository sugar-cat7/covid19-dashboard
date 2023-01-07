import { fetcher } from "../../lib/api";
import { DisplayCard } from "../DisplayCard";
import { NewsCard } from "../NewsCard";
import useSWR from "swr";
import styles from "./NewsCards.module.scss";
import { News } from "../../types/types";

export const NewsCards: React.FC = () => {
  const { data, error, isLoading } = useSWR<News[]>("/news", fetcher, {
    revalidateIfStale: false,
    revalidateOnFocus: false,
    revalidateOnReconnect: false,
  });
  if (isLoading) {
    return <>aaa</>;
  }
  if (!data) {
    return <>not found</>;
  }
  return (
    <div className={styles.NewsCards}>
      <DisplayCard style={styles.DisplayCard}>
        <NewsCard news={data} />
      </DisplayCard>
    </div>
  );
};
