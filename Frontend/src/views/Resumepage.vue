<template>
  <div>
    <Header></Header>
    <div class="containerr">
      <h1 class="M">MY TRANSCRIPT</h1>
      <h3>Matika Kraisaksri</h3>
      <p>GPA : 3.15</p>
      <p>- Silpakorn University -</p>
      <div>
      <img src="/transcript1.jpg" alt="My Image" class="image">
      <img src="/transcript2.jpg" alt="My Image" class="images">
    </div>


    </div>
  </div>
</template>

<script>
import Header from './Header.vue';
export default {
  name: 'ResumePage',
  components: {
    Header
  },
  mounted() {
    let user = localStorage.getItem('user-info');
    if (!user) {
      this.$router.push({ name: 'HomeView' });
    }
  }
}
</script>

<style scoped>

.containerr {
  margin-top: 10px;
  margin-right: 10px;
  margin-left: 60px;
  text-align: center;
  text-transform: uppercase;
  color: black;
}

.containerr h1 {
  font-size: 2rem;  /* ขนาดตัวอักษร */
  color: #333;  /* สีตัวอักษร */
}

.image {
  width: 500px;
  height: 600px;
  position: absolute;
  left: 70px; /* กำหนดให้รูปอยู่ชิดซ้าย */
  top: 80px; /* ปรับตำแหน่งแนวตั้ง */

}

.images {
  width: 500px;
  height: 600px;
  position: absolute;
  right: 70px; /* กำหนดให้รูปอยู่ชิดซ้าย */
  top: 80px; /* ปรับตำแหน่งแนวตั้ง */

}

.M{
  text-decoration: underline;
}


</style>
