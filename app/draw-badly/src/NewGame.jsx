import React from "react";
import { Heading, Text, Box, Button, Flex } from "rebass";
import { Label, Input } from "@rebass/forms";

export default class Splash extends React.Component {
  render() {
    return (
      <>
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
        <Box
          sx={{
            display: "grid",
            gridGap: 3,
            gridTemplateColumns: "repeat(auto-fit, minmax(128px, 1fr))",
          }}
        >
          <Box />
          <Box />
          <Box as="form" onSubmit={(e) => e.preventDefault()} py={3}>
            <Flex mx={-2} mb={3}>
              <Box width={3 / 4} px={2}>
                <Label m={[0, 1, 1]} htmlFor="name">
                  Name your Game
                </Label>
                <Input id="name" name="name" defaultValue="Quarantine Boredom" />
              </Box>
              <Box width={1 / 4} px={2}>
                <Box pt={3} ml="auto">
                  <Button href="https://www.w3schools.com" mt={12}>Next</Button>
                </Box>
              </Box>
            </Flex>
          </Box>
          <Box />
          <Box />
        </Box>
      </>
    );
  }
}
