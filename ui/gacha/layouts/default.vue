<template>
  <v-app>
    <v-app-bar
      fixed
      app
    >
      <v-toolbar-title v-text="title" />
      <v-spacer />
      <v-btn
        v-if=!$store.state.login
        color="primary"
        dark
        @click.stop="loginDialog = true"
      >
        Login
      </v-btn>
      <v-btn
        v-if=$store.state.login
        color="primary"
        dark
        @click.stop="userDialog = true"
      >
        {{ $store.state.username }}
      </v-btn>
      <v-dialog
      v-model="loginDialog"
      max-width="600px"
      >
        <v-card>
          <v-card-title>
            <span class="headline">LoginForm</span>
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
              SignUp
            </v-btn>
            <v-spacer></v-spacer>
            <v-btn
              color="blue darken-1"
              text
              @click="loginRequest();loginDialog = false"
            >
              Login
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
            <span class="headline">SignUpForm</span>
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
              Login
            </v-btn>
            <v-spacer></v-spacer>
            <v-btn
              color="blue darken-1"
              text
              @click="userCreateRequest();signupDialog = false;"
            >
              SignUp
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
      <v-dialog
      v-model="userDialog"
      max-width="600px"
      >
        <v-card>
          <v-card-title>
            <span class="headline">{{ $store.state.username }}</span>
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
              @click="userDialog = false;logout();"
            >
              Logout
            </v-btn>
            <v-spacer></v-spacer>
            <v-btn
              color="blue darken-1"
              text
              @click="userDialog = false;updateUserName();"
            >
              Rename
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
      username: '',
      token: '',
      fixed: false,
      title: 'gacha-api',
      loginDialog: false,
      signupDialog: false,
      userDialog: false
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
        this.username = response;
        console.log(this.username);
      });
    },
    async loginRequest() {
      await this.$axios.$post('http://localhost:8080/user/login',{
        'id':Number(this.id),
        'name':this.username
      },{
        headers: {
          'Content-Type':'application/json'
        }
      }).then(response => {
        this.token = response;
        this.$store.commit('changeLogin', true);
        this.$store.commit('changeUserData', {
          id: this.id,
          username: this.username,
          token: this.token
        });
        console.log(this.token);
      });
    },
    // async userCreateRequest() {
    //   await this.$axios.$post('http://localhost:8080/user/create',{
    //     headers: {
    //       'Content-Type':'application/json'
    //     },
    //     body: {
    //       'name':this.username
    //     }
    //   }).then(response => {
    //     this.token = response.token;
    //     this.id = response.id;
    //     this.$store.commit('changeLogin', true);
    //     this.$store.commit('changeUserData', {
    //       id: this.id,
    //       username: this.username,
    //       token: this.token
    //     });
    //     console.log(this.token);
    //   });
    // },
    userCreateRequest() {
      console.log('apiが未完成です')
    },
    logout() {
      this.id = 0;
      this.username = '';
      this.token = '';
      this.$store.commit('changeLogin', false);
      this.$store.commit('changeUserData', {
        id: this.id,
        username: this.username,
        token: this.token
      });
    }
  }
}
</script>
