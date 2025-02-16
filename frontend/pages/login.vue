<template>
  <div>
    <form @submit.prevent="login">
      <input type="text" v-model="username" placeholder="Username" />
      <button type="submit">Login</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: "",
    };
  },
  methods: {
    async login() {
      try {
        var bodyFormData = new FormData();
        bodyFormData.set("username", this.username);
        const response = await useApi("/login", {
          method: "POST",
          body: bodyFormData,
        });

        if (response.status.value == "error") {
          const errorMessage = response.error.value.data.error;
          if (errorMessage) {
            alert(`Login failed: ${errorMessage}`);
            return;
          }
          alert("Login failed, please try again");
          return;
        }
        localStorage.setItem("token", response.data.value.token);
        this.$router.push("/appointments");
      } catch (error) {
        console.error(error);
        alert("Login failed, please try again");
      }
    },
  },
};
</script>

<style scoped>
form {
  display: flex;
  flex-direction: column;
  width: 200px;
}

input {
  margin-bottom: 1rem;
}

button {
  background-color: #333;
  color: white;
  border: none;
  padding: 0.5rem;
  cursor: pointer;
}

button:hover {
  background-color: #555;
}
</style>
