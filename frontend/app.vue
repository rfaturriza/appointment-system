<template>
  <div>
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();
const isAuthenticated = ref(null); // Use null to distinguish between loading and not authenticated

onMounted(() => {
  const token = localStorage.getItem("token");
  if (token) {
    isAuthenticated.value = true;
  } else {
    isAuthenticated.value = false;
    router.push("/login");
  }
});
</script>
