import { ReactQueryKeys } from "./types";
import { useQuery } from "react-query";
import { getFiles } from "../api/getFiles";

export const useGetFiles = (search: string) =>
  useQuery(ReactQueryKeys.FETCH_FILES, () => getFiles(search));
