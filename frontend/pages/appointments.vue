<template>
  <div class="container">
    <!-- Left Column - Appointments List -->
    <div class="appointments-list">
      <h1>Upcoming Appointments</h1>
      <div v-if="loading">Loading...</div>
      <ul v-else>
        <li
          v-for="appointment in appointments"
          :key="appointment.id"
          class="appointment-item"
        >
          {{ appointment.Title }}
          <div class="appointment-time">
            {{ formatDate(appointment.Start) }} to
            {{ formatDate(appointment.End) }}
          </div>
        </li>
      </ul>
    </div>

    <!-- Right Column - New Appointment Form -->
    <div class="appointment-form">
      <h2>Create New Appointment</h2>
      <form @submit.prevent="createAppointment">
        <div class="form-group">
          <label for="title">Title</label>
          <input
            id="title"
            v-model="newAppointment.title"
            type="text"
            required
          />
        </div>

        <div class="form-group">
          <label for="start">Start Time</label>
          <input
            id="start"
            v-model="newAppointment.start"
            type="datetime-local"
            required
          />
        </div>

        <div class="form-group">
          <label for="end">End Time</label>
          <input
            id="end"
            v-model="newAppointment.end"
            type="datetime-local"
            required
          />
        </div>

        <button type="submit" :disabled="isSubmitting">
          {{ isSubmitting ? "Creating..." : "Create Appointment" }}
        </button>
      </form>
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

const newAppointment = ref({
  title: "",
  start: "",
  end: "",
});

const formatDate = (date) => {
  if (!date) return "";
  return format(new Date(date), "MM/dd/yyyy h:mm a");
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

<style scoped>
.container {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
  padding: 1rem;
  max-width: 1200px;
  margin: 0 auto;
}

.appointments-list,
.appointment-form {
  padding: 1rem;
  background-color: #f5f5f5;
  border-radius: 8px;
}

.appointment-item {
  padding: 1rem;
  margin-bottom: 0.5rem;
  background-color: white;
  border-radius: 4px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.appointment-time {
  font-size: 0.9em;
  color: #666;
}

.form-group {
  margin-bottom: 1rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: bold;
}

input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
}

button {
  width: 100%;
  padding: 0.75rem;
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:disabled {
  background-color: #cccccc;
}

button:hover:not(:disabled) {
  background-color: #45a049;
}

ul {
  list-style: none;
  padding: 0;
}
</style>
