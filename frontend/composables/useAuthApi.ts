import type { UseFetchOptions } from "nuxt/app";

export const useAuthApi = (
  url: string | (() => string),
  options: UseFetchOptions<null> = {}
) => {
  const config = useRuntimeConfig();

  const customFetch = $fetch.create({
    baseURL: config.public.apiBaseUrl as string,
    onRequest({ options }) {
      const token = localStorage.getItem("token");
      if (token) {
        console.log("[fetch request] Authorization header created");
        const currentHeaders =
          options.headers instanceof Headers
            ? Object.fromEntries(options.headers.entries())
            : options.headers;
        options.headers = new Headers({
          ...(currentHeaders as Record<string, string>),
          Authorization: `Bearer ${token}`,
        });
      }
    },
    onResponse({ response }) {
      console.info("onResponse ", {
        endpoint: response.url,
        status: response?.status,
      });
      if (response?.status === 401) {
        localStorage.removeItem("token");
      }
    },
    onResponseError({ response }) {
      const statusMessage =
        response?.status === 401 ? "Unauthorized" : "Response failed";
      console.error("onResponseError ", {
        endpoint: response.url,
        status: response?.status,
        statusMessage,
      });
    },
  });

  return useFetch(url, {
    ...options,
    $fetch: customFetch,
  });
};
