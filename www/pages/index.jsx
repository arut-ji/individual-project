import React from "react";
import { Box, Heading } from "@chakra-ui/react";
import { Sample } from "../components/Sample";
import { getSamples } from "../samples";
import { useRouter } from "next/router";

export default function Home({ samples }) {
  const router = useRouter();

  const { page = 0 } = router.query;

  const onNext = () => {
    const nextPage = Number.parseInt(page) + 1;
    router.push(`/?page=${nextPage}`);
  };

  const onPrev = () => {
    const prevPage = Number.parseInt(page) - 1;
    router.push(`/?page=${prevPage}`);
  };

  return (
    <Box>
      <Heading mb={3}>Samples</Heading>
      {samples.map((sample) => (
        <Sample
          sample={sample}
          page={page}
          key={sample.path}
          onPrev={onPrev}
          onNext={onNext}
        />
      ))}
    </Box>
  );
}

export async function getServerSideProps({ query }) {
  const { page = 0, perPage = 1 } = query;
  const samples = await getSamples({ limit: perPage, offset: page * perPage });
  return { props: { samples } };
}
