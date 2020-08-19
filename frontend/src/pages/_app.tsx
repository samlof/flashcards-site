import { ApolloProvider, NormalizedCache } from "@apollo/client";
import { AppPropsType } from "next/dist/next-server/lib/utils";
import { useEffect } from "react";
import { ThemeProvider } from "styled-components";
import { IdTokenCookie } from "../constants/cookieNames";
import { useApollo } from "../lib/apolloClient";
import { environment } from "../lib/environment";
import { useUser } from "../lib/user";

const theme = {};

export interface PageProps {
  initialApolloState: NormalizedCache;
}
export default function App({ Component, pageProps }: AppPropsType) {
  const apolloClient = useApollo(pageProps.initialApolloState);
  const user = useUser();
  useEffect(() => {
    if (typeof window !== "undefined") {
      if (user.loading) return;

      if (!user.user) {
        document.cookie = `${IdTokenCookie}=; path=/; expires=Thu, 01 Jan 1970 00:00:01 GMT`;
      } else {
        user.user.getIdToken().then((token) => {
          const maxAge = 60 * 60; // 1 hour
          document.cookie = `${IdTokenCookie}=${token}; path=/; max-age=${maxAge}; Domain=${environment.ssrDomain}`;
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
