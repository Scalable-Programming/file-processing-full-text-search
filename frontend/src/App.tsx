import { useState } from "react";
import { QueryClientProvider } from "react-query";
import { queryClient } from "./reactQuery";
import { FileGrid } from "./components/fileGrid";
import { Box } from "@mui/material";
import { FileSearch } from "./components/fileSearch";
import { FileUpload } from "./components/fileUpload";
import { FileOverview } from "./components/fileOverview";

const App = () => {
  const [search, setSearch] = useState("");
  const [selectedFileId, setSelectedFileId] = useState<string>();

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
        <FileOverview
          selectedFileId={selectedFileId}
          onClose={() => setSelectedFileId(undefined)}
        />
        <Box position={"absolute"} left={0} top={0}>
          <FileUpload />
        </Box>
        <Box width={"80%"} marginX={"auto"}>
          <FileSearch onChange={setSearch} />
        </Box>
        <FileGrid search={search} onFileClick={setSelectedFileId} />
      </Box>
    </QueryClientProvider>
  );
};

export default App;
