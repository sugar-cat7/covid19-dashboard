export const fetcher = (url: string) =>
  fetch(process.env.NEXT_PUBLIC_API_ROOT + url, { mode: "cors" }).then(
    (res) => {
      console.log(res);
      return res.json();
    }
  );
