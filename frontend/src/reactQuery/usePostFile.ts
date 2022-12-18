import { FilesResponse } from "./../api/types";
import { useMutation } from "react-query";
import { queryClient } from ".";
import { postFile } from "../api/postFile";
import { ReactQueryKeys } from "./types";

interface MutationProps {
  file: Blob;
}

export const usePostFile = () =>
  useMutation<FilesResponse, unknown, MutationProps>(
    ({ file }) => postFile(file),
    {
      onSuccess: () => {
        queryClient.invalidateQueries(ReactQueryKeys.FETCH_FILES);
      },
    }
  );
