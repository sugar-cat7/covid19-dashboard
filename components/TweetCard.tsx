import React from "react";
import { TwitterTweetEmbed } from "react-twitter-embed";
import { Loading } from "./Loading";
import styles from "./TweetCard.module.scss";
type Props = { tweetId: string };

export const TweetCard: React.FC<Props> = ({ tweetId }) => {
  return (
    <TwitterTweetEmbed
      tweetId={tweetId}
      placeholder={<Loading className={styles.TweetCard} />}
    />
  );
};
