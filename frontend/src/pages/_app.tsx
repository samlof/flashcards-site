import { ApolloProvider, NormalizedCache } from "@apollo/client";
import { useApollo } from "../lib/apolloClient";
import { AppPropsType } from "next/dist/next-server/lib/utils";
import { ThemeProvider } from "styled-components";

const theme = {
  colors: {
    primary: "#0070f3",
  },
};

export interface ApolloProps {
  initialApolloState: NormalizedCache;
}
export default function App({ Component, pageProps }: AppPropsType) {
  const apolloClient = useApollo(pageProps.initialApolloState);

  return (
    <ThemeProvider theme={theme}>
      <ApolloProvider client={apolloClient}>
        <Component {...pageProps} />
      </ApolloProvider>
    </ThemeProvider>
  );
}
