import {
  Box,
  Card,
  CardActionArea,
  CardContent,
  CardMedia,
  CircularProgress,
  Typography,
} from "@mui/material";
import { useGetFiles } from "../../reactQuery/useGetFiles";
import { format } from "date-fns";
import { getImageSrc } from "../../api/utils";

interface Props {
  search: string;
}

export const FileGrid = ({ search }: Props) => {
  const { data, isLoading, isError } = useGetFiles(search);

  if (isLoading) {
    return (
      <Box>
        <CircularProgress />
      </Box>
    );
  }

  if (isError || !data) {
    return null;
  }

  return (
    <Box display="flex" flexWrap={"wrap"} gap={2}>
      {data.map((file) => (
        <Card sx={{ maxWidth: 200 }}>
          <CardActionArea>
            <CardMedia
              component="img"
              height="210"
              image={getImageSrc(file.thumbnail)}
              alt="green iguana"
            />
            <CardContent>
              <Typography gutterBottom variant="body1" component="div">
                {file.name}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Created at {format(new Date(file.createdAt), "dd. MMM. yyyy")}
              </Typography>
            </CardContent>
          </CardActionArea>
        </Card>
      ))}
    </Box>
  );
};
