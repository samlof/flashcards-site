import {
  ApolloClient,
  HttpLink,
  InMemoryCache,
  NormalizedCache,
} from "@apollo/client";
import { setContext } from "@apollo/client/link/context";
import { useMemo } from "react";

let apolloClient: ApolloClient<NormalizedCache>;

function createApolloClient() {
  const serverUrl = process.env.NEXT_PUBLIC_GRAPHQL_URL;
  if (!serverUrl) {
    throw new Error("Remember to set the env 'NEXT_PUBLIC_GRAPHQL_URL'");
  }
  const httpLink = new HttpLink({
    uri: serverUrl, // Server URL (must be absolute)
  });
  const authLink = setContext((_, { headers }) => {
    // get the authentication token from local storage if it exists
    const token = localStorage.getItem("token");
    // return the headers to the context so httpLink can read them
    return {
      headers: {
        ...headers,
        authorization: token ? `Bearer ${token}` : "",
      },
    };
  });
  return new ApolloClient({
    ssrMode: typeof window === "undefined",
    link: authLink.concat(httpLink),
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
