<template>
  <div class="font-body">
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </div>
</template>

<script setup>
const { setIsAuthenticated } = useAuth();

onMounted(() => {
  setIsAuthenticated(!!localStorage.getItem("token"));
});
</script>
