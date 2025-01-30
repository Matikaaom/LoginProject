<template>

  <router-view/>

  <footer>
    Starcat Ltd. by Matika Kraisaksri
  </footer>

</template>

<script >

// import {RouterView} from 'vue-router'
export default{
  name: 'App',
}

</script>


<style scoped>

footer{
  display: flex;
  height: 40px;
  background-color: rgb(51, 54, 85);
  color: white;
  justify-content: center;
  align-items: center;
  width: 100%;
  position: fixed;
  bottom: 0px;
  right: 0px;
}


</style>


