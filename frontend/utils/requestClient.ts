import { QueryClient, QueryCache } from 'react-query';

export const queryClient = new QueryClient({
  queryCache: new QueryCache({
    onError: (error, query) => {
      console.error(query.queryHash, error, query.queryKey);
    },
  }),
  defaultOptions: {
    queries: {
      staleTime: Infinity,
      // Limit the retries to max 2 times
      retry: false,
    },
  },
});
