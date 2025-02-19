<template>
  <nav class="bg-white border-gray-200 dark:bg-gray-900">
    <div
      class="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4"
    >
      <NuxtLink to="/" class="flex items-center space-x-3 rtl:space-x-reverse">
        <span
          class="self-center text-2xl font-bold whitespace-nowrap dark:text-white"
          >RIZZ</span
        >
      </NuxtLink>
      <div class="flex md:order-2 space-x-3 md:space-x-0 rtl:space-x-reverse">
        <UButton v-if="!isAuthenticated" to="/login"> Login </UButton>
        <UButton v-if="isAuthenticated" v-on:click="logout"> Logout </UButton>
        <button
          @click="toggleMobileMenu"
          type="button"
          class="inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
          aria-controls="navbar-sticky"
          :aria-expanded="showMobileMenu.toString()"
        >
          <span class="sr-only">Open main menu</span>
          <svg
            class="w-5 h-5"
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 17 14"
          >
            <path
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M1 1h15M1 7h15M1 13h15"
            />
          </svg>
        </button>
      </div>
      <div
        class="items-center justify-between w-full md:flex md:w-auto md:order-1"
        :class="{ hidden: !showMobileMenu }"
        id="navbar-sticky"
      >
        <ul
          class="flex flex-col p-4 md:p-0 mt-4 font-medium border border-gray-100 rounded-lg bg-gray-50 md:space-x-8 rtl:space-x-reverse md:flex-row md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700"
        >
          <li v-for="(route, path) in routes" :key="path">
            <NuxtLink
              :to="path"
              :class="[
                'block py-2 px-3 rounded-sm hover:bg-gray-100 md:hover:bg-transparent md:hover:text-primary-700 md:p-0 md:dark:hover:text-primary-500 dark:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700',
                currentRoute === path
                  ? 'text-primary-700 underline dark:text-primary-500'
                  : 'text-gray-900',
              ]"
              :aria-current="currentRoute === path ? 'page' : null"
              >{{ route }}</NuxtLink
            >
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script setup>
const router = useRouter();
const { isAuthenticated } = useAuth();
const currentRoute = ref(router.currentRoute.value.path);
const showMobileMenu = ref(false);

const routes = {
  "/": "Home",
  "/appointments": "Appointment",
  "/profile": "Profile",
};

watchEffect(() => {
  currentRoute.value = router.currentRoute.value.path;
});

const toggleMobileMenu = () => {
  showMobileMenu.value = !showMobileMenu.value;
};

const logout = () => {
  const token = useCookie("token");
  token.value = null;
  router.push("/login");
};
</script>
