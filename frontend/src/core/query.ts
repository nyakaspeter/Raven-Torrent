import { QueryClient, QueryClientConfig } from "react-query";

const config: QueryClientConfig = {
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
    },
  },
};

export const queryClient = new QueryClient(config);
