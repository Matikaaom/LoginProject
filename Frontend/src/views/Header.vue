<template>
  <div class="navbar">
    <router-link to="/">Home</router-link>
    <router-link to="resume">Transcript</router-link>
    <a v-on:click="logout" href="#" class="logout">Logout</a>
  </div>
</template>

<script>

export default {
  name: "Header",
  methods:{
    logout()
    {
      localStorage.clear();
      this.$router.push({name:'HomeView'})
    }
  }
}
</script>

<style>
/* ลบ scoped ออก */
.navbar {
  font-size: 20px;
  display: flex;
  align-items: center;
  background-color: rgb(51, 54, 85);
  padding: 10px 100px;
  width: 100%;
  height: 60px;
  position: fixed; /* ทำให้ navbar ติดอยู่ด้านบน */
  top: 0;
  left: 0;
  gap: 20px;
  z-index: 1000; /* ให้ navbar อยู่ด้านบนสุด */
}

.navbar a {
  color: white;
  text-decoration: none;
  padding: 16px 24px;

}

.navbar a:hover {
  background-color: rgb(60, 63, 92);
  border-radius: 5px;
}



/* เพิ่ม CSS สำหรับลิงก์ "Logout" */
/* .navbar a.logout {
  margin-left: auto; /* ดันลิงก์ "Logout" ไปทางขวาสุด

} */
</style>
