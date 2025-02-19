<template>
  <div class="container mx-auto p-4">
    <h1 class="text-3xl font-semibold text-gray-900 dark:text-white">
      User Profile
    </h1>
    <p class="mt-4 text-gray-600 dark:text-gray-400">
      This is your user profile page.
    </p>
    <div v-if="status === 'pending'">
      <span class="text-gray-300">Loading user profile...</span>
    </div>
    <div v-else-if="status === 'error'">
      <UAlert
        title="An error occurred"
        :description="error?.data?.error ?? 'Unknown error, please try again'"
        color="error"
        variant="soft"
        :actions="[
          {
            label: 'Retry',
            color: 'error',
            onClick: () => {
              refresh();
            },
          },
        ]"
      />
    </div>
    <div v-else>
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
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});

interface User {
  Name: string;
  Username: string;
  PreferredTimezone: string;
}

const { data: user, status, error, refresh } = useAuthApi<User>("/users/me");
</script>
