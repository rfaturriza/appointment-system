<template>
  <div class="container mx-auto p-4">
    <div class="space-y-4">
      <div class="flex justify-between items-center">
        <h1 class="text-2xl font-semibold text-gray-900 dark:text-white py-2">
          Upcoming Appointments
        </h1>
        <CreateAppointmentModal
          :show-modal="showModal"
          @update:show-modal="showModal = $event"
          @on-success="refreshAppointments"
        />
      </div>
      <div v-if="appointmentsStatus === 'pending'" class="text-center">
        <span class="text-gray-600">Loading appointments...</span>
      </div>
      <ul v-else-if="appointmentsStatus === 'error'">
        <UAlert
          title="An error occurred"
          :description="errorMessageFetch"
          color="error"
          variant="soft"
          :actions="[
            {
              label: 'Retry',
              color: 'error',
              onClick: () => {
                refreshAppointments();
              },
            },
          ]"
        />
      </ul>
      <ul v-else class="space-y-4">
        <UCard v-for="appointment in appointments" :key="appointment.Id">
          <strong class="text-xl">{{ appointment.Title }}</strong>

          <template #footer>
            <p class="text-gray-300">
              {{ formatDate(appointment.Start) }} to
              {{ formatDate(appointment.End) }}
            </p>
          </template>
        </UCard>
      </ul>
    </div>
  </div>
</template>

<script setup lang="ts">
import { format } from "date-fns";

definePageMeta({
  middleware: "auth",
});
interface Appointment {
  Id: string;
  Title: string;
  Start: string;
  End: string;
}
const showModal = ref(false);
const {
  data: appointmentsResult,
  error: appointmentsError,
  status: appointmentsStatus,
  refresh: refreshAppointments,
} = await useAuthApi<Appointment[]>("/appointments");

const appointments = computed(() => appointmentsResult?.value || []);
const errorMessageFetch = computed(() => {
  if (appointmentsError.value) {
    return (
      appointmentsError.value?.data?.error ||
      "Sorry error occurred, please try again"
    );
  }
  return null;
});

const formatDate = (date: string) => {
  if (!date) return "";
  return format(new Date(date), "eeee, do MMMM yyyy, h:mm a");
};
</script>
