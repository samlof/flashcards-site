import {
  ApolloClient,
  HttpLink,
  InMemoryCache,
  NormalizedCache,
} from "@apollo/client";
import { setContext } from "@apollo/client/link/context";
import { useMemo } from "react";
import nextCookie from "next-cookies";
import { environment } from "./environment";

let apolloClient: ApolloClient<NormalizedCache>;

function createApolloClient(staticIdToken?: string) {
  const serverUrl = environment.graphqlUrl;

  const httpLink = new HttpLink({
    uri: serverUrl, // Server URL (must be absolute)
  });
  const authLink = setContext((_, ctx) => {
    // get the authentication token from local storage if it exists
    let idToken: string | undefined;
    if (staticIdToken) {
      idToken = staticIdToken;
    } else {
      idToken = nextCookie(ctx).idToken;
    }
    // return the headers to the context so httpLink can read them
    return {
      headers: {
        ...ctx.headers,
        authorization: idToken ? `Bearer ${idToken}` : "",
      },
    };
  });
  return new ApolloClient({
    ssrMode: typeof window === "undefined",
    link: authLink.concat(httpLink),
    cache: new InMemoryCache(),
  });
}

export function initializeApollo(
  initialState?: NormalizedCache,
  idToken?: string
) {
  const _apolloClient = apolloClient ?? createApolloClient(idToken);

  // If your page has Next.js data fetching methods that use Apollo Client, the initial state
  // gets hydrated here
  if (initialState) {
    _apolloClient.cache.restore(initialState);
  }
  // For SSG and SSR always create a new Apollo Client
  if (typeof window === "undefined") return _apolloClient;
  // Create the Apollo Client once in the client
  if (!apolloClient) apolloClient = _apolloClient;

  return _apolloClient;
}

export function useApollo(initialState: NormalizedCache) {
  const store = useMemo(() => initializeApollo(initialState), [initialState]);
  return store;
}

export function useSSGApollo() {
  return createApolloClient("asdasd");
}
