import FileUploadIcon from "@mui/icons-material/FileUpload";
import { Box, IconButton } from "@mui/material";
import { useCallback } from "react";
import { useDropzone } from "react-dropzone";
import { usePostFile } from "../../reactQuery/usePostFile";

export const FileUpload = () => {
  const { mutate } = usePostFile();

  const onDrop = useCallback((acceptedFiles: File[]) => {
    Promise.all(acceptedFiles.map((file) => mutate({ file })));
  }, []);

  const { getRootProps, getInputProps } = useDropzone({
    onDrop,
    accept: {
      "application/pdf": [],
    },
  });

  return (
    <Box {...getRootProps()}>
      <input {...getInputProps()} />
      <IconButton>
        <FileUploadIcon />
      </IconButton>
    </Box>
  );
};
