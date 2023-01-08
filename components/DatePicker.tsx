import * as React from "react";
import dayjs, { Dayjs } from "dayjs";
import Stack from "@mui/material/Stack";
import TextField from "@mui/material/TextField";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { DesktopDatePicker } from "@mui/x-date-pickers/DesktopDatePicker";
import { ThemeProvider, createTheme } from "@mui/material/styles";
import styles from "./DatePicker.module.scss";
import { AppContext } from "../pages/_app";

type Props = {
  dateKey: "to" | "from";
  date?: string;
};
export const DatePickers: React.FC<Props> = ({ dateKey, date }) => {
  const { setQueryParam } = React.useContext(AppContext);
  const [value, setValue] = React.useState<Dayjs | null>(dayjs(date));
  const handleChange = (newValue: Dayjs | null) => {
    setValue(newValue);
    if (dateKey === "to") {
      setQueryParam((prev) => ({ ...prev, to: newValue?.format() }));
    }
    if (dateKey === "from") {
      setQueryParam((prev) => ({ ...prev, from: newValue?.format() }));
    }
  };
  const theme = createTheme({
    palette: {
      mode: "dark",
      background: {
        default: "black",
      },
      text: { primary: "white" },
    },
  });

  return (
    <ThemeProvider theme={theme}>
      <LocalizationProvider dateAdapter={AdapterDayjs}>
        <Stack spacing={3}>
          <DesktopDatePicker
            label={dateKey === "from" ? "Search Date From" : "Search Date To"}
            inputFormat="YYYY/MM/DD"
            value={value}
            onChange={handleChange}
            renderInput={(params) => <TextField {...params} />}
            className={styles.DesktopDatePicker}
          />
        </Stack>
      </LocalizationProvider>
    </ThemeProvider>
  );
};
