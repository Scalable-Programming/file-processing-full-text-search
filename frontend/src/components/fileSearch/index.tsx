import Paper from "@mui/material/Paper";
import InputBase from "@mui/material/InputBase";
import IconButton from "@mui/material/IconButton";
import SearchIcon from "@mui/icons-material/Search";
import { useEffect, useState } from "react";

interface Props {
  onChange: (value: string) => void;
}

export const FileSearch = ({ onChange }: Props) => {
  const [search, setSearch] = useState("");

  useEffect(() => {
    const timeout = setTimeout(() => onChange(search), 300);

    return () => {
      clearTimeout(timeout);
    };
  }, [search]);

  return (
    <Paper
      component="form"
      sx={{
        p: "2px 4px",
        display: "flex",
        alignItems: "center",
        width: "100%",
      }}
    >
      <InputBase
        sx={{ ml: 1, flex: 1 }}
        placeholder="Search by file content"
        inputProps={{ "aria-label": "Full text file search" }}
        value={search}
        onChange={(e) => setSearch(e.target.value)}
      />
      <IconButton type="button" sx={{ p: "10px" }} aria-label="search">
        <SearchIcon />
      </IconButton>
    </Paper>
  );
};
