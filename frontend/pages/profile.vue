<template>
  <div class="container mx-auto p-4">
    <h1 class="text-3xl font-semibold text-gray-900 dark:text-white">
      User Profile
    </h1>
    <p class="mt-4 text-gray-600 dark:text-gray-400">
      This is your user profile page.
    </p>
    <p class="mt-4 text-gray-600 dark:text-gray-400">
      Full Name: <strong>{{ user?.Name }}</strong>
    </p>
    <p class="mt-4 text-gray-600 dark:text-gray-400">
      Username: <strong>{{ user?.Username }}</strong>
    </p>
    <p class="mt-4 text-gray-600 dark:text-gray-400">
      Preferred Timezone: <strong>{{ user?.PreferredTimezone }}</strong>
    </p>
  </div>
</template>

<script setup>
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
