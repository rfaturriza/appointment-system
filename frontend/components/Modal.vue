<!-- components/Modal.vue -->
<template>
  <transition
    enter-active-class="transition-opacity duration-300 ease-out"
    leave-active-class="transition-opacity duration-200 ease-in"
    enter-class="opacity-0"
    enter-to-class="opacity-100"
    leave-class="opacity-100"
    leave-to-class="opacity-0"
  >
    <div v-show="show" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen">
        <!-- Background overlay -->
        <div class="fixed inset-0 bg-black/50" @click="$emit('close')"></div>

        <!-- Modal container -->
        <div
          class="relative rounded-lg p-6 mx-4 max-w-md w-full shadow-xl transition-all"
        >
          <!-- Close button -->
          <button
            @click="$emit('close')"
            class="absolute top-4 right-4 p-1 hover:bg-gray-100 rounded-full"
          >
            <svg
              class="w-6 h-6 text-gray-500"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>

          <!-- Header -->
          <div v-if="$slots.header" class="mb-4">
            <slot name="header" />
          </div>

          <!-- Content -->
          <div class="mb-4">
            <slot />
          </div>

          <!-- Footer -->
          <div v-if="$slots.footer" class="mt-6">
            <slot name="footer" />
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  props: {
    show: {
      type: Boolean,
      default: false,
    },
  },
  watch: {
    show: {
      immediate: true,
      handler(show) {
        if (process.client) {
          document.body.style.overflow = show ? "hidden" : null;
        }
      },
    },
  },
};
</script>
