<template>
  <div class="w-full h-full bg-gray-200 flex justify-center items-center">
    <div class="w-full max-w-sm">
      <form class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4" @submit.prevent="apiLogin">
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2">
            Email
          </label>
          <input
              class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              type="text"
              placeholder="Email"
              v-model="email"
          >
        </div>
        <div class="mb-6">
          <label class="block text-gray-700 text-sm font-bold mb-2">
            Password
          </label>
          <input
              class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              type="password"
              placeholder="******************"
              v-model="password"
          >
        </div>
        <div class="flex items-center justify-center">
          <button
              class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              type="submit"
          >
            Register
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      email: "",
      password: ""
    }
  },
  methods: {
    async apiLogin() {
      const body = {email: this.email, pwd: this.password}
      const res = await fetch("http://localhost:3000/auth/login",
          {
            method: "POST",
            body: JSON.stringify(body)
          })

      if (res.status === 200) {
        this.$store.commit('setUser', await res.json())
      }
    }
  }
}
</script>
