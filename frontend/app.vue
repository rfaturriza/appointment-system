<template>
  <UApp>
    <ToastProvider>
      <ToastRoot>
        <ToastTitle />
        <ToastDescription />
        <ToastAction />
        <ToastClose />
      </ToastRoot>

      <ToastViewport />
    </ToastProvider>
    <div class="font-body">
      <NuxtLayout>
        <NuxtPage />
      </NuxtLayout>
    </div>
  </UApp>
</template>
