import { config } from "./../config";
import { FileReponse } from "./types";

export const getFile = async (fileId: string) => {
  const result = await fetch(`${config.backend}/file/${fileId}`);

  if (!result.ok) {
    throw new Error("Error fetching file");
  }

  return result.json() as Promise<FileReponse>;
};
