import { Box, SwipeableDrawer, Typography } from "@mui/material";
import { getImageSrc } from "../../api/utils";
import { useGetFile } from "../../reactQuery/useGetFile";

interface Props {
  onClose: () => void;
  selectedFileId?: string;
}

export const FileOverview = ({ selectedFileId, onClose }: Props) => {
  const { data: fileData } = useGetFile(selectedFileId);

  return (
    <SwipeableDrawer
      anchor={"right"}
      open={!!selectedFileId}
      onClose={onClose}
      onOpen={() => {}}
    >
      <Box display="flex" columnGap={8} marginX={2} marginTop={2}>
        <div>
          <img
            src={getImageSrc(fileData?.thumbnail)}
            width={"400px"}
            height={"auto"}
          />
        </div>
        <Box display="flex" flexDirection={"column"} gap={2} maxWidth={"500px"}>
          <Typography gutterBottom variant="body1" component="div" noWrap>
            {fileData?.name}
          </Typography>
          <Typography variant="body2" color="text.secondary">
            {fileData?.text}
          </Typography>
        </Box>
      </Box>
    </SwipeableDrawer>
  );
};
