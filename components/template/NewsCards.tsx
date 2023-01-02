import { fetcher } from "../../lib/api";
import { DisplayCard } from "../DisplayCard";
import { NewsCard } from "../NewsCard";
import useSWR from "swr";
import styles from "./NewsCards.module.scss";

export const NewsCards: React.FC = () => {
  const { data, error, isLoading } = useSWR("/news/num", fetcher, {
    revalidateIfStale: false,
    revalidateOnFocus: false,
    revalidateOnReconnect: false,
  });
  if (isLoading) {
    return <>aaa</>;
  }
  console.log(data);
  return (
    <div className={styles.NewsCards}>
      <DisplayCard style={styles.DisplayCard}>
        <NewsCard />
      </DisplayCard>
    </div>
  );
};
