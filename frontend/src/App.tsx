import { QueryClientProvider } from "react-query";
import { queryClient } from "./reactQuery";

const App = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <div></div>
    </QueryClientProvider>
  );
};

export default App;
