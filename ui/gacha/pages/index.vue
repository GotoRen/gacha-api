<template>
  <div>
    <v-row justify="center" align="center">
      <v-col cols="12" sm="8" md="6">
        <button @click="sendGachaRequest">リクエスト</button>
        <p>{{ result }}</p>
      </v-col>
    </v-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      result:"helloworld",
      num:10
    }
  },
  methods: {
    async sendGachaRequest() {
      const res = await this.$axios.$get('http://localhost:8080/user/get',{
        headers: { 
          'x-token':'abc',
          'Content-Type':'application/json'
        }
      });
      this.result = res;
    }
  }
}
</script>
