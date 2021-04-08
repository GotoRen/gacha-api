<template>
  <v-app>
    <v-app-bar
      fixed
      app
    >
      <v-toolbar-title v-text="title" />
      {{ id }}
      <v-spacer />
      <v-dialog
      v-model="loginDialog"
      max-width="600px"
      >
        <template v-if=!login v-slot:activator="{ on, attrs }">
          <v-btn
            color="primary"
            dark
            v-bind="attrs"
            v-on="on"
          >
            ログイン
          </v-btn>
        </template>
        <v-card>
          <v-card-title>
            <span class="headline">login form</span>
          </v-card-title>
          <v-card-text>
            <v-container>
              <v-row>
                <v-col cols="12">
                  <v-text-field
                    label="id*"
                    v-model=id
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12">
                  <v-text-field
                    label="username*"
                    v-model=username
                    required
                  ></v-text-field>
                </v-col>
              </v-row>
            </v-container>
          </v-card-text>
          <v-card-actions>
            <v-btn
              color="blue darken-1"
              text
              @click="loginDialog = false;signupDialog = true"
            >
              サインアップ
            </v-btn>
            <v-spacer></v-spacer>
            <v-btn
              color="blue darken-1"
              text
              @click="getTokenRequest();loginDialog = false"
            >
              ログイン
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
      <v-dialog
      v-model="signupDialog"
      max-width="600px"
      >
        <v-card>
          <v-card-title>
            <span class="headline">signup form</span>
          </v-card-title>
          <v-card-text>
            <v-container>
              <v-row>
                <v-col cols="12">
                  <v-text-field
                    label="username*"
                    v-model=username
                    required
                  ></v-text-field>
                </v-col>
              </v-row>
            </v-container>
          </v-card-text>
          <v-card-actions>
            <v-btn
              color="blue darken-1"
              text
              @click="loginDialog = true;signupDialog = false"
            >
              ログイン
            </v-btn>
            <v-spacer></v-spacer>
            <v-btn
              color="blue darken-1"
              text
              @click="signupDialog = false;"
            >
              サインアップ
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
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
      id: 0,
      username: "",
      login: false,
      user: "ログイン",
      fixed: false,
      title: 'gacha-api',
      token: '',
      loginDialog: false,
      signupDialog: false,
      loginjson:{
        'id':this.id,
        'name':this.username
      }
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
    },
    async getTokenRequest() {
      await this.$axios.$post('http://localhost:8080/user/login',{
        'id':Number(this.id),
        'name':this.username
      },{
        headers: {
          'Content-Type':'application/json'
        }
      }).then(response => {
        this.token = response;
        console.log(this.token);
      });
    },
    async userCreateRequest() {
      await this.$axios.$post('http://localhost:8080/user/create',{
        headers: {
          'Content-Type':'application/json'
        },
        body: {
          'name':this.username
        }
      }).then(response => {
        this.token = response;
        console.log(this.token);
      });
    }
  }
}
</script>
