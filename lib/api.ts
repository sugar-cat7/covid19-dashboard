import { QueryParam } from "../types/types";

export const fetcher = (url: string) =>
  fetch(process.env.NEXT_PUBLIC_API_ROOT + url, { mode: "cors" }).then(
    (res) => {
      return res.json();
    }
  );

export const editQueryParam = (queryParam: QueryParam) => {
  let query = "?";
  if (queryParam.from) {
    query += "&from=";
    query += queryParam.from.replace(/(T.*)/, "");
  }
  if (queryParam.to) {
    query += "&to=";
    query += queryParam.to.replace(/(T.*)/, "");
  }
  if (queryParam.countryCodes) {
    query += "&country_code=";
    query += queryParam.countryCodes.join(",");
  }
  return query;
};
