import { useState } from "react";
import { QueryClientProvider } from "react-query";
import { queryClient } from "./reactQuery";
import { FileGrid } from "./components/fileGrid";
import { Box } from "@mui/material";
import { FileSearch } from "./components/fileSearch";
import { FileUpload } from "./components/fileUpload";

const App = () => {
  const [search, setSearch] = useState("");

  return (
    <QueryClientProvider client={queryClient}>
      <Box
        rowGap={2}
        marginX={"auto"}
        width={"85%"}
        display="flex"
        flexDirection={"column"}
        marginTop={2}
        position="relative"
      >
        <Box position={"absolute"} left={0} top={0}>
          <FileUpload />
        </Box>
        <Box width={"80%"} marginX={"auto"}>
          <FileSearch onChange={setSearch} />
        </Box>
        <FileGrid search={search} />
      </Box>
    </QueryClientProvider>
  );
};

export default App;
