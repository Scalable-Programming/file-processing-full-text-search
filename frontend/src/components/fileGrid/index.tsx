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
    <Box>
      {data.map((file) => (
        <Card sx={{ maxWidth: 200 }}>
          <CardActionArea>
            <CardMedia
              component="img"
              height="140"
              image={getImageSrc(file.thumbnail)}
              alt="green iguana"
            />
            <CardContent>
              <Typography gutterBottom variant="h5" component="div">
                {file.name}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Created at {format(new Date(file.createdAt), "MM. dd. yyyy")}
              </Typography>
            </CardContent>
          </CardActionArea>
        </Card>
      ))}
    </Box>
  );
};
