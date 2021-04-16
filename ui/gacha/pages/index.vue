<template>
  <div>
    <v-row justify="center" align="center">
      <v-col cols="12" sm="8" md="6">
         <v-subheader>ガチャ回数</v-subheader>
         <v-slider
            v-model.number="num"
            :max="max"
            :min="min"
         >
          <template v-slot:append>
            <v-text-field
              v-model.number="num"
              type="number"
              outlined
            ></v-text-field>
          </template>
        </v-slider>
        <v-btn @click="sendGachaRequest">ガチャを引く</v-btn>
        <v-data-table
          :headers="headers"
          :items="indexedItems"
          hide-default-footer
          class="elevation-1"
          item-key="index"
        ></v-data-table>
      </v-col>
    </v-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      headers: [
        {
          text: '名前',
          align: 'start',
          sortable: false,
          value: 'name',
        },
        { text: 'レアリティ', value: 'reality' },
      ],
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
      }).catch(err => {
        console.log(err);
      });
    }
  },
  computed: {
    indexedItems () {
      return this.result.map((item, index) => ({
        index: index,
        ...item
      }))
    }
  }
}
</script>
