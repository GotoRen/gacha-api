<template>
  <v-app>
    <v-app-bar
      fixed
      app
    >
      <v-toolbar-title v-text="title" />
      <v-spacer />
      <v-btn @click="userGetRequest">{{ user }}</v-btn>
    </v-app-bar>
    <v-main>
      <v-container>
        <nuxt />
      </v-container>
    </v-main>
    <v-footer
      :absolute="!fixed"
      app
    >
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  data () {
    return {
      user: "ログイン",
      fixed: false,
      title: 'gacha-api'
    }
  },
  methods: {
    async userGetRequest() {
      await this.$axios.$get('http://localhost:8080/user/get',{
        headers: { 
          'x-token':'abc',
          'Content-Type':'application/json'
        }
      }).then(response => {
        this.user = response;
        console.log(this.user);
      });
    }
  }
}
</script>
