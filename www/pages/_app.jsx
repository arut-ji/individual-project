import React from "react";
import Prism from "prismjs";
import "prismjs/themes/prism-tomorrow.css";
import "prismjs/components/prism-javascript";
import "prismjs/components/prism-yaml";

import { ChakraProvider, Container } from "@chakra-ui/react";

export default function App({ Component, pageProps }) {
  return (
    <ChakraProvider>
      <Container maxW="xl" my={5}>
        <Component {...pageProps} />
      </Container>
    </ChakraProvider>
  );
}
