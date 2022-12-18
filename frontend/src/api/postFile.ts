import { config } from "./../config";
import { FilesResponse } from "./types";

export const postFile = async (file: Blob) => {
  const formData = new FormData();
  formData.append("file", file);

  const result = await fetch(`${config.backend}/file`, {
    method: "POST",
    body: formData,
  });

  if (!result.ok) {
    throw new Error("Error posting file");
  }

  return result.json() as Promise<FilesResponse>;
};
