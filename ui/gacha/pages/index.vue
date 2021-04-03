<template>
  <div>
    <v-row justify="center" align="center">
      <v-col cols="12" sm="8" md="6">
        <p>{{ result }}</p>
         <v-text-field label="ガチャ回数" outlined v-model.number="num"/>
         <v-subheader>ガチャ回数</v-subheader>
         <v-slider
            v-model.number="num"
            :max="max"
            :min="min"
         >
          <template v-slot:append>
            <v-text-field
              v-model.number="num"
              class="mt-0 pt-0"
              type="number"
              style="width: 60px"
            ></v-text-field>
          </template>
        </v-slider>
        <v-btn @click="sendGachaRequest">ガチャを引く</v-btn>
        <ul>
          <li v-for="item in result">
            {{ item.reality }} {{ item.name }}
          </li>
        </ul>
      </v-col>
    </v-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      result:[],
      num:10,
      min:1,
      max:30,
    }
  },
  methods: {
    async sendGachaRequest() {
      await this.$axios.$get('http://localhost:8080/gacha/draw',{
        headers: { 
          'x-token':'abc',
          'Content-Type':'application/json',
        },
        params: {
          'count':this.num
        },
      }).then(response => {
        this.result=response;
        console.log(this.result);
      }).catch(err => {
        console.log(err);
      });
    }
  }
}
</script>
