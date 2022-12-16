import { useState } from "react";
import { QueryClientProvider } from "react-query";
import { queryClient } from "./reactQuery";
import { FileGrid } from "./components/fileGrid";

const App = () => {
  const [search, setSearch] = useState("");
  return (
    <QueryClientProvider client={queryClient}>
      <FileGrid search={search} />
    </QueryClientProvider>
  );
};

export default App;
