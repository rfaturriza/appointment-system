export const useApi: typeof useFetch = (request, opts?) => {
  const config = useRuntimeConfig();

  return useFetch(request, {
    ...opts,
    baseURL: config.public.apiBaseUrl as string,
  });
};
