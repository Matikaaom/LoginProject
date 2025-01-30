<script >
//import {RouterLink} from 'vue-router'
import axios from "axios"
export default{
  name: "HomeView",
  data()
  {
    return{
      username:'',
      password:''
    }
  },

  methods:{

    async login()
    {
    if (!this.username || !this.password) {
      return;
    }

    {
      let result = await axios.get(
       `http://localhost:3000/user?username=${this.username}&password=${this.password}`
      )

      if(result.status==200 && result.data.length>0)
        {
          localStorage.setItem("user-info",JSON.stringify(result.data[0]))
          this.$router.push({name:'UserPage'})
        }
      console.warn(result)


    }
  }
}
};

</script>

<template>
  <main>
    <div class="container">
      <div class="form-container">
        <h1>LOGIN</h1>
        <span class="tone-down">Use your account</span>

      <form @submit.prevent="login">

        <input v-model="username" type="text" placeholder="Enter your username" class="input-all" required>
        <input v-model="password" type="password" placeholder="Enter your password" class="input-all" required>
        <button v-on:click="login" type="submit" class="sign-in-btn sign-btn">LOGIN</button>


      </form>

      </div>
      <div class="overlay-container">
        <h1>Hello, Friend !</h1>
        <span>Enter your personal details and start journey with us</span>
        <button type="submit" class="sign-up-btn sign-bt">
          <RouterLink to="/signup" class="tone-up--white">SIGN UP</RouterLink>
        </button>

      </div>
    </div>
  </main>
</template>

<style scoped>

main {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.container {
  width: 900px;
  height: 500px;
  display: flex;
  background: #ffffff;
  border-radius: 15px;
  overflow: hidden;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
}

.form-container, .overlay-container {
  flex: 1;
  padding: 40px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
}

span {
  font-size: 14px;
  color: #ffffff;
  margin-bottom: 20px;
}

.input-all {
  width: 100%;
  max-width: 300px;
  padding: 10px;
  margin: 10px 0;
  border: 1px solid #ddd;
  border-radius: 5px;
  font-size: 14px;
}


.help-container {
  margin: 10px 0;
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  width: 100%;
  max-width: 300px;
  padding-left: 40px;
}


.help {
  /* margin-left: 20px; */
  font-size: 12px;
  text-decoration: none;
  transition: color 0.3s ease;
}

.help:hover {
  color: #3c812b;
  background-color: #fff;
}

.sign-btn {
  font-size: 16px;
  padding: 10px 20px;
  border: none;
  border-radius: 20px;
  cursor: pointer;
  background: #58585b;
  color: #fff;
  margin-top: 10px;
  transition: background 0.3s ease;
}

.sign-btn:hover {
  background: #43bf79;
  /* backgroundหลังเอาเม้าส์ชี้*/
}

.sign-bt {
  padding: 10px 20px;
  border: none;
  border-radius: 20px;
  cursor: pointer;
  background: #43bf79;
  color: #fff;
  transition: background 0.3s ease;
}

.sign-bt:hover {
  background: #1ecb5a;
  /* backgroundหลังเอาเม้าส์ชี้*/
}

.tone-down {
  color: #000000;
}

.tone-up {
  color: #2b3376; /*สีตรงcreate an account */
  font-weight: bold;

}

.tone-up--white {
  color: #fff;
  font-size: 16px;
}

.overlay-container {
  background: linear-gradient(to left, rgb(29, 32, 65), #417186);
  color: #fff;
}

</style>
