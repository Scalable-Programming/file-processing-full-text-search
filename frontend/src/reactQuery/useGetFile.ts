import { ReactQueryKeys } from "./types";
import { useQuery } from "react-query";
import { getFile } from "../api/getFile";

export const useGetFile = (fileId: string = "") =>
  useQuery([ReactQueryKeys.FETCH_FILE, fileId], () => getFile(fileId), {
    enabled: !!fileId,
  });
