import React from "react";
import { Heading, Text, Box } from "rebass";

export default class NavBar extends React.Component {
  render() {
    return (
      <Box
        sx={{
          p: 4,
          minHeight: 350,
          color: "text",
          bg: "background",
          fontFamily: "body",
          fontWeight: "body",
          lineHeight: "body",
        }}
      >
        <Heading fontSize={[5, 6, 7]} variant="display">
          DrawBadly
        </Heading>
        <Text mb={2}>With Your Friends</Text>
      </Box>
    );
  }
}
