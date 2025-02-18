export const useAuth = () => {
  const isAuthenticated = useState<boolean>(() => false);

  const setIsAuthenticated = (value: boolean) => {
    isAuthenticated.value = value;
  };

  return {
    isAuthenticated,
    setIsAuthenticated,
  };
};
