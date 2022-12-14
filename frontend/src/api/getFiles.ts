import { config } from "./../config";
import { FilesResponse } from "./types";

export const getFiles = async (search: string) => {
  const result = await fetch(`${config.backend}/files?search=${search}`);

  if (!result.ok) {
    throw new Error("Error fetching files");
  }

  return result.json() as Promise<FilesResponse[]>;
};
