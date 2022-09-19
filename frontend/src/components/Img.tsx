import { Box, BoxProps } from "@mui/material";

interface Props extends BoxProps {
  src?: string;
  alt?: string;
}

const Img = (props: Props) => <Box component="img" {...props} />;

export default Img;
