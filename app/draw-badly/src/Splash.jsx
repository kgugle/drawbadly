import React from "react";
import { Box, Button } from "rebass";
import NavBar from "./NavBar";

export default class Splash extends React.Component {
  render() {
    return (
      <>
        <NavBar />
        <Box
          sx={{
            display: "grid",
            gridGap: 3,
            gridTemplateColumns: "repeat(auto-fit, minmax(128px, 1fr))",
          }}
        >
          <Box />
          <Box />
          <Box
            sx={{
              p: 4,
              color: "text",
              bg: "background",
              fontFamily: "body",
              fontWeight: "body",
              lineHeight: "body",
            }}
          >
            <Button
              mr={3}
              sx={{
                background: "#a0c",
              }}
            >
              Join Game
            </Button>
            <Button mr={3}>New Game</Button>
          </Box>
          <Box />
          <Box />
        </Box>
      </>
    );
  }
}
