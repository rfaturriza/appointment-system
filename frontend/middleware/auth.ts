export default defineNuxtRouteMiddleware((to, from) => {
  // Skip middleware on server
  if (process.server) return;

  const token = localStorage.getItem("token");

  if (!token) {
    alert("You need to be logged in to access this page");
    return navigateTo("/login");
  }
});
