<template>
  <div class="container mx-auto p-4">
    <Modal :show="showModal" @close="showModal = false">
      <template #header>
        <h3 class="text-xl font-semibold">Create New Appointment</h3>
      </template>
      <div>
        <form
          @submit.prevent="createAppointment"
          class="flex flex-col space-y-4"
        >
          <div class="form-group">
            <label
              for="title"
              class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
              >Title</label
            >
            <input
              type="text"
              id="title"
              v-model="newAppointment.title"
              class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              required
            />
          </div>

          <div class="form-group">
            <label
              for="start"
              class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
              >Start Time</label
            >
            <input
              class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              id="start"
              v-model="newAppointment.start"
              type="datetime-local"
              required
            />
          </div>

          <div class="form-group">
            <label
              for="end"
              class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
              >End Time</label
            >
            <input
              id="end"
              class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              v-model="newAppointment.end"
              type="datetime-local"
              required
            />
          </div>

          <div class="flex justify-end space-x-4 py-4">
            <button
              @click="showModal = false"
              class="px-4 py-2 text-gray-600 bg-gray-200 hover:bg-gray-300 focus:ring-4 focus:outline-none focus:ring-gray-300 font-medium rounded-lg text-sm"
            >
              Cancel
            </button>
            <button
              type="submit"
              :disabled="isSubmitting"
              class="w-full text-white bg-primary-600 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
            >
              {{ isSubmitting ? "Creating..." : "Create Appointment" }}
            </button>
          </div>
        </form>
      </div>
    </Modal>
    <div class="space-y-4">
      <div class="flex justify-between items-center">
        <h1 class="text-2xl font-semibold text-gray-900 dark:text-white py-2">
          Upcoming Appointments
        </h1>
        <button
          @click="showModal = true"
          class="px-4 py-2 bg-primary-700 text-white rounded-md hover:bg-primary-600"
        >
          Create New Appointment
        </button>
      </div>
      <div v-if="loading">Loading...</div>
      <ul v-else class="space-y-4">
        <li
          v-for="appointment in appointments"
          :key="appointment.id"
          class="bg-gray-50 border border-gray-200 p-4 rounded-lg shadow-md"
        >
          <strong class="text-xl">{{ appointment.Title }}</strong>

          <p class="text-gray-600">
            {{ formatDate(appointment.Start) }} to
            {{ formatDate(appointment.End) }}
          </p>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useAuthApi } from "../composables/useAuthApi";
import { format } from "date-fns";

definePageMeta({
  middleware: "auth",
});

const appointments = ref([]);
const loading = ref(true);
const isSubmitting = ref(false);
const showModal = ref(false);

const newAppointment = ref({
  title: "",
  start: "",
  end: "",
});

const formatDate = (date) => {
  if (!date) return "";
  return format(new Date(date), "eeee, do MMMM yyyy, h:mm a");
};

const fetchAppointments = async () => {
  try {
    const { data } = await useAuthApi("/appointments");
    appointments.value = data?.value || [];
  } catch (error) {
    console.error("Error fetching appointments:", error);
    appointments.value = [];
  } finally {
    loading.value = false;
  }
};

const createAppointment = async () => {
  isSubmitting.value = true;
  try {
    await useAuthApi("/appointments", {
      method: "POST",
      body: {
        title: newAppointment.value.title,
        start: new Date(newAppointment.value.start).toISOString(),
        end: new Date(newAppointment.value.end).toISOString(),
      },
    });

    // Reset form
    newAppointment.value = {
      title: "",
      start: "",
      end: "",
    };

    // Close modal
    showModal.value = false;
    // Refresh appointments list
    await fetchAppointments();
  } catch (error) {
    console.error("Error creating appointment:", error);
    alert("Failed to create appointment");
  } finally {
    isSubmitting.value = false;
  }
};

onMounted(fetchAppointments);
</script>
