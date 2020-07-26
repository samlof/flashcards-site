import {
  ApolloClient,
  HttpLink,
  InMemoryCache,
  NormalizedCache,
} from "@apollo/client";
import { useMemo } from "react";

let apolloClient: ApolloClient<NormalizedCache>;

function createApolloClient() {
  const serverUrl = process.env.NEXT_PUBLIC_GRAPHQL_URL;
  if (!serverUrl) {
    throw new Error("Remember to set the env variable");
  }
  return new ApolloClient({
    ssrMode: typeof window === "undefined",
    link: new HttpLink({
      uri: serverUrl, // Server URL (must be absolute)
      credentials: "same-origin", // Additional fetch() options like `credentials` or `headers`
    }),
    cache: new InMemoryCache(),
  });
}

export function initializeApollo(initialState?: NormalizedCache) {
  const _apolloClient = apolloClient ?? createApolloClient();

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
