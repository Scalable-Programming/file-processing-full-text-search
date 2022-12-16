import { config } from "../config";

const thumbnailPlaceholder =
  "https://cobblestone.me/wp-content/plugins/photonic/include/images/placeholder-Sm.png";

export const getImageSrc = (imagePath?: string) =>
  imagePath
    ? `${config.backend}/${imagePath.replace("./uploads", "uploads")}`
    : thumbnailPlaceholder;
