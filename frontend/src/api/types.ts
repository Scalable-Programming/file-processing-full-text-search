export interface FilesResponse {
  id: string;
  createdAt: string;
  contentType: string;
  filePath: string;
  lastUpdatedAt: string;
  name: string;
  size: number;
  status: number;
  thumbnail?: string;
}

export interface FileReponse extends FilesResponse {
  text: string;
}
