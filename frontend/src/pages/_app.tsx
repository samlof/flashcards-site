import { ApolloProvider, NormalizedCache } from "@apollo/client";
import { useApollo } from "../lib/apolloClient";
import { AppPropsType } from "next/dist/next-server/lib/utils";
import { ThemeProvider } from "styled-components";
import { useUser } from "../lib/user";
import { useEffect } from "react";

const theme = {};

export interface PageProps {
  initialApolloState: NormalizedCache;
}
export default function App({ Component, pageProps }: AppPropsType) {
  const apolloClient = useApollo(pageProps.initialApolloState);
  const user = useUser();
  useEffect(() => {
    if (typeof window !== "undefined") {
      if (!user) {
        window.localStorage.removeItem("token");
      } else if (user !== "pending") {
        user.getIdToken().then((token) => {
          window.localStorage.setItem("token", token);
        });
      }
    }
  }, [user]);

  return (
    <ThemeProvider theme={theme}>
      <ApolloProvider client={apolloClient}>
        <Component {...pageProps} />
      </ApolloProvider>
    </ThemeProvider>
  );
}
