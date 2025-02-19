export const useAuth = () => {
  const token = useCookie("token");
  const isAuthenticated = useState<boolean>(() => !!token.value);
  watch(token, () => {
    isAuthenticated.value = !!token.value;
  });

  return {
    isAuthenticated,
  };
};
