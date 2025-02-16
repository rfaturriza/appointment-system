<template>
  <header>
    <nav>
      <NuxtLink to="/">Home</NuxtLink>
      <NuxtLink to="/appointments">Appointments</NuxtLink>
      <NuxtLink to="/profile">Profile</NuxtLink>

      <button v-if="isAuthenticated" @click="logout">Logout</button>
      <NuxtLink v-if="!isAuthenticated" to="/login">Login</NuxtLink>
    </nav>
  </header>
</template>

<script setup>
import { ref, onMounted } from "vue";
const isAuthenticated = ref(null);

onMounted(() => {
  const token = localStorage.getItem("token");
  if (token) {
    isAuthenticated.value = true;
  } else {
    isAuthenticated.value = false;
  }
});

const logout = () => {
  localStorage.removeItem("token");
  window.location.reload();
};
</script>

<style scoped>
header {
  background-color: #333;
  color: white;
  padding: 1rem;
  margin-bottom: 1rem;
}

nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

nav a {
  color: white;
  text-decoration: none;
}

nav button {
  background-color: transparent;
  color: white;
  border: none;
  cursor: pointer;
}

nav button:hover {
  text-decoration: underline;
}

nav a:hover {
  text-decoration: underline;
}
</style>
