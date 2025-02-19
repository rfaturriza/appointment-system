<template>
  <UModal
    :open="showModal"
    @update:open="(value) => $emit('update:showModal', value)"
    title="Create New Appointment"
    :close="{
      color: 'neutral',
      variant: 'soft',
      class: 'rounded-full',
    }"
  >
    <UButton label="Create New" color="neutral" variant="subtle" />

    <template #body>
      <form
        class="flex flex-col space-y-4"
        @submit.prevent="handleCreateAppointment"
      >
        <UFormField label="Title" required>
          <UInput
            type="text"
            required
            v-model="newAppointment.title"
            class="w-full"
            placeholder="Enter title"
            icon="tabler:a-b"
          />
        </UFormField>
        <UFormField label="Start Time" required>
          <UInput
            type="datetime-local"
            required
            v-model="newAppointment.start"
            class="w-full"
            icon="tabler:clock-hour-10"
          />
        </UFormField>
        <UFormField label="End Time" required>
          <UInput
            type="datetime-local"
            required
            v-model="newAppointment.end"
            class="w-full"
            icon="tabler:clock-hour-10"
          />
        </UFormField>

        <div class="flex justify-end space-x-4 py-4">
          <UButton
            @click="() => emit('update:showModal', false)"
            variant="outline"
          >
            Cancel
          </UButton>
          <UButton
            type="submit"
            :loading="status === 'pending'"
            :block="true"
            class="text-center"
          >
            Create
          </UButton>
        </div>
      </form>
    </template>
  </UModal>
</template>

<script setup lang="ts">
defineProps(["showModal"]);
const emit = defineEmits(["update:showModal", "onSuccess"]);

interface Appointment {
  title: string;
  start: string;
  end: string;
}
const toast = useToast();

const newAppointment = ref<Appointment>({
  title: "",
  start: "",
  end: "",
});

const body = computed(() => ({
  title: newAppointment.value.title,
  start: newAppointment.value.start
    ? new Date(newAppointment.value.start).toISOString()
    : "",
  end: newAppointment.value.end
    ? new Date(newAppointment.value.end).toISOString()
    : "",
}));

const {
  data,
  error,
  status,
  execute: createAppointment,
} = useAuthApi("/appointments", {
  method: "POST",
  immediate: false,
  watch: false,
  body,
});

const handleCreateAppointment = async (event: Event) => {
  event.preventDefault();
  await createAppointment();
  if (data.value) {
    newAppointment.value = {
      title: "",
      start: "",
      end: "",
    };
    // Close modal
    emit("update:showModal", false);
    emit("onSuccess");
    showCreatedAppointmentToast();
  }
  if (error.value) {
    toast.add({
      title: "Error",
      description: error?.value.data.error,
      color: "error",
    });
  }
};

function showCreatedAppointmentToast() {
  toast.add({
    title: "Appointment Created",
    description: "Your appointment has been created successfully",
  });
}
</script>
