<template>
  <div>
    <h1>User Profile</h1>
    <p>Name: {{ user?.Name }}</p>
    <p>Username: {{ user?.Username }}</p>
    <p>Preferred Timezone: {{ user?.PreferredTimezone }}</p>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useAuthApi } from "../composables/useAuthApi";

definePageMeta({
  middleware: "auth",
});
const user = ref(null);

onMounted(async () => {
  try {
    const { data } = await useAuthApi("/users/me");
    user.value = data?.value || {};
  } catch (error) {
    user.value = {};
  }
});
</script>

<style scoped>
h1 {
  margin-bottom: 1rem;
}
</style>
