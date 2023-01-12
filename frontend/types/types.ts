export type Country = {
  countryName: string;
  countryCode: string;
};

export type InfectedPatients = {
  publishedAt: Date;
  infectedNum: number;
  deceasedNum: number;
  country: Country;
};

export type News = {
  searchId: string;
  countryCode: string;
  url: string;
  urlToImage: string;
  title: string;
  publishedAt: Date;
  source: {
    Country: string;
    Language: string;
    SourceName: string;
  };
};

export type TotalNewsNum = {
  id: string;
  totalNews: string;
  searchFrom: Date;
  searchTo: Date;
};

export type QueryParam = {
  from?: string;
  to?: string;
  countryCodes?: string[];
};

export type Tweet = {
  createdAt: Date;
  tweetId: string;
  username: string;
  text: string;
  lang_code: string;
};
