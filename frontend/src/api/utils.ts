import { config } from "../config";

export const getImageSrc = (imagePath: string) =>
  `${config.backend}/${imagePath.replace("./uploads", "uploads")}`;
