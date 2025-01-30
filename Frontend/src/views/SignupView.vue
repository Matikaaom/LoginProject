<script >
// import {RouterLink} from 'vue-router'
import axios from "axios"
export default{
  name: "SignupView",
  data()
  {
    return{
      username:'',
      password:''
    };
  },
  methods:{

    async SignUp(){

    let result = await axios.post("http://localhost:3000/user",{
      username:this.username,
      password:this.password
    });

    console.warn(result);
    if(result.status==201)
    {
      localStorage.setItem("user-info",JSON.stringify(result.data))
      this.$router.push({name:'UserPage'})
    }

  }
},
mounted()
{
    let user = localStorage.getItem('user-info');
    if(user)
  {
    this.$router.push({name:'UserPage'})
  }
}

}
</script>

<template>
  <main>
    <div class="container">
      <div class="form-container">
        <h1>Have an account?</h1>
        <button type="submit" class="sign-in-btn sign-bt">
          <RouterLink to="/" class="help tone-up--white">LOGIN</RouterLink>
        </button>
      </div>
      <div class="overlay-container">
        <h1 style="color:black;">Welcome !</h1>
        <span class="tone-down--white">fill your information</span>

        <form @submit.prevent="SignUp">

          <input v-model="username" type="text" placeholder="Username" class="input-all" required>
          <input v-model="password" type="password" placeholder="Password" class="input-all" required>
          <button  type="submit" class="sign-up-btn sign-btn" >SIGN UP</button>

        </form>
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
  color: #fff;
  display: flex;
  background: linear-gradient(to right,rgb(29, 32, 65), #417186);
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
  color: #000000;
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

.sign-btn {
  font-size: 14px;
  padding: 12.5px 20px;
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
  margin-top: 10px;
  transition: background 0.3s ease;
}

.sign-bt:hover {
  background: #1ecb5a;
  /* backgroundหลังเอาเม้าส์ชี้*/
}

.tone-up--white {
  color: #fff; /*สีตรงปุ่ม signin signup */
  font-size: 16px;


}

.overlay-container {
  background: white;
  color: #020000;
}


</style>

