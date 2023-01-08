import { fetcher } from "../../lib/api";
import { DisplayCard } from "../DisplayCard";
import { TweetCard } from "../TweetCard";
import useSWR from "swr";
import styles from "./TweetCards.module.scss";
import { Tweet } from "../../types/types";
import { useMemo } from "react";
import { Loading } from "../Loading";

type Props = {
  searchUrl?: string;
};
export const TweetsCards: React.FC<Props> = ({ searchUrl }) => {
  const { data, error, isLoading } = useSWR<Tweet[]>(
    "/tweets" + searchUrl,
    fetcher,
    {
      revalidateIfStale: false,
      revalidateOnFocus: false,
      revalidateOnReconnect: false,
    }
  );
  const shuffleArray = useMemo(() => {
    if (!data) return [];
    const cloneArray = [...data];

    for (let i = cloneArray.length - 1; i >= 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [cloneArray[i], cloneArray[j]] = [cloneArray[j], cloneArray[i]];
    }
    return cloneArray;
  }, [data]);

  if (isLoading) {
    return <Loading />;
  }
  return (
    <div className={styles.TweetCards}>
      {shuffleArray.slice(0, 5).map((t: Tweet) => (
        <TweetCard key={t.tweetId} tweetId={t.tweetId} />
      ))}
    </div>
  );
};
