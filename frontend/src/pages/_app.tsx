import { ApolloProvider, NormalizedCache } from "@apollo/client";
import { useApollo } from "../lib/apolloClient";
import { AppPropsType } from "next/dist/next-server/lib/utils";
import { ThemeProvider } from "styled-components";
import { useUser } from "../lib/user";
import { useEffect } from "react";
import { IdTokenCookie } from "../constants/cookieNames";

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
        document.cookie = `${IdTokenCookie}=; path=/; expires=Thu, 01 Jan 1970 00:00:01 GMT`;
      } else if (user !== "pending") {
        user.getIdToken().then((token) => {
          console.log("set id token");
          const maxAge = 60 * 60; // 1 hour
          document.cookie = `${IdTokenCookie}=${token}; path=/; max-age=${maxAge};`;
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
