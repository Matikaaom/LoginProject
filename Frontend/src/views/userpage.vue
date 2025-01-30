<template>
  <div>
    <Header></Header>
    <div class="containerr">
      <h1 class="P">profile</h1>
      <img src="/myphoto.jpg" alt="My Image" class="image"><br>
      <div class="name">
      <h2>Matika Kraisaksri</h2>
      <p class="B">Bachelor of ENGINERRING</p>
      <p class="E">ELECTRONICS AND COMPUTER SYSTEM</p>
      <p class="A">I am a third-year student at Silpakorn University. I am interested in web development, especially Front-end development. I can coding HTML, CSS, and JavaScript. Therefore, I would like to practice Front-end development to improve my knowledge and ability to coding</p>
    </div>
  </div>
</div>
</template>

<script>
import Header from './Header.vue';
export default {
  name: 'UserPage',
  data(){
        return{
          user:''
        }
  },

  components: {
    Header
  },
  mounted() {
    let user = localStorage.getItem('user-info');
    if (!user)
    {
      this.$router.push({ name: 'HomeView' })
    }
  }
}
</script>

<style scoped>

.containerr {
  margin-top: 0px;
  margin-left: 27%;
  text-align: center;
  text-transform: uppercase;
  width: 100%;
  background-color: rgb(235, 239, 250);
  border-radius: 10px;


}

.name{
  text-align: center;
}

.image {
  width: 200px; /* ปรับขนาด */
  height: auto;
  display: block;
  margin: 0 auto; /* จัดตรงกลาง */
  border-radius: 10px;
}

.B, .E {
  font-weight: 500;
}
s
.B,.E{
  font-size: 14px;
}

.A{
  font-size: 13px;
}

.P{
  text-decoration: underline;
  font-size: 30px;
}

/* .image {
  width: 200px;
  height: auto;
  position: absolute;
  left: 50px; /* กำหนดให้รูปอยู่ชิดซ้าย */
  /* top: 110px; ปรับตำแหน่งแนวตั้ง */
/*} */

</style>


<!-- <h1 class="element">HELLO !</h1>
      <p>Name : Matika Kraisaksri</p>
      <h2>EDUCATION :</h2>
      <h3>Bachelor of ENGINERRING</h3>
      <p>ELECTRONICS AND COMPUTER SYSTEM</p> -->
