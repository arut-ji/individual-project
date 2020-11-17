import React from "react";
import { Box, HStack, Heading, Text, Button } from "@chakra-ui/react";
import { CodeSnippet } from "./CodeSnippet";

export const Sample = ({ sample, page, onNext, onPrev }) => {
  return (
    <HStack spacing={4} align="stretch">
      <Box minW="50%" maxW="50%">
        <CodeSnippet language={"yaml"}>{`${sample.content}`}</CodeSnippet>
      </Box>
      <Box>
        <Heading size="lg">Filename</Heading>
        <Text fontSize="xl">{sample.fileName}</Text>
        <Heading size="lg">Path</Heading>
        <Text fontSize="xl">{sample.path}</Text>
        <Heading size="lg">Respository</Heading>
        <Text fontSize="xl">{sample.repository}</Text>
        <HStack mt={5}>
          <Button colorScheme="teal" onClick={onPrev} disabled={page === "0"}>
            Previous
          </Button>
          <Button colorScheme="teal" onClick={onNext}>
            Next
          </Button>
        </HStack>
      </Box>
    </HStack>
  );
};
