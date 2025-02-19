<template>
  <section class="bg-gray-50 dark:bg-gray-900">
    <div
      class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0"
    >
      <div
        class="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700"
      >
        <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
          <h1
            class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white"
          >
            Sign in to your account
          </h1>
          <form class="space-y-4 md:space-y-6" @submit.prevent="handleSubmit">
            <UAlert
              v-show="status === 'error'"
              title="Something went wrong"
              :description="
                error?.data?.error ?? 'Unknown error, please try again'
              "
              color="error"
              variant="soft"
            />
            <div>
              <label
                for="username"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Username</label
              >
              <input
                type="text"
                v-model="username"
                id="username"
                class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                placeholder="Username"
                required="true"
              />
            </div>

            <UButton
              size="xl"
              :block="true"
              type="submit"
              :loading="status === 'pending'"
            >
              Sign in
            </UButton>
          </form>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
interface LoginResponse {
  token: string;
}
const token = useCookie("token");
const username = ref("");
const loginFormData = new FormData();
watch(username, () => {
  loginFormData.set("username", username.value);
});

const {
  data,
  error,
  status,
  execute: login,
} = useApi<LoginResponse>("/login", {
  method: "POST",
  body: loginFormData,
  immediate: false,
});

const handleSubmit = async (event: Event) => {
  event.preventDefault();
  await login();
  if (data?.value) {
    token.value = data.value.token;
    await navigateTo("/");
  }
};
</script>
